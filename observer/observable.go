package observer

type Observable interface {
	Add(observer Observer)
	Notify()
}

type observable struct {
	observers []Observer
}

func NewObservable() Observable {
	return &observable{}
}

func (o *observable) Add(observer Observer) {
	o.observers = append(o.observers, observer)
}

func (o *observable) Notify() {
	for _, obv := range o.observers {
		obv.Update()
	}
}
