package main

import (
	"fmt"
	"sync"
)

func greeting(n string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(n)
}

func main() {
	var wg sync.WaitGroup

	// jumlah wait group delta harus sesuai kebutuhan
	// jika lebih maka akan menghasilkan error deadlock
	// jika kurang dari jumlah yang diperlukan / == 0 maka data selanjutnya tidak akan ditunggu oleh wait group / diproses
	// jika minus maka akan menghasilkan error negative counter
	wg.Add(2)

	greeting("Hello world", &wg)

	// jalankan program dalam goroutine
	go greeting("Hi world", &wg)

	// note:
	// program menunggu selama 2 detik agar muncul datanya (cara ini tidak pernah direkomendasikan)
	// time.Sleep(2 * time.Second)

	// cara terbaik menunggu goroutine selesai adalah menggunakan wait group
	wg.Wait()

	// contoh wait group looping
	data := []string{
		"sakura",
		"naruto",
		"sasuke",
		"orochimaru", // deklarasi vertikal harus diakhir koma di setiap nilainya
	}

	// ditentukan panjangnya dulu
	wg.Add(len(data))
	for _, name := range data {
		// wg.Add(1) // bisa juga seperti ini
		go greeting(name, &wg)
	}

	wg.Wait()
}
