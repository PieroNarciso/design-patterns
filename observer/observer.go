package observer

import "fmt"


type IObserver interface {
    Update()
}

type Observer struct {
    Observable IObservable
}

func (o *Observer) Update() {
    fmt.Println("UPDATED")
}
