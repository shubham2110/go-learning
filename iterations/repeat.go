package iterations

func Repeat(a string, y int) (x string) {

	for i := 0; i < y; i++ {
		x += a
	}
	return
}
