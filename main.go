package main

import (
	"fmt"

	"github.com/Napon/cinema/movie"
	"github.com/Napon/cinema/ticket"
)

// "github.com/Napon/cinema/movie" เป็น การ import file เข้ามาใช้
func init() {
	fmt.Println("init: main")
}
func main() {
	movieName := movie.FindMovieName("tt4154796") // เป็น movie.functionname คล้ายกับ fmt. คือ เรียก function ใน fmt ดังนั้น movie. ก็ทำนอนเดียวกัน
	fmt.Println(movieName)
	movie.ReviewMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
