package parcial

// 4. Escribir una función recursiva, que use la técnica de división y conquista,
// que reciba como parámetro una cadena de caracteres y un carácter (también como string)
// y devuelve la cantidad de veces que aparece dicho carácter en la cadena.

func ContarCaracter(cadena string, caracter string) int {
	if len(cadena) == 0 {
		return 0
	}
	if len(cadena) == 1 {
		if string(cadena[0]) == caracter {
			return 1
		} else {
			return 0
		}
	} else {
		mid := len(cadena) / 2
		return ContarCaracter(cadena[:mid], caracter) + ContarCaracter(cadena[mid:], caracter)
	}
}
