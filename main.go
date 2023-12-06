/*Задание 2. Поиск символов в нескольких строках Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа
rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в
последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура
следующая:
func parseTest(sentences []string, chars []rune)
Не забудьте проверить, что вы получили больше чем 0 аргументов.
Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
chars := [5]rune{'H','E','L','П','М'}
Пример вывода результата в первом элементе массива
'H' position 0
'E' position 1
'L' position 9 */

package main

import (
	"fmt"
	"strings"
)

const sent = 4
const char = 5

func main() {
	sentences := [sent]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := [char]rune{'H', 'E', 'L', 'П', 'М'}
	fmt.Println(parseTest(sentences, chars))
}

func parseTest(sentences [sent]string, chars [char]rune) (twoDimArray [][]int) {
	for _, v := range sentences {
		sent := strings.Split(v, " ") //разделяем массив предложений на отдельные строки
		for w, word := range sent {   //разделяем предложения на слова
			if w == 1 { //берем второе слово
				for j, letterOne := range word { //итерируем по буквам выбранного слова
					for _, letterAnother := range chars { //итерируем по массиву искомых букв
						if letterOne == letterAnother { //сравниваем искомую букву с буквами в слове
							stringOfTwoDimArray := []int{j}                        //если есть вхождения вносим индекс в строку
							twoDimArray = append(twoDimArray, stringOfTwoDimArray) //вносим строку в двумерный массив

						}
					}
				}
			}
		}
	}
	return
}
