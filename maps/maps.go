package main

import "fmt"
func main(){
	fmt.Println("Maps in golang")
	languages:=make(map[string]string)
	languages["js"]="Javascript"
	languages["dart"]="dart"
	languages["flutter"]="flutter"
	languages["py"]="python"

	fmt.Println(languages)

	fmt.Println(languages["js"])

	//deleting value from map

	delete(languages,"py")
	fmt.Println(languages)

	//loops in golang
	
	for key,value:=range languages{
		fmt.Printf("key %v,value %v",key,value)

	}
}