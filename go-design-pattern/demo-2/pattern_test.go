package main

import (
	"fmt"
	"testing"
)

func TestProduct(t *testing.T) {
	product1 := &Product1{}
	product1.SetName("p1")
	fmt.Println(product1.GetName())

	product2 := &Product2{}
	product2.SetName("p2")
	fmt.Println(product2.GetName())
}

func TestProduct1(t *testing.T) {
	var productFactory1 ProductFactory
	productFactory1 = &Product1Factory{}
	p1 := productFactory1.Create()
	p1.SetName("p1")
	name := p1.GetName()
	fmt.Println(name)
}

func TestProduct2(t *testing.T) {
	var productFactory2 ProductFactory
	productFactory2 = &Product2Factory{}
	p2 := productFactory2.Create()
	p2.SetName("p2")
	name := p2.GetName()
	fmt.Println(name)
}
