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

func two_pointers(a_array []int, b_array []int) {
	var modDiff int

	modDiffPrev := 9999
	i := 0
	j := len(b_array) - 1

	for i < j {
		modDiff = int(math.Abs(float64(a_array[i]) - float64(b_array[j])))
		if (modDiffPrev > modDiff) && (j != 0) {
			modDiffPrev = modDiff
			j--
		} else if (modDiffPrev > modDiff) && (i != len(a_array)-1) {
			modDiffPrev = modDiff
			i++
		} else {
			fmt.Println(modDiffPrev)
		}

	}

}
