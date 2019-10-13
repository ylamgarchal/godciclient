package main

import (
	"fmt"

	"github.com/ylamgarchal/godciclient/dci"
)

func main() {
	var dciAPI = dci.GetClient(
		"http://127.0.0.1:5000/api/v1",
		"admin",
		"admin")
	mytopic, err := dciAPI.GetTopicByName("RHEL-8.1")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Topic id: %s\n", mytopic.ID)
	fmt.Printf("Topic name: %s\n", mytopic.Name)
}
