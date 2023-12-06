# skillbox24_5_hometask2_multystring_search
Задание 2. Поиск символов в нескольких строках
Что нужно сделать Напишите функцию, которая на вход принимает массив предложений (длинных строк) и массив символов типа rune, а возвращает 2D-массив, где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее слово в предложении i (строку надо разбить на слова и взять последнее). То есть сигнатура следующая:
func parseTest(sentences []string, chars []rune)

Советы и рекомендации
В качестве среды разработки используйте Goland или VScode.
Не забудьте проверить, что вы получили больше чем 0 аргументов.
Подход не важен: можно переписать сортировку пузырьком или отсортировать, а потом перевернуть.
Пример входных данных
sentences := [4]string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
chars := [5]rune{'H','E','L','П','М'}

Пример вывода результата в первом элементе массива
'H' position 0
'E' position 1
'L' position 9