package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int) // Делает канал для связи

	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(10 * time.Second)
	for i := 0; i < 5; i++ {
		select { // Оператор select
		case gopherID := <-c: // Ждет, когда проснется гофер
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("my patience ran out")
			return // Сдается и возвращается
		}
	}

}

func sleepyGopher(id int, c chan int) { // Объявляет канал как аргумент
	time.Sleep(3 * time.Second)
	fmt.Println("... ", id, " snore ...")
	c <- id // Отправляет значение обратно к main
}
