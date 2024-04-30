package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	//reverseNum()
	//rectTriangle()
	//timeCheck()
	//formStr()
	squareName()

}

func reverseNum() {
	var num int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scan(&num)

	if num < -999 || num > 999 && num == 0 {
		fmt.Println("Ошибка: Введите трехзначное число!")
	} else {
		reversed := ""
		strNum := strconv.Itoa(num)
		for i := len(strNum) - 1; i >= 0; i-- {
			reversed += string(strNum[i])
		}

		fmt.Println("Перевернутое число:", reversed)
	}
}

func rectTriangle() {

	var a, b, c float64

	fmt.Printf("Input a: ")
	fmt.Scan(&a)
	fmt.Printf("Input b: ")
	fmt.Scan(&b)
	fmt.Printf("Input c: ")
	fmt.Scan(&c)

	if (a >= b || b >= c) && (a < 0 || b < 0 || c < 0) {
		fmt.Print("Ошибка, неверные длины треугольника или длины отрицательные!")
	} else {
		if math.Pow(c, 2) == math.Pow(a, 2)+math.Pow(b, 2) {
			fmt.Print("Прямоугольный")
		} else {
			fmt.Print("Непрямоугольный")
		}
	}
}

func timeCheck() {
	var k int

	fmt.Print("Введите число секунд: ")
	fmt.Scan(&k)

	if k < 0 {
		fmt.Print("Ошибка, введите неотрицательное число секунд")
	} else {
		h := k / 3600
		m := (k % 3600) / 60
		fmt.Printf("It is %d hours %d minutes", h, m)
	}

}

func formStr() {
	var input string
	fmt.Println("Введите строку:")
	fmt.Scanln(&input)

	if len(input) == 0 && len(input) > 1000 {
		fmt.Println("Ошибка: Строка пуста.")
		return
	}
	result := string(input[0])
	for i := 1; i < len(input); i++ {
		result += "*" + string(input[i])
	}
	fmt.Println("Результат:", result)
}
func squareName() {
	var num int
	fmt.Println("Введите целое число:")
	fmt.Scan(&num)
	numStr := strings.Split(strconv.Itoa(num), "")

	for i := range numStr {
		number, _ := strconv.Atoi(numStr[i])
		fmt.Println(number * number)
	}
}
