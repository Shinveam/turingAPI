package main

import (
	"log"
	"turingAPI/turing"
)

const ApiKey string = "eb58b64b8cd34a68b3c8fe588ded8191"

func TuringRobot() {
	_, err := turing.Robots(ApiKey,
		turing.ReqType(0),
		"董明珠的图片",
	)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	TuringRobot()
}
