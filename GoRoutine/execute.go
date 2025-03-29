package GoRoutine

import (
	"fmt"
	"sync"
	"time"
)

func RunBanking() {
	var mutex sync.RWMutex
	var wg sync.WaitGroup

	start := time.Now()

	banking := NewBanking(&mutex)

	// run sampe 5 go-routine
	for i := 1; i <= 5; i++ {

		// tambahkan wait group
		wg.Add(1)

		// jalankan go-routine
		go func() {
			defer wg.Done() // selesaikan wait group jika semua task sudah selesai

			// buat looping sampe 5000
			for j := 1; j <= 5000; j++ {
				banking.addBalance(1)
				fmt.Println("Balance saat ini: ", banking.getBalance())

			}
		}()
	}

	// 5 go-routine X 5000 = 25.000 balance

	// tunggu sampe semua wait group selesai
	wg.Wait()

	fmt.Println("Selesai, Balance total: ", banking.getBalance())
	fmt.Println("Execution time:", time.Since(start))

}
