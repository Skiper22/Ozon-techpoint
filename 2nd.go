package main
import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	arr := make([]int, t)
	for i:=0; i< t; i++{
		var count, p int
		fmt.Scan(&count, &p)
		price := 0
		for j:=0; j < count; j++{
			fmt.Scan(&price)
			arr[i]+= price*p%100
		}
	}
	for i:=0; i < t; i++{
		fmt.Printf("%.2f", float32(arr[i])/100)
		fmt.Println()
	}
}