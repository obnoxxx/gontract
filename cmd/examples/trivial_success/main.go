package main

import "github.com/gontract/gontract"

func exampleFunc() string {

	gontract.PreCondition(true, "trivially true")

	gontract.PostCondition(true, "trivially true")
	return "OK"
}

func main() {
	message := exampleFunc()

	println(message)

}
