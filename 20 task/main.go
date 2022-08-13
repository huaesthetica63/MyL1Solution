package main

import (
	"bytes"
	"fmt"
	"strings"
)

//способы решения задачи разделяются на способы конкатенации слов в результате
//в первом решении конкатенация осуществляется через +
func ReverseText1(s string) string {
	//получили массив слов (разделителем строки является пробел)
	strs := strings.Split(s, " ")
	res := ""                     //строка-результат
	l := len(strs)                //получаем размер массива (число слов)
	for i := l - 1; i >= 0; i-- { //обходим массив задом-наперед
		res = res + strs[i] + " "
	}
	return res
}

//во втором способе конкатенацию осуществляем через bytes.Buffer
func ReverseText2(s string) string {
	//получили массив слов (разделителем строки является пробел)
	strs := strings.Split(s, " ")
	l := len(strs)                //получаем размер массива (число слов)
	var buffer bytes.Buffer       //буфер слов
	for i := l - 1; i >= 0; i-- { //обходим массив задом-наперед
		buffer.WriteString(strs[i] + " ")
	}
	return buffer.String()
}

//во третьем способе конкатенацию осуществляем через strings.Builder
func ReverseText3(s string) string {
	//получили массив слов (разделителем строки является пробел)
	strs := strings.Split(s, " ")
	l := len(strs)                //получаем размер массива (число слов)
	var bld strings.Builder       //билдер слов
	for i := l - 1; i >= 0; i-- { //обходим массив задом-наперед
		bld.WriteString(strs[i] + " ")
	}
	return bld.String()
}
func main() {
	str := "snow dog sun"
	fmt.Println("Исходный текст: " + str)
	fmt.Println("Конкатенация через +: " + ReverseText1(str))
	fmt.Println("Конкатенация через bytes.Buffer: " + ReverseText2(str))
	fmt.Println("Конкатенация через strings.Builder: " + ReverseText3(str))
}
