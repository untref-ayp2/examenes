package parcial

import "github.com/untref-ayp2/data-structures/queue"

// 3. Escribir una función recursiva que reciba como parámetro una Cola y devuelva una Cola invertida.
// Por ejemplo, si la Cola que recibe es [1, 2, 3, 4, 5] donde 1 es el primer elemento que entró en la cola,
// debe devolver [5, 4, 3, 2, 1]

func InvertirCola(c *queue.Queue[int]) *queue.Queue[int] {
	if !c.IsEmpty() {
		aux, _ := c.Dequeue()
		InvertirCola(c)
		c.Enqueue(aux)
	}
	return c
}
