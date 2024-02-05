package main

// import (
// 	"fmt"

// 	turboGenerator "github.com/nats-io/go-nats-streaming/services/generator"
// 	sender "github.com/nats-io/go-nats-streaming/services/sender"
// )

// func main() {
// 	var jsonFile string
// 	var err error
// 	jsonFile, err = turboGenerator.GenerateJSON()
// 	if err != nil {
// 		fmt.Println("1")
// 		panic(err)
// 	}
// 	fmt.Println("1-")
// 	err = sender.Confirm(jsonFile)
// 	if err != nil {
// 		fmt.Println("2")
// 		panic(err)
// 	}
// }
