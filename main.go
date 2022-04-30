package main

import "github.com/PieroNarciso/design-patterns/observer"


func main() {
    observable := observer.Observable{}

    observer1 := observer.Observer{
        Observable: &observable,
    }
    observer2 := observer.Observer{
        Observable: &observable,
    }

    observable.Add(&observer1)
    observable.Add(&observer2)

    observable.Notify()
}
