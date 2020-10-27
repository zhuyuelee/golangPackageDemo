package main

import (
	"fmt"
	"net/http"
)

func main() {
	client := &http.Client{}
	url := "http://www.qq.com"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	response, _ := client.Do(reqest)
	fmt.Println(response.Status)

	fmt.Println(reqest.URL)
}
