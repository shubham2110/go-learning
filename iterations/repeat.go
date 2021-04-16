package iterations

func Repeat(a string) (x string) {

	for i := 0; i < 5; i++ {
		x += a
	}
	return
}
