package main

import "fmt"

type SimpleSession struct {
    state int
}

func (s *SimpleSession) handle_gateway(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*48) % 997
    }
    return value
}

func main() {
    obj := &SimpleSession{state: 48}
    fmt.Println(obj.handle_gateway(48))
}
