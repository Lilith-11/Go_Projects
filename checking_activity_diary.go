package main
import ("fmt"
"time"
"bufio"
"os")
func main(){
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println("enter what are you doing")
	activity,err:=reader.ReadString('\n')
	if err!=nil{
		return
	}
	timestamp:=time.Now().Format("2006-01-02 15:04:05")
	
    entry := fmt.Sprintf("[%s]%s",timestamp,activity)
	file,err:=os.OpenFile("diary1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
	if err!=nil{
		fmt.Println("there is error in opening file")
		return
	}
	defer file.Close()
	file.WriteString(entry)
	fmt.Println("activity added")
}