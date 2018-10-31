package main

import (
	"fmt"
	"log"
	"turingAPI/turing"
)

const ApiKey string = "eb58b64b8cd34a68b3c8fe588ded8191"

func TuringRobot() {
	resp, err := turing.Robots(ApiKey,
		turing.ReqType(0),
		"你好",
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

func main() {
	TuringRobot()
}
