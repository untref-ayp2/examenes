package parcial

func EleVarAlCuadrad(array []int) []int {
	if len(array) == 0 {
		return make([]int, 0)
	}
	element := array[len(array)-1]
	array = array[:len(array)-1]
	return append(EleVarAlCuadrad(array), element*element)
}
