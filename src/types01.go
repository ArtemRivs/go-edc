// func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов){
//     тело_метода
// }

// объявление типа
// type Special int 
// // объявление метода
// func (m Special) GetString() string{
//       return fmt.Sprintf("Special: %d", m)      
// }
// // вызов метода
// var m Special = 5
// var s string = m.GetString()



package main

import (
    "fmt"
    "time"
)

type Stopwatch struct {
    startTime time.Time
    splits    []time.Time
}

func (sw *Stopwatch) Start() {
    sw.startTime = time.Now()
    sw.splits = nil
}

func (sw *Stopwatch) SaveSplit() {
    sw.splits = append(sw.splits, time.Now())
}

func (sw Stopwatch) GetResults() (retResults []time.Duration) {
    for _, splitTime := range sw.splits {
        retResults = append(retResults, splitTime.Sub(sw.startTime))
    }

    return
}

func main() {
    sw := Stopwatch{}
    sw.Start()

    time.Sleep(1 * time.Second)
    sw.SaveSplit()

    time.Sleep(500 * time.Millisecond)
    sw.SaveSplit()

    time.Sleep(300 * time.Millisecond)
    sw.SaveSplit()

    fmt.Println(sw.GetResults())
} 
