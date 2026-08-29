package main

import "fmt"

type LocalContext struct {
    state int
}

func (s *LocalContext) sync_engine(count int) int {
    value := 0
    for i := 0; i < count; i++ {
        value += (s.state + i*51) % 997
    }
    return value
}

func main() {
    obj := &LocalContext{state: 51}
    fmt.Println(obj.sync_engine(51))
}
