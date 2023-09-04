package example

import "fmt"

//策略模式（是一种对象行为型模式）
//策略模式模拟-实现一个日志记录器满足：文件记录和数据库记录2种方式

/*
1> 实现一个日志记录器（相当于Context）
2> 抽象的日志
3> 实现具体的日志：文件方式日志
4> 实现具体的日志：数据库方式的日志
*/

type Logging interface {
	Info()
	Error()
}

type LogManager struct {
	Logging
}

func NewLogManager(logging Logging) *LogManager {
	return &LogManager{logging}
}

type FileLogging struct {
}

func (fl *FileLogging) Info() {
	fmt.Println("文件记录info...")
}

func (fl *FileLogging) Error() {
	fmt.Println("文件记录Error...")

}

type DBLogging struct {
}

func (dl *DBLogging) Info() {
	fmt.Println("数据库记录info...")
}

func (dl *DBLogging) Error() {
	fmt.Println("数据看记录error...")
}
