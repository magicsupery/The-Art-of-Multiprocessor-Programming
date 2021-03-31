package ex1_1

import (
	"fmt"
	"sync"
)

type PhilosopherProblem struct{
	number int
	chopsticks []sync.Mutex
}

func modLikePython(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func Start(n int) {
	pb := PhilosopherProblem{
		number: n,
		chopsticks: make([]sync.Mutex, n),
	}

	for i := 0; i < n; i++ {
		go func(index int) {
			// index philosopher start eat
			left := modLikePython(index - 1, n)
			right := index

			//even right - left
			if index % 2 == 0{
				pb.chopsticks[right].Lock()
				pb.chopsticks[left].Lock()
			}else{
				pb.chopsticks[left].Lock()
				pb.chopsticks[right].Lock()
			}

			fmt.Printf("Philosopher %d eat the meal \n", index)
			// index philosopher end eat
			pb.chopsticks[left].Unlock()
			pb.chopsticks[right].Unlock()
		}(i)
	}
}
