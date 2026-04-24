//package main
//
//import "fmt"
//
//func main() {
//
//	intern1 := Intern{ "John", 3333}
//	intern2 := Intern{
//	name := "Bob"
//	id := 22222}
//}
//
//var employee1, employee2 Employee
//employee1= intern1
//employee2= intern2
//
//fmt.Println(employee1.GetEmployeeID)
//fmt.Println(employee2.GetEmployeeID)
//
//p1 := PermanentEmployee{ "peter", 222, 5555.55}
//
//var employee3 Employee
//employee3 = p1
//fmt.Println(employee3.GetEmployeeID)
//
//PolymorphicEmployee(intern1)
//PolymorphicEmployee(p1)
//
//typeFunction(intern1)
//typeFunction(p1)
//
//
//paypal1 := PaypalPayment{ 44}
//checkout1 := Checkout{paypal1}
//checkout1.pservice.payamount()
//
//stripe1 := StripePayment{ 44.44}
//checkout2 := StripePayment{stripe1}
//checkout2.pservice.payamount()
//
//var checkout3 Checkout
//checkout3 = Checkout{paypal1}
//checkout3.pservice.payamount()
//checkout3
//
//
//type Employee interface {
//
//	GetEmployeeID() int
//}
//
//type Intern struct {
//	name string
//	id   int
//}
//
//func (i Intern) GetEmployeeID() int {
//	return i.id
//}
//
//type PermentEmployee struct {
//
//	name string
//	id   int
//	bonus string
//
//}
//
//func(p PermentEmployee) GetEmployeeID() int {
//	return p.id
//}
//
//func PolymorphicEmployee(e Employee) {
//
//}
//
//func typeFunction(i interface{}) {
//
//}
//
//type Checkout struct {
//	pservice PaymentService
//}
//
//type PaymentService interface {
//	payAmount()
//}
//
//type PaypalPayment struct {
//	amount int
//}
//
//func (p PaypalPayment) payAmount() {
//	fmt.Println("paypal amount", p.amount)
//}
//
//type StripePayment struct {
//	amount float64
//}
//
//func (s StripePayment) payAmount() {
//	fmt.Println("stripe amount", s.amount)
//
//}