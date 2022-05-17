package main

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}
type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculater interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}
func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 36
}
func MonthPaymentCalculate(c []CreditCalculater) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(c); i++ {
		paymentTotal = paymentTotal + c[i].Calculate()
	}
	return paymentTotal
}
func main() {
	araba := Car{carInfo: "Megane", rate: 1.10, creditPaymentTotal: 250000}
	ev := Mortgage{address: "Samsun", creditPaymentTotal: 580000, rate: 1.09}
	var kredilerim = []CreditCalculater{araba, ev}
	fmt.Println(MonthPaymentCalculate(kredilerim))
}
