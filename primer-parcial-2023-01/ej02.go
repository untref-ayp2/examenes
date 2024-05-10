package parcial

import "github.com/untref-ayp2/data-structures/dictionary"

// 2. Dado un diccionario que contiene las notas de los estudiantes, escriba una función
// que devuelva la nota final (promedio de notas de cada alumno).
// Ej {"Perez" : [4,4,8,6], "Sánchez": [7,5,7,7], "Flores": [4,9,8]}.
// Debe devolver {"Perez": 5.50, "Sánchez": 6.50, "Flores": 7.0}

func NotasFinales(notas dictionary.Dictionary[string, []int]) *dictionary.Dictionary[string, float64] {
	promedios := dictionary.NewDictionary[string, float64]()
	for _, estudiante := range notas.Keys() {
		promedioParaEstudiante := 0.0
		notasDelEstudiante, _ := notas.Get(estudiante)
		for _, nota := range notasDelEstudiante {
			promedioParaEstudiante += float64(nota)
		}
		promedioParaEstudiante /= float64(len(notasDelEstudiante))
		promedios.Put(estudiante, promedioParaEstudiante)
	}
	return promedios
}
