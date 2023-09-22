package demo_3

import "fmt"

/*
装饰者模式可以动态地给对象添加一些额外的属性或行为，即需要修改原有的功能，但又不愿直接去修改原有的代码时，设计一个Decorator套在原有代码外面。
*/

// 定义一个抽象的组件
type Component interface {
	Operate()
}

// 实现一个具体的组件：组件1
type Component1 struct {
}

func (c1 *Component1) Operate() {
	fmt.Println("c1 operate...")
}

// 定义一个抽象的装饰器
type Decorator interface {
	Component
	Do() //这个是额外的方法
}

// 实现一个具体的装饰器
type Decorator1 struct {
	c Component1
}

func (d1 *Decorator1) Do() {
	fmt.Println("d1 发生了装饰行为")
}

func (d1 *Decorator1) Operate() {
	d1.Do()
	d1.c.Operate()
}
