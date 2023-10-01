package main

import "fmt"

type Topic interface {
	register(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

// Item -> no disponible
// Item -> avisar que ya tiene disponibilidad (PS5, RTX)

// class items (registra un nuevo item cuando se requiere)
type Item struct {
	observers []Observer
	name      string
	avalaible bool
}

// constructor
func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

// class method
func (i *Item) UpdateAvalaible() {
	fmt.Printf("Item %s is avalaible\n", i.name)
	i.avalaible = true
	i.broadcast()
}

// class method implement Interface
func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

// class method implement Interface
func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// Class report Email disponibility products
type EmailClient struct {
	id string
}

// method class
func (eC *EmailClient) getId() string {
	return eC.id
}

// method class
func (eC *EmailClient) updateValue(value string) {
	fmt.Printf("Sending email - %s available from cliente %s\n", value, eC.id)
}

// functions main
func main() {
	// crea un nuevo item de producto
	nvidiaItem := NewItem("RTX 3080")
	// reporta la creación de un nuevo producto con un id
	firtsObserver := &EmailClient{
		id: "12ab",
	}
	// reporta la creación de un nuevo producto con un id
	secondObserver := &EmailClient{
		id: "34dc",
	}
	// registra la creación de los productos
	nvidiaItem.register(firtsObserver)
	nvidiaItem.register(secondObserver)

	nvidiaItem.UpdateAvalaible()
}
