package main

import "fmt"
func main(){
	fmt.Println("Hello World!");
	a :=100;
	fmt.Println(a);
	var b string="Hello Nishak";
	c:="Variable";
	fmt.Println(c);
	fmt.Println(b);
	/*
	Basic Data Types:
	-int 
	-float32
	-bool
	-string
	*/

	var x int =10;
	fmt.Println(x);
	//syntax for boolean value
	isDone:=true
	isDone=false
	//fmt println package usage
	fmt.Println(isDone)
	//syntax constant in golang
	const pi=3.1416
	fmt.Println(pi)

	//if else syntax
	age:=20
	if age> 18{
		fmt.Println("Your age is greater 18")
	}else if age<18{
		fmt.Println("your age is less than 18")
	}else if age==18{
		fmt.Println("your age is equla to 18")
	}else{
		fmt.Println("Your age is",age)
	}


}