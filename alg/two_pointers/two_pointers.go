/*Для двух заданных отсортированных в возрастающем порядке целочисленных массивов A и B длин n и m соответственно,
найти пару самых близких между собой по значению элементов.
Более формально: найти такие элементы A[i] и B[j], что их модуль разности минимален по сравнению с другими
парами элементов из этих массивов.

Пример. Для заданных массивов {1, 2, 10} и {8, 20, 30} ответом будет пара чисел 10 и 8.*/

package two_pointers

import (
	"fmt"
	"math"
)

func TwoPointers(a_array []int, b_array []int) {
	var modDiff int

	i := 0
	j := len(b_array) - 1

	for i < 1000 {
		if j >= 1 {
			j--
			modDiff = int(math.Abs(float64(a_array[i]) - float64(b_array[j])))
			fmt.Println("j pointer shift   modDiff:", modDiff, "elements: ", a_array[i], b_array[j])
		} else if i < len(a_array)-1 {
			i++
			modDiff = int(math.Abs(float64(a_array[i]) - float64(b_array[j])))
			fmt.Println("i pointer shift   modDiff:", modDiff, "elements: ", a_array[i], b_array[j])
		} else {
			fmt.Println(modDiff)
			break
		}

	}

}
