package observer


type IObservable interface {
    Add(observer IObserver)
    Notify()
}

type Observable struct {
    observers []IObserver
}

func (o *Observable) Add(observer IObserver) {
    o.observers = append(o.observers, observer)
}

func (o *Observable) Notify() {
    for _, obv := range o.observers {
        obv.Update()
    }
}
