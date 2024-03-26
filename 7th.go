package main

import (
	"fmt"
	"strings"
)


func cmp(str1, str2 string) int {
	if strings.EqualFold(string(str1), str2) {
			return 1
		}
	runes := []rune(str1)
	for i := 0; i < len(runes)-1; i++ {
		swapped := make([]rune, len(runes))
		copy(swapped, runes)
		swapped[i], swapped[i+1] = swapped[i+1], swapped[i]

		if strings.EqualFold(string(swapped), str2) {
			return 1
		}
	}

	return 0
}


func main(){
	var work_count int
	fmt.Scan(&work_count)
	login := make([]string, work_count)
	for i:=0; i< work_count; i++{
		fmt.Scan(&login[i])
	}
	/*Duplicate of code is bad, i know it, and can create function, but im busy now :*_*```
*/	var new_work_count int
	fmt.Scan(&new_work_count)
	new_login := make([]string, new_work_count)
	for i:=0; i< new_work_count; i++{
		fmt.Scan(&new_login[i])
	}
	for i:=0; i< new_work_count; i++{
		res := 0 
		for j:=0; j < work_count; j++{
			if len(new_login[i]) == len(login[j]){
				res += cmp(new_login[i], login[j])
			}
		}
		if res != 0{
			fmt.Println(1)
		} else {fmt.Println(0)}
	}

}