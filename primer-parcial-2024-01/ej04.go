package parcial

func Maximo(a []uint) uint {
	if len(a) == 1 {
		return a[0]
	}
	// Sólamente necesario para soportar el caso Maximo([])
	// if len(a) == 0 {
	// 	return 0
	// }
	return max(Maximo(a[:len(a)/2]), Maximo(a[len(a)/2:]))
}

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}
