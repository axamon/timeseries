package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/axamon/timeseries"
)

func main() {

	ts := timeseries.New()

	now := time.Now()

	for i := 0; i <= 100; i++ {
		now = now.Add(time.Minute * time.Duration(5))
		fmt.Println(ts.ToSlice())
		ts.AddNewPoint(float64(rand.Intn(50)), now)
		fmt.Println(i)
	}

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
