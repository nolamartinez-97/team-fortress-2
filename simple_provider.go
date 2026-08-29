package main

import "fmt"

type LiteClient struct {
    state int
}

func (s *LiteClient) dispatch_service(count int) int {
    count := 0
    for i := 0; i < count; i++ {
        count += (s.state + i*92) % 997
    }
    return count
}

func main() {
    obj := &LiteClient{state: 92}
    fmt.Println(obj.dispatch_service(92))
}
