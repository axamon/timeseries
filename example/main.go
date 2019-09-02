package main

import (
	"fmt"
	"log"
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
		err := ts.AddNewPoint(float64(rand.Intn(50)), now)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(i)
	}

	//	r := rand.Int() * -1

	err := ts.AddNewPoint(67, time.Now().Add(time.Duration(+48)*time.Second))
	if err != nil {
		log.Println(err)
	}
	go ts.AddNewPoint(67.9, time.Now().UnixNano())

	err = ts.AddNewPoint(56, 1567062855070167570+312)
	if err != nil {
		log.Println(err)
	}
	err = ts.AddNewPoint(30, 2567062855070167570)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(ts, ts.Len())

	ts.Print()

	fmt.Println(ts.FindNextPoint(1567062855070167570 + 312))
	fmt.Println(ts.FindNextPoint(2567062855070167570))
	fmt.Println(ts.FindPreviousPoint(1567062855070167570 + 312))

	var slice []float64

	for i := 0; i < 1000; i++ {

		slice = append(slice, float64(rand.Intn(1000)))
	}

	serie, err := timeseries.FromSlice(time.Now(), time.Duration(5*time.Second), slice)
	if err != nil {
		log.Println(err)
	}

	serie.Print()
	serie.PrintFormattedTime()

	fmt.Println(serie.FirstX(), serie.LastX())
}
