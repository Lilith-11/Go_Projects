package main 
import "fmt"
type student struct{
	name string
	score int
}
func updatescore(m *int,add int){
	*m =*m+add
}
func main(){
    s:=student{name:"lilith",score:80}
	fmt.Println("before score",s.score)
	updatescore(&s.score,10)
	fmt.Println("after score",s.score)
}
