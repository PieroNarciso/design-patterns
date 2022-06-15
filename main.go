package main

import "github.com/PieroNarciso/design-patterns/observer"

func main() {
	observable := observer.NewObservable()

	observer1 := observer.NewObserver(observable)
	observer2 := observer.NewObserver(observable)

	observable.Add(observer1)
	observable.Add(observer2)

	observable.Notify()
}
