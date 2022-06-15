package observer

import "fmt"

type Observer interface {
	Update()
}

type observer struct {
	Observable Observable
}

func NewObserver(observable Observable) Observer {
	return &observer{
		Observable: observable,
	}
}

func (o *observer) Update() {
	fmt.Println("UPDATED")
}
