package main

import (
	"fmt"
	"time"
)

//global scope

var date=400;

func div(x int,y int)(int){
z:=x/y;
return z
}

func main(){
	fmt.Println("Hello World!");
	//local scope
	a :=100;
	//--
	fmt.Println(a);
	var b string="Hello Nishak";
	c:="Variable";
	fmt.Println(c);
	fmt.Println(b);
	// res:=div(time,date)
	// fmt.Println("result is :",res)

	fmt.Println("Enter the value  ")
	fmt.Scanln()
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
	// isDone=false
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
		fmt.Println("your age is equal to 18")
	}else{
		fmt.Println("Your age is",age)
	}
	//not syntax with bool
	isSuccess:=true
	if !isSuccess{
		fmt.Println("Failed")
	}else{
		fmt.Println("Success")
	}
	//or and syntax
	if (isDone && isSuccess){
		fmt.Println("Success")
	}else{
		fmt.Println("Failed")
	}

	value1:=10
	value2:=20
	if(value1==10 || value2==10){
		fmt.Println(value1)
	}

	//switch case syntax
	switch value1{
	case 1:
		fmt.Println("Value is ",value1)
	case 2,3:
		fmt.Println("Value is 2,3")
	case value2:
		fmt.Println("value is ",value2)
	case value1:
		fmt.Println("Value is ",value1)
	default:
		fmt.Println("Default Value")

	}
	//sum of two number
	num1:=10
	num2:=20
	sum:=num1+num2
	fmt.Println("sum of num1 and num 2 ",sum)

	sum1:=add(10,20)
	sum2:=add(value1,value2)
	fmt.Println(sum1)
	fmt.Println(sum2)
	getValues(10,"Nishak")

	//time syntax
	fmt.Println("Welcome to time")
	presentTime:=time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))


}
//return multiple value using function
func getValues(a int,b string)(int, string){
	return a,b

}
func add(num1 int,num2 int)int{
	sum:=num1+num2
	return sum
}