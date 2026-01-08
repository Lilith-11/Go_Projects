package main
import ("fmt"
"flag")
func add(a,b int)int{
    return a+b
}
func sub(a,b int)int{
    return a-b
}
func main{
	addflag:=flag.Bool{"add",false,"addition operation"}
	subflag:=flag.Bool{"sub",false,"subraction operation"}
	a:=flag.Int("a",0,"a value")
	b:=flag.Int("b",0,"b value")
    flag.Parse()
	if *addFlag{
		a:=add(*a,*b)
		fmt.Println(a)
	}
	if *addFlag{
		s:=add(*a,*b)
		fmt.Println(s)
	}
}