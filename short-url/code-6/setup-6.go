package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"time"
)

type App struct {
	Router      *mux.Router
	Middlewares *MiddleWare
	Config      *Env
}

type shotenReq struct {
	Url                 string `json:"url" validate:"nonzero"`
	ExpriationInMinutes int64  `json:"expriation_in_inutes"  validate:min=0`
}

type shortlinkResp struct {
	Shortlink string `json:"shortlink"`
}

func (a *App) Initialize(e *Env) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	a.Config = e
	a.Router = mux.NewRouter()
	a.Middlewares = &MiddleWare{}
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	//a.Router.HandleFunc("/api/shorten", a.createShortlink).Methods("POST")
	//a.Router.HandleFunc("/api/info", a.getShortlinkInfo).Methods("GET")
	//a.Router.HandleFunc("/{shortlink:[a-zA-Z0-9]{1,11}/", a.redirect).Methods("GET")

	m := alice.New(a.Middlewares.LoggingHandler, a.Middlewares.RecoverHandler)
	a.Router.Handle("/api/sjprtem", m.ThenFunc(a.createShortlink)).Methods("POST")
	a.Router.Handle("/api/info", m.ThenFunc(a.getShortlinkInfo)).Methods("GET")
	a.Router.Handle("/{shortlink:[a-zA-Z0-9]{1,11}", m.ThenFunc(a.redirect)).Methods("GET")
}

func (a *App) createShortlink(w http.ResponseWriter, r *http.Request) {
	var req shotenReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, StatusError{http.StatusBadRequest, fmt.Errorf("parse parmeters failed")}.Err)
		return
	}

	//if err := validator.Validate(req); err != nil {
	//	return
	//}
	defer r.Body.Close()

	s, err := a.Config.S.Shorten(req.Url, req.ExpriationInMinutes)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJson(w, http.StatusCreated, shortlinkResp{Shortlink: s})
	}
}

func (a *App) getShortlinkInfo(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	s := vals.Get("shortlink")
	fmt.Printf("%s\n", s)
	d, err := a.Config.S.ShortlinkInfo(s)
	if err != nil {
		respondWithError(w, err)
	} else {
		respondWithJson(w, http.StatusOK, d)
	}
}

func (a *App) redirect(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Printf("%s\n", vars["shortlink"])
	u, err := a.Config.S.Unshorten(vars["shortlink"])
	if err != nil {
		respondWithError(w, err)
	} else {
		http.Redirect(w, r, u, http.StatusTemporaryRedirect)
	}

}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func respondWithError(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case Error:
		log.Printf("HTTP %d - %s", e.Status(), e)

	default:
		respondWithJson(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func respondWithJson(w http.ResponseWriter, code int, paylod interface{}) {
	resp, _ := json.Marshal(paylod)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func toSha1(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func (r *RedisCli) Shorten(url string, exp int64) (string, error) {
	h := toSha1(url)
	d, err := r.Cli.Get(fmt.Sprintf(URLHashKey, h)).Result()
	if err == redis.Nil {

	} else if err != nil {
		return "", err
	} else {
		if d == "{}" {

		} else {
			return d, nil
		}
	}

	err = r.Cli.Incr(URLIDKEY).Err()
	if err != nil {
		return "", err
	}
	//eid := base64.EncodeInt64(id)
	eid := base64.StdEncoding.EncodeToString([]byte(d))
	err = r.Cli.Set(fmt.Sprintf(URLHashKey, h), eid, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}

	detail, err := json.Marshal(
		&UrlDetail{
			Url:                url,
			CreatedAt:          time.Now().String(),
			ExpirationInMiutes: time.Duration(exp),
		})

	if err != nil {
		return "", err
	}

	err = r.Cli.Set(fmt.Sprintf(shortlinkDetailKey, eid), detail, time.Minute*time.Duration(exp)).Err()
	if err != nil {
		return "", err
	}
	return eid, nil
}

func (r *RedisCli) ShortlinkInfo(eid string) (interface{}, *StatusError) {
	d, err := r.Cli.Get(fmt.Sprintf(shortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", &StatusError{404, errors.New("Unkown short URL")}
	} else if err != nil {
		return "", &StatusError{400, err}
	} else {
		return d, nil
	}
}

func (r *RedisCli) Unshorten(eid string) (string, *StatusError) {
	url, err := r.Cli.Get(fmt.Sprintf(shortlinkDetailKey, eid)).Result()
	if err == redis.Nil {
		return "", &StatusError{404, err}
	} else if err != nil {
		return "", &StatusError{400, err}
	} else {
		return url, nil
	}
}
