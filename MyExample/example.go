package main

import (
	"fmt"
	"log"
	"turingAPI/turing"
)

const ApiKey string = "***your apikey***"

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
