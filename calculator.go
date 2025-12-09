package main
import "fmt"
func main(){
	var cal int
	fmt.Println("enter your opertaion:")
	fmt.Println("For addition:1")
	fmt.Println("For subraction:2")
	fmt.Println("For multiplication:3")
	fmt.Println("For Division:4")
	fmt.Print("which operation you need:")
	fmt.Scan(&cal)
	var a,b int 
	fmt.Print("\nenter the two number:")
	fmt.Scan(&a,&b)
	switch cal{
	    case 1:
		  fmt.Println("addition of a and b:",a+b)
		case 2:
		  fmt.Println("subraction of a and b",a-b)
		case 3:
		  fmt.Println("multiplication of a and b",a*b)
		case 4:
		  fmt.Println("division of a and b ",a/b)
		case 5:
			fmt.Println("modulus of a and b",a%b)
		default:
			fmt.Println("invalid number")
	   
	}

}