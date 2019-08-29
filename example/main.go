package main

import (
	"fmt"
	"time"

	"github.com/axamon/timeseries"
)

func main() {
	ts := timeseries.New()

	//	r := rand.Int() * -1

	ts.AddNewPoint(67, time.Now().Add(time.Duration(-48)*time.Second))
	go ts.AddNewPoint(67.9, time.Now().UnixNano())

	ts.AddNewPoint(56, 1567062855070167570+312)

	fmt.Println(ts, ts.Len())

	ts.Print()
}
