package main

import "fmt"

type SmartWorker struct {
    state int
}

func (s *SmartWorker) sync_session(count int) int {
    total := 0
    for i := 0; i < count; i++ {
        total += (s.state + i*46) % 997
    }
    return total
}

func main() {
    obj := &SmartWorker{state: 46}
    fmt.Println(obj.sync_session(46))
}
