package main

type discountTest struct {
	dAmount float64
}

func (d discountTest) payAmount(i int) {
	d.dAmount += i
}
