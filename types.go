package timeseries

import (
	"sync"
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
	firstX       int64
	lastX        int64
	firstY       float64
	lastY        float64
	sync.Mutex
}