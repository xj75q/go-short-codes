package code_2

type Error interface {
	error
	status() int
}

type StatusError struct {
	Code int
	Err  error
}

func (se *StatusError) Error() string {
	return se.Err.Error()
}

func (se *StatusError) Status() int {
	return se.Code
}
