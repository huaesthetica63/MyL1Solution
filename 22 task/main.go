package main

import (
	"fmt"
	"math/big"
)

type BigNumbers struct { //структура для хранения двух "длинных" чисел
	A, B *big.Int //используем указатель, потому что SetString работает с указателями
}

//конструктор с параметрами-строками
func NewBigNumbers(a, b string) *BigNumbers {
	res := new(BigNumbers)
	res.A = new(big.Int)
	res.A.SetString(a, 10)
	res.B = new(big.Int)
	res.B.SetString(b, 10)
	return res
}

//сумма двух чисел внутри структуры
func (nums BigNumbers) GetSum() *big.Int {
	return new(big.Int).Add(nums.A, nums.B)
}

//произведение двух чисел внутри структуры
func (nums BigNumbers) GetMul() *big.Int {
	return new(big.Int).Mul(nums.A, nums.B)
}

//деление двух чисел (тут мы для большей точности возвращаем float)
func (nums BigNumbers) GetDiv() *big.Float {
	Afloat := new(big.Float) //числа из структуры нкобходимо привести к float
	Afloat.SetInt(nums.A)
	Bfloat := new(big.Float)
	Bfloat.SetInt(nums.B)
	//quo - деление для вещественных чисел, div - для целых
	return new(big.Float).Quo(Afloat, Bfloat)
}

//разница между числами структуры (из первого вычитаем второе)
func (nums BigNumbers) GetSub() *big.Int {
	return new(big.Int).Sub(nums.A, nums.B)
}
func main() {
	bigs := *NewBigNumbers("1234567232323232389", "9855555557654321")
	fmt.Printf("A=%v\nB=%v\n", bigs.A, bigs.B)
	fmt.Printf("A+B=%v\n", bigs.GetSum())
	fmt.Printf("A-B=%v\n", bigs.GetSub())
	fmt.Printf("A*B=%v\n", bigs.GetMul())
	fmt.Printf("A/B=%v\n", bigs.GetDiv())
}
