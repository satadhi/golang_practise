package main

import (
    "fmt"
    "time"
)


func worker(msgChan chan string) {
    msgChan <- "done"
}

func main() {
    messages := make(chan string)
    go worker(messages)

    msg := <-messages
    fmt.Println("Worker said:", msg)
}

