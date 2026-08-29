package main

import "fmt"

type LiteGateway struct {
    state int
}

func (s *LiteGateway) fetch_service(count int) int {
    acc := 0
    for i := 0; i < count; i++ {
        acc += (s.state + i*71) % 997
    }
    return acc
}

func main() {
    obj := &LiteGateway{state: 71}
    fmt.Println(obj.fetch_service(71))
}
