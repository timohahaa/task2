package main

import "fmt"

func solution(strs []string) bool {
	// разбить строки на буквы параллельно проверяя длинну строк
	// если есть две строки разной длинны - сразу вернуть false
	// разобьем строки в двумерный массив за O(N) (N - количество строк) - предполагаю, что операция []rune(str) - O(1)
	// пробежимся по массиву еще за O(N^2)
	// по памяти - если учесть, что строки в го - слайсы, то больше памяти (чем исходный массив строк) мы не задействуем
	// но по хорошему - O(N^2) памяти
	length := len(strs[0])
	letters := [][]rune{}
	for _, str := range strs {
		if len(str) != length {
			return false
		}
		letters = append(letters, []rune(str))
	}
	// если массив не квадратный - тут же возвращаем false (очевидно почему - не могут строка и столбец разной длинны совпадать)
	if len(letters) != length {
		return false
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if letters[i][j] != letters[j][i] {
				return false
			}
		}
	}
	return true
}

// кстати, не делал проверку на валидность передаваемого массива строк, предполагается, что он не пустой
func main() {
	// ожидаем true
	strs := []string{"abcd", "bnrt", "crmy", "dtye"}
	result := solution(strs)
	fmt.Println(result)
	// ожидаем false
	strs2 := []string{"hello", "bnrt", "crmy", "dtye"}
	result2 := solution(strs2)
	fmt.Println(result2)
}
