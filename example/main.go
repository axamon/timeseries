package main

import (
	"fmt"
	"time"

	"github.com/axamon/timeseries"
)

func main() {
	ts := timeseries.New()

	//	r := rand.Int() * -1

	ts.AddNewPoint(67, time.Now().Add(time.Duration(+48)*time.Second))
	go ts.AddNewPoint(67.9, time.Now().UnixNano())

	ts.AddNewPoint(56, 1567062855070167570+312)
	ts.AddNewPoint(30, 2567062855070167570)

	fmt.Println(ts, ts.Len())

	ts.Print()

	fmt.Println(ts.FindNextPoint(1567062855070167570 + 312))
	fmt.Println(ts.FindNextPoint(2567062855070167570))
	fmt.Println(ts.FindPreviousPoint(1567062855070167570 + 312))
}
