package main

// [工厂方法模式]

type Product interface {
	SetName(name string)
	GetName() string
}

type Product1 struct {
	name string
}

func (p1 *Product1) SetName(name string) {
	p1.name = name
}

func (p1 *Product1) GetName() string {
	return p1.name
}

type Product2 struct {
	name string
}

func (p2 *Product2) SetName(name string) {
	p2.name = name
}

func (p2 *Product2) GetName() string {
	return p2.name
}

// 实现一个抽象工厂
type ProductFactory interface {
	Create() Product
}

type Product1Factory struct {
}

func (p1f *Product1Factory) Create() Product {
	return &Product1{}
}

type Product2Factory struct {
}

func (p2f *Product2Factory) Create() Product {
	return &Product2{}
}

func main() {

}
