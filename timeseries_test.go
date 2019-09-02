// Copyright 2019 Alberto Bregliano. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package timeseries_test

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	"github.com/axamon/timeseries"
)

func TestTimeseries_AddNewPoint(t *testing.T) {
	var ts = timeseries.New()

	type args struct {
		v float64
		x interface{}
	}
	tests := []struct {
		name     string
		ts       *timeseries.Timeseries
		args     args
		expected float64
		wantErr  bool
	}{
		{"Primo", ts, args{v: 50, x: int64(10)}, 50, false},
		{"Secondo", ts, args{v: 50, x: int64(8)}, 40, true},
		//{"no", ts, args{v: 50, x: "otto"}, 40, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.ts.AddNewPoint(tt.args.v, tt.args.x)
			if err != nil {
				t.Errorf("Timeseries.AddNewPoint() wantErr %v, expected %v, got error: %v", tt.wantErr, tt.expected, err)
			}
			ok := tt.expected != ts.XY[tt.args.x.(int64)]
			if ok != tt.wantErr {
				t.Errorf("Timeseries.AddNewPoint() wantErr %v, expected %v, got error: %v", tt.wantErr, tt.expected, err)
			}
		})
	}
}

func TestTimeseries_AddValueToIndex(t *testing.T) {
	var ts = timeseries.New()

	type args struct {
		v float64
		x int64
	}
	tests := []struct {
		name     string
		ts       *timeseries.Timeseries
		args     args
		expected float64
		wantErr  bool
	}{
		{"Primo", ts, args{v: 50, x: int64(10)}, 50, false},
		{"Secondo", ts, args{v: 50, x: int64(8)}, 40, true},
		{"terzo", ts, args{v: 50, x: int64(10)}, 100, false},
		{"quarto", ts, args{v: -70, x: int64(10)}, 30, false},
		{"quarto", ts, args{v: -40, x: int64(10)}, 10, true},
		{"quinto", ts, args{v: 0, x: int64(10)}, -10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ts.AddValueToIndex(tt.args.v, tt.args.x)
			ok := tt.expected != ts.XY[tt.args.x]
			if ok != tt.wantErr {
				t.Errorf("Timeseries.AddNewPoint() wantErr %v, expected %v, got %v", tt.wantErr, tt.expected, ts.XY[tt.args.x])
			}
		})
	}
}

func TestTimeseries_Len(t *testing.T) {
	var ts = timeseries.New()
	var now = time.Now()

	ts.AddNewPoint(5, 23)
	ts.AddNewPoint(5, now)
	ts.AddNewPoint(6, now.Add(4*time.Second))

	tests := []struct {
		name string
		ts   *timeseries.Timeseries
		want int
	}{
		{"primo", ts, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ts.Len(); got != tt.want {
				t.Errorf("Timeseries.Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeseries_ToSlice(t *testing.T) {
	var ts = timeseries.New()

	tests := []struct {
		name string
		want []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ts.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Timeseries.ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkAddNewPoint(b *testing.B) {
	ts := timeseries.New()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		ts.AddNewPoint(float64(rand.Intn(1000)), time.Now().Add(time.Duration(n)*time.Hour))
	}
}
