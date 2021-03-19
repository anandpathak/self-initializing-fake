package main

import (
	"self_initializing_fake/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

}
