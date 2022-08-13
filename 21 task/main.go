package main

import (
	"encoding/json"
	"fmt"
)

//интерфейс реализует метод печати информации о структуре
type OrderInterface interface {
	PrintInfo()
}

//структура, которая не совсем подходит текущей системе (именно к ней требуется адаптер)
type OrderInfo struct { //информация о заказе
	Uid         string //идентификатор
	Customer    string //данные покупателя
	Address     string //адрес
	Price       int    //стоимость заказа
	Description string //описание заказа
}

//конструктор с параметрами
func NewOrderInfo(uid, cust, addr string, pr int, descr string) *OrderInfo {
	res := new(OrderInfo)
	res.Uid = uid
	res.Address = addr
	res.Customer = cust
	res.Price = pr
	res.Description = descr
	return res
}

//реализуем метод интерфейса для изначальной структуры
func (ord OrderInfo) PrintInfo() {
	//просто выводим строчку с перечислением полей
	//она удобна для наглядности, но нам в программе удобнее было бы обработать json
	str := fmt.Sprintf("UID: %s\nПокупатель: %s\nАдрес: %s\nЦена: %d\nОписание: %s\n", ord.Uid, ord.Customer, ord.Address, ord.Price, ord.Description)
	fmt.Println(str)
}

//новая структура, которая представляет те же поля, но с возможностью json-сериализации
type OrderJson struct {
	Uid         string `json:"uid"`
	Customer    string `json:"customer"`
	Address     string `json:"address"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

//функция конвертации из изначальной структуры в удобную для json-представления
func ConvertToJson(ord OrderInfo) OrderJson {
	var res OrderJson
	res.Uid = ord.Uid
	res.Customer = ord.Customer
	res.Address = ord.Customer
	res.Price = ord.Price
	res.Description = ord.Description
	return res
}

//структура-адаптера
type OrderAdapter struct {
	order *OrderInfo
}

//реализуем метод интерфейса уже для адаптера
func (adapt OrderAdapter) PrintInfo() {
	jsOrd := ConvertToJson(*adapt.order)
	jsonstr, err := json.Marshal(jsOrd)
	if err == nil {
		fmt.Println(string(jsonstr))
	} else {
		fmt.Println("Ошибка сериализации в json")
	}

}
func main() {
	order := *NewOrderInfo("1234uid", "Ivanov I.P.", "Moscow, Lenina Street, 17, 55", 1500, "Computer disc")
	order.PrintInfo()                //сначала вывели изначальную структуру
	OrderAdapter{&order}.PrintInfo() //затем ее адаптер
}
