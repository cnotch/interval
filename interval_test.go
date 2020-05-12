// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package interval

import (
	"reflect"
	"testing"
	"time"
)

func TestThisDay(t *testing.T) {
	tests := []struct {
		name         string
		rtime        time.Time
		wantStarting time.Time
		wantEnding   time.Time
	}{
		{
			"ThisDay",
			time.Date(2017, 11, 21, 16, 33, 54, 100, time.UTC),
			time.Date(2017, 11, 21, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 11, 22, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReferenceTime(tt.rtime).ThisDay()
			if !reflect.DeepEqual(got.From, tt.wantStarting) {
				t.Errorf("ThisDay() got.From = %v, want %v", got.From, tt.wantStarting)
			}
			if !reflect.DeepEqual(got.To, tt.wantEnding) {
				t.Errorf("ThisDay() got.To = %v, want %v", got.To, tt.wantEnding)
			}
		})
	}
}

func TestThisWeek(t *testing.T) {

	tests := []struct {
		name         string
		rtime        time.Time
		wantStarting time.Time
		wantEnding   time.Time
	}{
		{
			"ThisWeek",
			time.Date(2017, 11, 21, 16, 33, 54, 100, time.UTC),
			time.Date(2017, 11, 19, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 11, 26, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReferenceTime(tt.rtime).ThisWeek(time.Sunday)
			if !reflect.DeepEqual(got.From, tt.wantStarting) {
				t.Errorf("ThisWeek() got.From = %v, want %v", got.From, tt.wantStarting)
			}
			if !reflect.DeepEqual(got.To, tt.wantEnding) {
				t.Errorf("ThisWeek() got.To = %v, want %v", got.To, tt.wantEnding)
			}
		})
	}
}

func TestThisMonth(t *testing.T) {

	tests := []struct {
		name         string
		rtime        time.Time
		wantStarting time.Time
		wantEnding   time.Time
	}{
		{
			"ThisMonth",
			time.Date(2017, 11, 21, 16, 33, 54, 100, time.UTC),
			time.Date(2017, 11, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2017, 12, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReferenceTime(tt.rtime).ThisMonth()
			if !reflect.DeepEqual(got.From, tt.wantStarting) {
				t.Errorf("ThisMonth() got.From = %v, want %v", got.From, tt.wantStarting)
			}
			if !reflect.DeepEqual(got.To, tt.wantEnding) {
				t.Errorf("ThisMonth() got.To = %v, want %v", got.To, tt.wantEnding)
			}
		})
	}
}

func TestThisQuarter(t *testing.T) {
	tests := []struct {
		name         string
		rtime        time.Time
		wantStarting time.Time
		wantEnding   time.Time
	}{
		{
			"ThisQuarter",
			time.Date(2017, 11, 21, 16, 33, 54, 100, time.UTC),
			time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReferenceTime(tt.rtime).ThisQuarter()
			if !reflect.DeepEqual(got.From, tt.wantStarting) {
				t.Errorf("ThisQuarter() got.From = %v, want %v", got.From, tt.wantStarting)
			}
			if !reflect.DeepEqual(got.To, tt.wantEnding) {
				t.Errorf("ThisQuarter() got.To = %v, want %v", got.To, tt.wantEnding)
			}
		})
	}
}

func TestThisYear(t *testing.T) {
	tests := []struct {
		name         string
		rtime        time.Time
		wantStarting time.Time
		wantEnding   time.Time
	}{
		{
			"ThisYear",
			time.Date(2017, 11, 21, 16, 33, 54, 100, time.UTC),
			time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReferenceTime(tt.rtime).ThisYear()
			if !reflect.DeepEqual(got.From, tt.wantStarting) {
				t.Errorf("ThisYear() got.From = %v, want %v", got.From, tt.wantStarting)
			}
			if !reflect.DeepEqual(got.To, tt.wantEnding) {
				t.Errorf("ThisYear() got.To = %v, want %v", got.To, tt.wantEnding)
			}
		})
	}
}
