package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/applee/flag"
)

type hobbies []string

func (p *hobbies) String() string {
	return fmt.Sprint(*p)
}

func (p *hobbies) Set(value string) error {
	var err error
	b := []byte(value)
	if IsJSON(b) {
		err = json.Unmarshal(b, p)
	} else {
		*p = append(*p, value)
	}
	return err
}

func IsJSON(b []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(b, &js) == nil
}

func main() {
	var (
		config              string
		length              float64
		age                 int
		name                string
		female              bool
		extHobbies          hobbies
		extLinkmanParentDad string
		extLinkmanParentMom string
		duration            time.Duration
	)

	flag.StringVar(&config, "config", "config.toml", "help message")
	flag.StringVar(&name, "name", "", "help message")
	flag.IntVar(&age, "age", 0, "help message")
	flag.Float64Var(&length, "length", 0, "help message")
	flag.BoolVar(&female, "female", false, "help message")
	flag.DurationVar(&duration, "duration", time.Second*5, "")
	flag.Var(&extHobbies, "ext.hobbies", "")
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
