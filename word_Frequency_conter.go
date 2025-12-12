package main
import ("fmt"
        "strings")
func main(){
	words:="I am Lilith. I am a go developer"
	countwords:=make(map[string]int)
	text:=strings.Fields(strings.ToLower(words))
	for _,txt:= range text{
		txt=strings.Trim(txt,".")
		countwords[txt]++
	}
	fmt.Println(countwords)
	

}