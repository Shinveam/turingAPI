package main

import (
	"log"
	"turingAPI/turing"
)

const ApiKey string = "Your APIKEY"

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
