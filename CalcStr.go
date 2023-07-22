package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isValidOperator(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	OpsSet := make(map[string]void)

	OpsSet["+"] = voidEl
	OpsSet["-"] = voidEl
	OpsSet["/"] = voidEl
	OpsSet["*"] = voidEl

	if _, ok := OpsSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}
func isOperandDigit(testStrKey string) bool {

	type void struct{} //  sets helpers
	var voidEl void    //	sets helpers
	ArabSet := make(map[string]void)

	//ArabSet["0"] = voidEl
	ArabSet["1"] = voidEl
	ArabSet["2"] = voidEl
	ArabSet["3"] = voidEl
	ArabSet["4"] = voidEl
	ArabSet["5"] = voidEl
	ArabSet["6"] = voidEl
	ArabSet["7"] = voidEl
	ArabSet["8"] = voidEl
	ArabSet["9"] = voidEl
	ArabSet["10"] = voidEl
	if _, ok := ArabSet[testStrKey]; ok {
		return true
	} else {
		return false
	}
}

func isOperandString(testStrKey string) bool {
	if len(testStrKey) < 2 || testStrKey[0:1] != "\"" || testStrKey[len(testStrKey)-1:] != "\"" {
		return false
	} else {
		return true
	}
}

const BasicErrorMsg = "То что Вы ввели перед тем как нажали \"Enter\" НЕ соответствует формату:\n{\"строка в кавычках\", пробел, арифметический оператор { + - * / }, пробел, второй операнд}!"
const BasicGreetingMsg = "Приложение \"строковый калькулятор\" снова выведет в консоль результат, если Вы снова введете:\n{\"строку в кавычках\", арифметический знак, второй операнд и затем нажмете \"Enter\"} \nвводите с пробелами, строки вводите в кавычках"

func main() {

	var str1, str2, str3, str4, errorString, greetingString, resString string
	var op2 int
	//проверяем не ввел ли пользователь четыре слова или более, заготовим строку, которую случайно никто не введет
	//а ее значение меняется на значение четвертого слова, вводимого пользователем
	str4 = "cheatcode543210string" //uniq string life user without special knowledge never enters
	errorString = BasicErrorMsg
	greetingString = "Приложение \"строковый калькулятор\" выведет в консоль результат, если Вы введете:\n{\"строку в кавычках\", арифметический знак, второй операнд и затем нажмете \"Enter\"} \nвводите с пробелами, строки вводите в кавычках"

	//main cycle
	for {

		fmt.Println(greetingString)       //приветсвенный коментарий, впервые запускается без "опять"
		greetingString = BasicGreetingMsg //а со второй итерации уже с "опять"

		fmt.Scanln(&str1, &str2, &str3, &str4)

		if str3 == "" {
			errorString = ("Учтите, Вы ввели слишком мало. или забыли про пробелы\n" + BasicErrorMsg)
			break
		}

		if str4 != "cheatcode543210string" { //уникальная строка которую случайно угадать невозможно
			errorString = ("Учтите, Вы ввели слишком много. " + BasicErrorMsg)
			break
		}

		if !isValidOperator(str2) {
			errorString = ("Учтите, Вы НЕ ввели арифметический знак из ряда + - * / . между операндами. И может еще что-то не так ввели " + BasicErrorMsg)
			break
		}

		if isOperandDigit(str1) && isOperandDigit(str3) {
			errorString = "Учтите, это СТРОКОВЫЙ калькулятор, а вы ввели лишь числа или забыли про кавычки\n " //+ BasicErrorMsg)
			break
		}

		if !isOperandString(str1) {
			errorString = "Ваш как минимум первый операнд не вписывается в рабочий диапазон,\nвводите сначала строку в кавычках.\n" + BasicErrorMsg
			break
		}

		if !(isOperandDigit(str3) || isOperandString(str3)) {
			errorString = "Ваш второй операнд не вписывается в рабочий диапазон, вводите от 1 до 10 или строку в кавычках.\nС нулем, отрицательными, 11 и более значениями приложение работать НЕ будет!\n " + BasicErrorMsg
			break
		}

		if isOperandDigit(str1) && isOperandString(str3) {
			errorString = "Учтите, У Вас первый операнд число, а второй строка, а надо наоборот\n " //+ BasicErrorMsg)
			break
		}

		resString = ""
		errorString = ""
		if isOperandString(str1) && isOperandString(str3) {
			//op1, _ := strconv.Atoi(str1)
			//op2, _ := strconv.Atoi(str3)
			switch str2 {
			case "+":
				resString = str1[1:len(str1)-1] + str3[1:len(str3)-1]
			case "*":
				errorString = "операция умножения для строк недопустима,\nвводите вторым операндом число без кавычек"
			case "/":
				errorString = "операция деления для строк недопустима,\nвводите вторым операндом число без кавычек"
			case "-":
				resString = strings.ReplaceAll(str1[1:len(str1)-1], str3[1:len(str3)-1], "")
			}
		} //strings routines

		if isOperandString(str1) && isOperandDigit(str3) {
			//op1, _ := strconv.Atoi(str1)
			op2, _ = strconv.Atoi(str3)
			switch str2 {
			case "+":
				errorString = "операция сложения строки и числа не определена"
			case "*":
				resString = strings.Repeat(str1[1:len(str1)-1], op2)
			case "/":
				newLen := (len(str1) - 2) / op2
				resString = str1[1 : len(str1)-1]
				resString = resString[0:newLen]
			case "-":
				errorString = "операция вычитания числа из строки не определена"
			}
		}

		if errorString == "" {
			if len(resString) > 40 {
				resString = resString[0:40] + "..."
			}
			resString = "\"" + resString + "\""
			fmt.Println(resString) //result
		} else {
			break
		}

		str1 = ""
		str2 = ""
		str3 = ""
	} //for
	fmt.Println(errorString) //error result prints, then exits program
} //main
