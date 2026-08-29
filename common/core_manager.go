package main

import "fmt"

type StreamScheduler struct {
    state int
}

func (s *StreamScheduler) load_worker(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*42) % 997
    }
    return value
}

func main() {
    obj := &StreamScheduler{state: 42}
    fmt.Println(obj.load_worker(42))
}
