package timeseries

import (
	"fmt"
	"sort"
	"time"
)

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

// AddValueToIndex adds the value v at the value present
// at the index specified.
func (ts *Timeseries) AddValueToIndex(v float64, i int64) {
	//ts.Lock()
	//defer ts.Unlock()

	if oldvalue, exists := ts.XY[i]; exists {
		ts.Lock()
		ts.XY[i] = oldvalue + v
		ts.Unlock()
		return
	}
	ts.Lock()
	ts.XY[i] = v
	ts.Unlock()

	return
}

// AddValueToTime adds the value v at the value present
// at the timestamp specified.
func (ts *Timeseries) AddValueToTime(v float64, t time.Time) {
	ts.Lock()
	i := t.Unix()
	ts.XY[i] = float64(ts.XY[i]) + v
	ts.Unlock()
}

// AddNewPointKeepLen adds a point to the time serie and keep a certain number of points.
func (ts *Timeseries) AddNewPointKeepLen(v float64, x interface{}) error {
	ts.orderIndex()
	l := len(ts.XY)
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
	newl := len(ts.XY)

	keys := make([]float64, 0, len(ts.XY))
	for k := range ts.XY {
		keys = append(keys, float64(k))
	}
	sort.Float64s(keys)
	if newl > l {
		delete(ts.XY, int64(keys[0]))
	}
	ts.orderIndex()

	return nil
}
