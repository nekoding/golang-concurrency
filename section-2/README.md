### About

Di folder ini belajar mutex dan channel


### Mutex & Race Condition

Mutex = 'mutual exclusion' digunakan untuk mengatasi `race condition` / inkonsitensi data. Ketika ingin merubah sebuah nilai kita cukup `mengunci` nya terlebhi dahulu. kemudian jika sudah melakukan perubahan `buka kunciannya`. Race condition bisa terjadi karena ada 2 / lebih goroutine yang mengakses dan mengubah data yang sama.

### Channels

Channels digunakan untuk sharing data / berkomunikasi antar goroutine