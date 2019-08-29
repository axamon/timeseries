// Copyright 2019 Alberto Bregliano. All rights reserved.

// Use of this source code is governed by a BSD-style

// license that can be found in the LICENSE file.

package timeseries_test

import (
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ts.AddNewPoint(tt.args.v, tt.args.x)
			ok := tt.expected != ts.XY[tt.args.x.(int64)]
			if ok != tt.wantErr {
				t.Errorf("Timeseries.AddNewPoint() wantErr %v, expected %v, got %v", tt.wantErr, tt.expected, ts.XY[tt.args.x.(int64)])
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
