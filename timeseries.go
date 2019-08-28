package timeseries

import (
	"fmt"
	"sync"
	"time"
)

// Timeseries is the type for time series.
type Timeseries struct {
	XY map[int64]float64
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

	switch T := x.(type) {
	case int64:
		ts.Lock()
		ts.XY[T] = v
		ts.Unlock()
	case time.Time:
		ts.Lock()
		ts.XY[T.UnixNano()] = v
		ts.Unlock()
	case int:
		ts.Lock()
		ts.XY[int64(T)] = v
		ts.Unlock()
	default:
		return fmt.Errorf("Adding point not possible")
	}

	return nil
}

func (ts *Timeseries) AddValueToIndex() {

}

func (ts *Timeseries) AddValueToTime() {

}
