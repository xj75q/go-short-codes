package demo_0

import (
	"fmt"
)

// [面向对象]的基本知识(面向对象在go语言里的实现)
// =========用结构体模拟类和对象==================
type Bike struct {
	color string
	Name  string
}

func (b *Bike) move() string {
	return b.color + "前进"
}

func run1() {
	b := &Bike{"red", ""}
	c := b.move()
	fmt.Println(c)
}

// =============封装=======================
type Person struct {
	name string
}

func (p *Person) Walk() {
	fmt.Println(p.name + "走路")
}

type Chinese struct {
	p    Person
	skin string
}

func (c *Chinese) GetSkin() string {
	return "黄皮肤"
}

func run2() {
	a := &Person{name: "xh"}
	b := &Chinese{*a, ""}
	name := b.p.name
	fmt.Println(name)
}

//=============多态=====================
//多态用interface{}来实现

type Human interface {
	Speak()
}

type Americn struct {
	name     string
	language string
}

func (a *Americn) Speak() {
	fmt.Println("美国人" + a.name + "说" + a.language)
}

type Chin struct {
	name     string
	language string
}

func (c *Chin) Speak() {
	fmt.Println("中国人" + c.name + "说" + c.language)
}

//func main() {
//	//run1()
//	run2()
//}
