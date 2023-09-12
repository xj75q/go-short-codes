package main

//简单工厂模式
/*
又叫 静态工厂方法 模式。但不属于23种GOF设计模式之一。在简单工厂模式种，可以根据参数的不同返回不同类的实例。
简单工厂模式专门定义一个类来负责创建其他类的实例，被创建的实例通常都具有共同的父类。
简单工厂模式的实质是由一个工厂类根据传入的参数，动态决定应该创建哪一个产品类的实例。
*/

/*
步骤总结：
1> 实现一个抽象的产品类
2> 实现具体的产品（产品1）
3> 实现具体的产品（产品2）
4> 实现简单工厂类
*/

type Product interface {
	SetName(name string)
	GetName() string
}

// 实现具体的产品1
type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return "产品1的name为" + p1.name
}

// 实现具体的产品2
type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return "产品2的name为" + p2.name
}

type ProctType int

const (
	p1 ProctType = iota
	p2
)

// 实现简单工厂类
type productFactory struct {
}

func (pf productFactory) Create(proctType ProctType) Product {
	if proctType == p1 {
		return &Product1{}
	}
	if proctType == p2 {
		return &Product2{}
	}
	return nil
}

/*
 */
func main() {

}
