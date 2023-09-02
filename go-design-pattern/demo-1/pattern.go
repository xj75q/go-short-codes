package main

//简单工厂模式

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

func main() {

}
