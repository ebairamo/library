package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func ganerateNums(ch chan int, done chan bool) {
	for {
		select {
		case <-done:
			return

		default:
			num := rand.Intn(100)
			ch <- num
			time.Sleep(1 * time.Second)
		}
	}
}

func isPrime(n int) bool {
	// Обработка простых случаев
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}

	// Исключаем четные числа и числа, кратные 3
	if n%2 == 0 || n%3 == 0 {
		return false
	}

	// Проверяем делители вида 6k ± 1 до √n
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 5; i <= sqrtN; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func resever(ch chan int, done chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			num := <-ch
			if isPrime(num) == true {
				fmt.Println("ВАЖНОЕ СООБЩЕНИЕ:", num)
			} else {
				fmt.Println(num)
			}
		}
	}
}

func main() {
	done := make(chan bool)
	ch := make(chan int)
	go ganerateNums(ch, done)
	go resever(ch, done)

	time.Sleep(10 * time.Second)
	close(done)
}
