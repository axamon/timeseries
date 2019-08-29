package timeseries

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// Timeseries is the type for time series.
type Timeseries struct {
	XY           map[int64]float64
	orderedIndex []int64
	sync.Mutex
}

// New creates a timeserie
func New() *Timeseries {

	s := new(Timeseries)
	s.XY = make(map[int64]float64, 0)

	return s
}

// AddNewPoint adds a point to the time serie.
func (ts *Timeseries) AddNewPoint(v float64, x interface{}) error {
	ts.Lock()
	defer ts.Unlock()

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

func (ts *Timeseries) Len() int {

	ts.Lock()
	l := len(ts.XY)
	ts.Unlock()

	return l

}

func (ts *Timeseries) orderIndex() {

	// converts ts indexes to strings
	var indexes []string
	for n := range ts.XY {
		indexes = append(indexes, fmt.Sprint(n))
	}

	// sorts in ascending order the indexes
	sort.Strings(indexes)

	ts.Lock()
	defer ts.Unlock()
	ts.orderedIndex = []int64{}

	// loops through the ts.XY map and prints
	// in ascending order its content converting
	// the ordered indexes back to int64
	for _, s := range indexes {
		i, _ := strconv.ParseInt(s, 10, 64) // convers the string in int64

		ts.orderedIndex = append(ts.orderedIndex, i)
		// fmt.Println(i, ts.XY[int64(i)])
	}

	return
}

func (ts *Timeseries) Print() {
	ts.orderIndex()

	for n, i := range ts.orderedIndex {
		fmt.Println(n, i, "\t", ts.XY[i])
	}

	return
}

func (ts *Timeseries) AddValueToIndex() {

}

func (ts *Timeseries) AddValueToTime() {

}
