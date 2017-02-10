package main

import (
	"fmt"
	"log"

	"github.com/applee/flag"
)

func main() {
	var (
		config     string
		length     float64
		age        int
		name       string
		female     bool
		extHobbies string
	)

	flag.StringVar(&config, "config", "", "help message")
	flag.StringVar(&name, "name", "", "help message")
	flag.IntVar(&age, "age", 0, "help message")
	flag.Float64Var(&length, "length", 0, "help message")
	flag.BoolVar(&female, "female", false, "help message")
	flag.StringVar(&extHobbies, "ext.hobbies", "", "")

	err := flag.Parse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("length:", length)
	fmt.Println("age:", age)
	fmt.Println("name:", name)
	fmt.Println("female:", female)
	fmt.Println("hobies:", extHobbies)
}
