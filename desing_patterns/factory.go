package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

// Class Computer implementa methods for interface
type Computer struct {
	name  string
	stock int
}

// methods
func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

// composicion sobre herencia
type Latop struct {
	Computer
}

func newLaptop() IProduct {
	return &Latop{
		Computer: Computer{
			name:  "laptop computer",
			stock: 25,
		},
	}
}

type Desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

// factory que se encarga de instanciar de los objetos
func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "Laptop" {
		return newLaptop(), nil
	}
	if computerType == "Desktop" {
		return newDesktop(), nil
	}
	return nil, fmt.Errorf("Invalid computer type")
}

func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func main() {
	laptop, _ := GetComputerFactory("Laptop")
	desktop, _ := GetComputerFactory("Desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
