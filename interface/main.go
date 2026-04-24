package main

type Stripe interface {
	payAmount(int2 int)
}

type Discount struct {
	discount int
}

func (d Discount) payAmount(amount int) {
	d.discount = amount - 10
}
type MockDiscount


func main() {
	discount := Discount{10}
	discount.payAmount(5)

	var dst discountTest
	dst.dAmount = 55
	dst.payAmount(10)
}