package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"net/http"
)
//goroutine for Avengers
func avengers() {
	response, err := http.Get("http://www.mocky.io/v2/5ecfd5dc3200006200e3d64b")
	if err != nil {
		fmt.Printf("HTTP request failed with the error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
//goroutine for Antiheros
func antiheros() {
	response, err := http.Get("http://www.mocky.io/v2/5ecfd630320000f1aee3d64d")
	if err != nil {
		fmt.Printf("HTTP request failed with the error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))
	}
}
//goroutine for mutants
func mutants() {
	response, err := http.Get("http://www.mocky.io/v2/5ecfd6473200009dc1e3d64e")
	if err != nil {
		fmt.Printf("HTTP request failed with the error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		fmt.Println(string(data))
	}
}
func main() {

	go avengers() //calling Avengers goroutine
	time.Sleep(1 * time.Second)
	go antiheros()  //calling Antiheros goroutine
	time.Sleep(1 * time.Second)
	go mutants() //calling mutants goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("main function executing")



}

