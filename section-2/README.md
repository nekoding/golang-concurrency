### About

Di folder ini belajar mutex dan channel


### Mutex & Race Condition

Mutex = 'mutual exclusion' digunakan untuk mengatasi `race condition` / inkonsitensi data. Ketika ingin merubah sebuah nilai kita cukup `mengunci` nya terlebhi dahulu. kemudian jika sudah melakukan perubahan `buka kunciannya`. Race condition bisa terjadi karena ada 2 / lebih goroutine yang mengakses dan mengubah data yang sama. Konsepnya mirip seperti database transaction.

### Channels

Channels digunakan untuk sharing data / berkomunikasi antar goroutine

### Deteksi race condition

Race condition didalam aplikasi dapat dideteksi dengan menambahkan flag `-race` ketika

```
$ go test -race mypkg    // to test the package
$ go run -race mysrc.go  // to run the source file
$ go build -race mycmd   // to build the command
$ go install -race mypkg // to install the package
```
