package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//returns the Greeting message for a specific name
func Greetings(name string) (string,error){
	//handle case for empty name/when name is not given
	if name==""{
		return "",errors.New("empty name")
	}
	//build message to display message to user
message:= fmt.Sprintf(randomFormat(),name)
return message,nil
}

//return random greeting message
func randomFormat() string{
	//define a slice of greeting messages
	formats:=[]string{
		"Hi, Happy Birthday!! %v",
		"Welcome to the team %v!!",
		"Have a great day, %v!!",
	}
	//return any of the messages above by picking random index from 0 to len of formats
	return formats[rand.Intn(len(formats))]
}

//return random name from given set of names
func randomName() string{
	names:= [] string{
		"Rahul",
		"Vishal",
		"Rohan",
	}

	return names[rand.Intn(len(names))]

}

//Wish a random message to random person
func Wish() (string,error){
	message:= fmt.Sprintf(randomFormat(),randomName())
	if message!=""{
		return message,nil
	}
	return "",errors.New("Could not wish!!")
}


func Greet(){

	names:=[]string{
		"Rahul",
		"Vishal",
		"Rohan",
	}
	for i := 0; i < len(names); i++ {
		msg:=fmt.Sprintf(randomFormat(),names[i])
		fmt.Println(msg)
	}

}

func Greets(names[] string) (map[string]string,error){
	messages:= make(map[string]string)

	for _, name:=range names{
		message,err:=Greetings(name)
		if err!=nil{
			fmt.Println("non-null error returned!!!")
			return nil,err
		}
		messages[name] = message

	}
	return messages,nil
}

