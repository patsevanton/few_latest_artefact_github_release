package main

import (
	"fmt"
	// "reflect"
	"net/http"
	"io/ioutil"
)

var link_github_release = "https://api.github.com/repos/rabbitmq/erlang-rpm/releases"

func main() {
	
	response, err := http.Get(link_github_release)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	bodyString := string(body)

	fmt.Println(bodyString)

	// fmt.Println(reflect.TypeOf(bodyString))

	// for _, element := range bodyString {
	// 	fmt.Println(reflect.TypeOf(element))
	// }

}
