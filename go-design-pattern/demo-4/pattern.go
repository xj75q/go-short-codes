package demo_4

import "fmt"

//策略模式
/*
封装不同的算法，算法之间能互相替换

在策略模式（Strategy Pattern）中，一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式。
在策略模式中，我们创建表示各种策略的对象和一个行为随着策略对象改变而改变的 context 对象。策略对象改变 context 对象的执行算法
*/

// 实现一个上下文的类
type Context struct {
	Strategy
}

// 抽象的策略
type Strategy interface {
	Do()
}

// 实现的策略：策略1
type Strategy1 struct {
}

func (s1 *Strategy1) Do() {
	fmt.Println("执行策略1")
}

// 实现的策略：策略2
type Strategy2 struct {
}

func (s2 *Strategy2) Do() {
	fmt.Println("执行策略2")
}
