// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timeseries

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Point rappresent a point present in a time serie.
type Point struct {
	X int64
	Y float64
}

// Timeseries is the type for time series.
type Timeseries struct {
	XY           map[int64]float64
	orderedIndex []int64
	sync.Mutex
}

// New creates a timeserie
func New() *Timeseries {

	ts := new(Timeseries)
	ts.XY = make(map[int64]float64, 0)

	return ts
}

// AddNewPoint adds a point to the time serie.
func (ts *Timeseries) AddNewPoint(v float64, x interface{}) error {
	ts.Lock()
	defer ts.Unlock() // unlocks at the end

	switch T := x.(type) {
	case int64:
		ts.XY[T] = v
	case time.Time:
		ts.XY[T.UnixNano()] = v
	case int:
		ts.XY[int64(T)] = v
	default:
		return fmt.Errorf("Adding point not possible")
	}

	return nil
}

// Len returns the length of the timeserie.
func (ts *Timeseries) Len() int {

	ts.Lock()
	l := len(ts.XY)
	ts.Unlock()

	return l

}

func (ts *Timeseries) orderIndex() {

	// locks the time serie untill the end
	ts.Lock()
	defer ts.Unlock()

	// creates string slices to contain indexes
	var indexes []string

	// cicles through the XY map and appends the indexes
	// trasformed in string to the slice
	for n := range ts.XY {
		indexes = append(indexes, fmt.Sprint(n))
	}

	// sorts in ascending order the indexes
	sort.Strings(indexes)

	// clears the existing ts.orderedIndex
	ts.orderedIndex = []int64{}

	// loops through the ts.XY map and prints
	// in ascending order its content converting
	// the ordered indexes back to int64
	for _, s := range indexes {
		i, err := strconv.ParseInt(s, 10, 64) // convers the string in int64
		if err != nil {
			log.Fatal(err)
		}

		ts.orderedIndex = append(ts.orderedIndex, i)
	}

	return
}

// Print prints all the points in the timeserie.
func (ts *Timeseries) Print() {
	ts.orderIndex()

	for n, i := range ts.orderedIndex {
		fmt.Println(n, "\t", i, "\t", ts.XY[i])
	}

	return
}

// PrintFormattedTime prints all the points in the timeserie,
// with times formatted as RFC339
func (ts *Timeseries) PrintFormattedTime() {
	ts.orderIndex()

	for n, i := range ts.orderedIndex {
		fmt.Println(n, "\t", time.Unix(i/1000000000, 0).Format(time.RFC3339), "\t", ts.XY[i])
	}

	return
}

// AddValueToIndex adds the value v at the value present
// at the index specified.
func (ts *Timeseries) AddValueToIndex(v float64, i int64) {
	ts.Lock()
	ts.XY[i] = float64(ts.XY[i]) + v
	ts.Unlock()
}

// AddValueToTime adds the value v at the value present
// at the timestamp specified.
func (ts *Timeseries) AddValueToTime(v float64, t time.Time) {
	ts.Lock()
	i := t.Unix()
	ts.XY[i] = float64(ts.XY[i]) + v
	ts.Unlock()
}

// FindNextPoint returns the next recorded index in the time serie.
// If not available it will return a Point with 0,0
func (ts *Timeseries) FindNextPoint(i int64) Point {

	ts.orderIndex()

	var p = new(Point)

	for n, j := range ts.orderedIndex {
		if j == i {
			if n+1 < len(ts.orderedIndex) {
				p.X = ts.orderedIndex[n+1]
				p.Y = ts.XY[p.X]
				return *p
			}
		}

	}

	return *p
}

// FindPreviousPoint returns the next recorded index in the time serie.
// If not available it will return a Point with 0,0
func (ts *Timeseries) FindPreviousPoint(i int64) Point {

	ts.orderIndex()

	var p = new(Point)

	for n, j := range ts.orderedIndex {
		if j == i {
			if n-1 > 0 {
				p.X = ts.orderedIndex[n-1]
				p.Y = ts.XY[p.X]
				return *p
			}
		}

	}

	return *p
}

// ToSlice creates a slice with the values of the time serie.
func (ts *Timeseries) ToSlice() []float64 {

	var slice []float64

	ts.orderIndex()

	ts.Lock()
	defer ts.Unlock()
	for _, i := range ts.orderedIndex {

		slice = append(slice, ts.XY[i])

	}

	return slice

}

// FromSlice returns a new timeserie created with the data in the slice passed
// as argument.
func FromSlice(start time.Time, step time.Duration, s []float64) (ts *Timeseries, err error) {

	ts = New()

	for n, v := range s {
		err = ts.AddNewPoint(v, start.Add(step*time.Duration(n)))
	}

	return ts, err

}
