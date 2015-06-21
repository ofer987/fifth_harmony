package main

import "fmt"

import "net/http"

func main() {
	fmt.Println("hello world")

	sendHttpGet("www.cnet.com")
	// var givenName string
	// givenName = askMeMyName()
	//
	// fmt.Printf("Your name is %s\n", givenName)
}

func askMeMyName() string {
	var givenName string
	fmt.Println("What is your name?")
	fmt.Scanf("%s", &givenName)

	return givenName
}

func sendHttpGet(url string) {
	var response http.Response
	var err error

	response, err := http.Get(url)
}
