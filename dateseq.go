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
var direction int

func usage() {
	fmt.Println("Like seq but for dates")
	fmt.Println()
	fmt.Println("dateseq 2024-01-01 2024-01-10")
	fmt.Println("  will list all the dates from 2024-01-01 to 2024-01-10 inclusive")
	fmt.Println()
	fmt.Println("dateseq 2024-01-10 2024-01-01")
	fmt.Println("  will list all the dates from 2024-01-01 to 2024-01-10 in the reverse order")
	fmt.Println()
	fmt.Println("dateseq 2024-01-01 7 2024-12-31")
	fmt.Println("  will list every 7th date from 2024-01-01 up to 2024-12-31 but")
	fmt.Println("  may not include the final date")
	fmt.Println()
	fmt.Println("dateseq 2024-12-31 7 2024-01-01")
	fmt.Println("  will list every 7th date from 2024-12-31 down to 2024-01-01 but")
	fmt.Println("  may not include the first date")
	fmt.Println()
	fmt.Println("Basically the step value is always a positive value")

	os.Exit(1)
}

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
		usage()
	}

	s, err := time.Parse(time.DateOnly, flag.Arg(start_index))
	if err != nil {
		fmt.Printf("Could not parse %s as a date: %s\n", flag.Arg(start_index), err)
		os.Exit(2)
	}

	e, err := time.Parse(time.DateOnly, flag.Arg(end_index))
	if err != nil {
		fmt.Printf("Could not parse %s as a date: %s\n", flag.Arg(end_index), err)
		os.Exit(2)
	}

	if s.Compare(e) == 1 {
		direction = int(-1)
	} else {
		direction = int(1)
	}

	d := int64(1)
	if day_index != -1 {
		d, err = strconv.ParseInt(flag.Arg(day_index), 10, 64)
		if err != nil {
			fmt.Printf("Could not parse %s as a number: %s\n", flag.Arg(day_index), err)
			os.Exit(2)
		}
		if d < 1 {
			fmt.Println("The step value needs to be positive")
			os.Exit(2)
		}
	}

	starting_time = s
	ending_time = e
	days, _ = time.ParseDuration(fmt.Sprintf("%dh", int64(direction)*d*24))
}

func main() {
	for {
		fmt.Println(starting_time.Format(time.DateOnly))
		starting_time = starting_time.Add(days)

		if starting_time.Compare(ending_time) == direction {
			break
		}
	}
}
