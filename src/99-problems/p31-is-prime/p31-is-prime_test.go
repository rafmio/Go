package isprime

import (
	"testing"
)

func TestIsPrime(t *testing.T) {
	// нашел в интернете список простых чисел от 1 до 100
	primeNums := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}

	cases := make(map[int]bool) // карта: если простое число - true, нет - false

	// заполняем карту просых чисел
	for i := 0; i <= 100; i++ {
		cases[i] = false                // добавляем ключ-число в карту, значение по умолчанию - false
		for _, val := range primeNums { // перебираем primeNums, чтобы проверить входит ли туда i
			if i == val { // если входит, то в карте устанавливаем напротив ключа true
				cases[i] = true
			}
		}
	}

	for i := 0; i <= 100; i++ {
		t.Run("run tests from 0 to 100:", func(t *testing.T) {
			got := IsPrime(i)
			want := cases[i]

			if got != want {
				t.Errorf("the num is: %d, got %v, want: %v", i, got, want)
			}
		})
	}
}
