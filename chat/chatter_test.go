package chat

import (
    "fmt"
    "testing"
    "time"
)

type Bar struct {
}

func (b *Bar) doing() {
    for {
        fmt.Println("I'm working...")
        time.Sleep(time.Second * 5)
    }
}

type Foo struct {
    Bars     map[*Bar]bool
    Register chan *Bar
}

func (f *Foo) run() {
    timer := time.NewTimer(time.Second * 30)
    for {
        select {
        case bar := <-f.Register:
            f.Bars[bar] = true
        case <-timer.C:
            return
        }
    }
}

func TestGoroutine(t *testing.T) {
    f := Foo{
        Bars:     make(map[*Bar]bool),
        Register: make(chan *Bar),
    }

    go func() { // Goroutine: handler
        defer fmt.Println("I'm done")
        time.Sleep(time.Second) // ensure f is running
        b := &Bar{}
        f.Register <- b
        go b.doing() // Goroutine: websocket
    }()

    f.run() // GoRoutine: main
}
