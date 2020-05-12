// Copyright (c) 2019,CAO HONGJU. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package interval

import "time"

// Interval presentat time interval [From, To).
type Interval struct {
	From time.Time
	To   time.Time
}

// Offset offsets time interval with the given years, months and days.
func (tv Interval) Offset(years, months, days int) Interval {
	tv.From = tv.From.AddDate(years, months, days)
	tv.To = tv.To.AddDate(years, months, days)
	return tv
}

// OffsetDay offsets time interval with the given days.
func (tv Interval) OffsetDay(days int) Interval {
	return tv.Offset(0, 0, days)
}

// OffsetMonth offsets time interval with the given months.
func (tv Interval) OffsetMonth(months int) Interval {
	return tv.Offset(0, months, 0)
}

// OffsetYear offsets time interval with the given years.
func (tv Interval) OffsetYear(years int) Interval {
	return tv.Offset(years, 0, 0)
}

// Today returns the time interval for today.
func Today() Interval {
	return ReferenceTime(time.Now()).ThisDay()
}

// ThisWeek returns the time interval for this week.
func ThisWeek(firstDay time.Weekday) Interval {
	return ReferenceTime(time.Now()).ThisWeek(firstDay)
}

// ThisMonth returns the time interval for this month.
func ThisMonth() Interval {
	return ReferenceTime(time.Now()).ThisMonth()
}

// ThisQuarter returns the time interval for this quarter.
func ThisQuarter() Interval {
	return ReferenceTime(time.Now()).ThisQuarter()
}

// ThisYear returns the time interval for this year.
func ThisYear() Interval {
	return ReferenceTime(time.Now()).ThisYear()
}

// ReferenceTime presentat refenced time of time interval.
type ReferenceTime time.Time

// ThisDay returns the time interval for the day of reference time.
func (rt ReferenceTime) ThisDay() Interval {
	t := time.Time(rt)
	year, month, day := t.Date()

	from := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	to := from.AddDate(0, 0, 1)

	return Interval{from, to}
}

// ThisWeek returns the time interval for the week of reference time.
func (rt ReferenceTime) ThisWeek(firstDay time.Weekday) Interval {
	t := time.Time(rt)
	year, month, day := t.Date()

	from := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	intervalDays := int(from.Weekday() - firstDay)
	if intervalDays < 0 {
		intervalDays = 7 + intervalDays
	}
	if intervalDays != 0 {
		from = from.AddDate(0, 0, -intervalDays)
	}
	to := from.AddDate(0, 0, 7)

	return Interval{from, to}
}

// ThisMonth returns the time interval for the month of reference time.
func (rt ReferenceTime) ThisMonth() Interval {
	t := time.Time(rt)
	year, month, _ := t.Date()

	from := time.Date(year, month, 1, 0, 0, 0, 0, t.Location())
	to := from.AddDate(0, 1, 0)

	return Interval{from, to}
}

const monthsPerQuarter = 3

// ThisQuarter returns the time interval for the quarter of reference time.
func (rt ReferenceTime) ThisQuarter() Interval {
	t := time.Time(rt)
	year, month, _ := t.Date()

	from := time.Date(year, ((month-1)/monthsPerQuarter)*monthsPerQuarter+1, 1, 0, 0, 0, 0, t.Location())
	to := from.AddDate(0, monthsPerQuarter, 0)

	return Interval{from, to}
}

// ThisYear returns the time interval for the year of reference time.
func (rt ReferenceTime) ThisYear() Interval {
	t := time.Time(rt)
	year, _, _ := t.Date()

	from := time.Date(year, 1, 1, 0, 0, 0, 0, t.Location())
	to := from.AddDate(1, 0, 0)

	return Interval{from, to}
}
