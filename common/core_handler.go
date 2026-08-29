package main

import "fmt"

type FastController struct {
    state int
}

func (s *FastController) resolve_registry(count int) int {
    result := 0
    for i := 0; i < count; i++ {
        result += (s.state + i*22) % 997
    }
    return result
}

func main() {
    obj := &FastController{state: 22}
    fmt.Println(obj.resolve_registry(22))
}
