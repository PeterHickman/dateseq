package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var starting_time time.Time
var ending_time time.Time
var days time.Duration

func init() {
	flag.Parse()

	start_index := 0
	day_index := 1
	end_index := 2

	switch len(flag.Args()) {
	case 2:
		day_index = -1
		end_index = 1
	case 3:
		// This is ok
	default:
		println("Wrong arguments")
		os.Exit(1)
	}

	s, err := time.Parse(time.DateOnly, flag.Arg(start_index))
	if err != nil {
		fmt.Printf("Could not parse %s as a date: %s\n", flag.Arg(start_index), err)
		os.Exit(2)
	}

	d := int64(1)
	if day_index != -1 {
		d, err = strconv.ParseInt(flag.Arg(day_index), 10, 64)
		if err != nil {
			fmt.Printf("Could not parse %s as a number: %s\n", flag.Arg(day_index), err)
			os.Exit(2)
		}

		if d < 1 {
			fmt.Printf("The incremenet %d needs to be 1 or greater\n", d)
			os.Exit(2)
		}
	}

	e, err := time.Parse(time.DateOnly, flag.Arg(end_index))
	if err != nil {
		fmt.Printf("Could not parse %s as a date: %s\n", flag.Arg(end_index), err)
		os.Exit(2)
	}

	starting_time = s
	ending_time = e
	days, _ = time.ParseDuration(fmt.Sprintf("%dh", d*24))
}

func main() {
	for {
		fmt.Println(starting_time.Format(time.DateOnly))
		starting_time = starting_time.Add(days)

		if starting_time.Compare(ending_time) == 1 {
			break
		}
	}
}
