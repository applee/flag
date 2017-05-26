package main

import (
	"fmt"
	"log"
	"time"

	"github.com/applee/flag"
)

func main() {
	var (
		config              string
		length              float64
		age                 int
		name                string
		female              bool
		extHobbies          string
		extLinkmanParentDad string
		extLinkmanParentMom string
		duration            time.Duration
	)

	flag.StringVar(&config, "config", "", "help message")
	flag.StringVar(&name, "name", "", "help message")
	flag.IntVar(&age, "age", 0, "help message")
	flag.Float64Var(&length, "length", 0, "help message")
	flag.BoolVar(&female, "female", false, "help message")
	flag.DurationVar(&duration, "duration", time.Second*5, "")
	flag.StringVar(&extHobbies, "ext.hobbies", "", "")
	flag.StringVar(&extLinkmanParentDad, "ext.linkman.parent.dad", "", "")
	flag.StringVar(&extLinkmanParentMom, "ext.linkman.parent.mom", "", "")

	err := flag.Parse()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("length:", length)
	fmt.Println("age:", age)
	fmt.Println("name:", name)
	fmt.Println("female:", female)
	fmt.Println("hobbies:", extHobbies)
	fmt.Println("duration", duration.Seconds())
	fmt.Println("dad", extLinkmanParentDad)
	fmt.Println("mom", extLinkmanParentMom)
}
