package main

import (
	"fmt"
	. "github.com/Sa2Knight/go-zaim"
	"os"
)

func main() {
	fmt.Println(NewClient(
		os.Getenv("ZAIM_CUSTOMER_KEY"),
		os.Getenv("ZAIM_CUStOMER_SECRET"),
		os.Getenv("ZAIM_TOKEN"),
		os.Getenv("ZAIM_SECRET"),
	))
}