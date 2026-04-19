package main

import (
"fmt"
"log"

"src/greetings"
)

func main() {
	//configure logs
	log.SetPrefix("greetings:")
	log.SetFlags(0)


	// fmt.Println("Hello World")
	// fmt.Println("Welcome to Go Tutotial")
	//message,err:= greetings.Greetings("random")
	// message,err:= greetings.Wish()
	// if err!=nil{
	// 	log.Fatal("Non nil error returned")
	// }
	// fmt.Println(message)
	messages,err:=greetings.Greets([]string{
		"Rohan","Vishal","Rahul",
	})
	if err!=nil{
		log.Fatal("error")
	}
	fmt.Println(messages)
	fmt.Println("Exiting from Main.go file...")
}


