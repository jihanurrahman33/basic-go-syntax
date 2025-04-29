package main

import (
	"fmt"
	"sort"
)
func main(){
	fmt.Println("Welcome to slices syntax")
	//slice syntax
	var fruitList=[]string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitList is %T\n",fruitList)
	//adding item to slice
	fruitList = append(fruitList, "Potato","Banana")
	fmt.Println(fruitList)
	//slicing of slice
	fruitList =append(fruitList[1:] )
	fmt.Println(fruitList)
	fruitList =append(fruitList[:3] )
	fmt.Println(fruitList)
	fruitList =append(fruitList[1:3] )
	fmt.Println(fruitList)

	//create something using make function
	highScore:=make([]int,4)
	highScore[0]=234
	highScore[1]=744
	highScore[2]=699
	highScore[3]=233
	/*
	but this syntax not works as value is 4
	highscore[4]=555
	*/
	//append works perfectly because it reallocate the memory
	highScore=append(highScore, 999,333,666)
	fmt.Println(highScore)
	//sorting int
	sort.Ints(highScore)
	fmt.Println(highScore)
	//return bool value
	fmt.Println(sort.IntsAreSorted(highScore))
}