package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Requests")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:5000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	fmt.Println(content)
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println(responseString.String())

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:5000/post"

	data := strings.NewReader(`
	{
		"name" : "Sai",
		"age" : "22",
		"country" : "India"
	}
	`)

	response, err := http.Post(myurl, "application/json", data)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println("Post Response: ", response)

	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Content: ", content)
	fmt.Println("String COntent: ", string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:5000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "sai")
	data.Add("lastname", "kokate")
	data.Add("email", "sk@go.dev")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))

}
