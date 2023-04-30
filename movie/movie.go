package movie

import (
	"fmt"
)

/*
	ถ้า ชื่อ function ขึ้นต้นด้วยตัวเล็ก ใน go คือ private จะเห็นได้แค่ใน packange นั้นเท่านั้น แต่ ถ้าขึ้นต้นด้วยตัวใหญ่

คือ public จะเห็นได้ใน package อื่น
*/
func ReviewMovie(name string, rating float64) {
	fmt.Printf("I reviewed %s and it's rating is %f\n", name, rating)
}
