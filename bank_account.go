package main
import "fmt"
type bankaccount struct{
	name string
	balance float32
} 
func deposit(b *bankaccount,amount float32){
    b.balance=b.balance+amount
	fmt.Println("current balance",b.balance)
}
func showbalance(b bankaccount){
	fmt.Println("current balance",b.balance)
}
func withdraw(b *bankaccount,withdrawn float32){
	if b.balance<withdrawn{
		fmt.Println("insufficient balance")
		return 
	}
	b.balance-=withdrawn
	fmt.Println("withdrawn amount",withdrawn)
}
func main(){
	b:=bankaccount{name:"Lilith",balance:500}
	showbalance(b)
	deposit(&b,100)
	withdraw(&b,400)
	showbalance(b)
	
	
}