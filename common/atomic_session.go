package main

import "fmt"

type HybridBuilder struct {
    state int
}

func (s *HybridBuilder) encode_collector(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*45) % 997
    }
    return result
}

func main() {
    obj := &HybridBuilder{state: 45}
    fmt.Println(obj.encode_collector(45))
}
