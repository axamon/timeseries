package main

import (
	"fmt"
	"time"

	"github.com/axamon/timeseries"
)

func main() {
	ts := timeseries.New()

	ts.AddNewPoint(67, time.Now())
	ts.AddNewPoint(67.9, time.Now().UnixNano())

	ts.AddNewPoint(56, 1567030959592611)
	fmt.Println(ts)
}
