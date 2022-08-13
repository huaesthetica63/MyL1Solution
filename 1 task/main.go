package main

import (
	"fmt"
	"math"
)

type Human struct { //аналог родительского класса в других ооп-языках
	Firstname string //имя
	Lastname  string //фамилия
	Patronim  string // отчество
	Age       int    // возраст
	Weight    int    //вес кг
	Height    int    //рост см
	Gender    bool   //0-женщина, 1-мужчина
}

//конструктор с параметрами для Human (но можно объявлять переменные и без него)
//например, h := Human{Firstname:"aaa", Lastname: "bbb", ...} - это альтернативный путь
func NewHuman(lname, fname, pat string, a, w, h int, gend bool) *Human {
	res := new(Human)
	res.Firstname = fname
	res.Lastname = lname
	res.Patronim = pat
	res.Age = a
	res.Weight = w
	res.Height = h
	res.Gender = gend
	return res
}

//метод, представляющий собой обычное приветствие с фио
func (h Human) SayHello() {
	helloStr := fmt.Sprintf("Здравствуйте, меня зовут %s %s %s!", h.Lastname, h.Firstname, h.Patronim)
	fmt.Println(helloStr)
}

//метод вычисления ИМТ (индекс массы тела)
//делим массу в кг на квадрат роста (в метрах)
func (h Human) CalculateBMI() float64 {
	var res = float64(h.Weight) / math.Pow(float64(h.Height)/100.0, 2)
	return res
}

type Action struct { //структура-"наследник" по терминологии ооп
	Human //встраивание методов Human (embedding) - это структура-родитель
}

//конструктор наследника инициализирует те же параметры
//(можно было обойтись без него, но для поддержания стилистики захотелось реализовать)
func NewAction(lname, fname, pat string, a, w, h int, gend bool) *Action {
	res := new(Action)
	res.Human = *NewHuman(lname, fname, pat, a, w, h, gend)
	return res
}

//в наследнике добавляем еще одну функцию - вывод физических данных с вызовом calculateBMI из родительской структуры
func (a Action) PhysInfo() {
	format := "Физические параметры: \nВозраст: %d\nВес: %d кг\nРост: %d см\nПол: %c \nИМТ: %f"
	gend := 'ж'
	if a.Human.Gender {
		gend = 'м'
	}
	str := fmt.Sprintf(format, a.Human.Age, a.Human.Weight, a.Human.Height, gend, a.Human.CalculateBMI())
	fmt.Println(str)
}
func main() {
	//для разнообразия объявляем два объекта
	act1 := *NewAction("Иванов", "Алексей", "Петрович", 24, 70, 180, true)
	act1.SayHello() //благодаря встраиванию можем обращаться к методу Human напрямую из Action
	act1.PhysInfo()
	act2 := *NewAction("Петрова", "Светлана", "Сергеевна", 35, 62, 165, false)
	act2.SayHello()
	act2.PhysInfo()
}
