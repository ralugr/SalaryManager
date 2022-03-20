package model

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNextPayDate(t *testing.T) {
	var tests = []struct {
		payDay   int
		from     time.Time
		expected time.Time
	}{
		{18, Date(2022, 3, 17), Date(2022, 3, 18)},
		{17, Date(2022, 3, 19), Date(2022, 4, 15)},
		{1, Date(2022, 5, 30), Date(2022, 6, 1)},
		{27, Date(2022, 7, 30), Date(2022, 8, 26)},
		{15, Date(2022, 3, 15), Date(2022, 4, 15)},
		{20, Date(2022, 3, 19), Date(2022, 4, 20)},
		{29, Date(2022, 2, 19), Date(2022, 3, 1)},
	}
	for _, data := range tests {
		testname := fmt.Sprintf("Pay day:%d,Current date: %v", data.payDay, data.from)
		t.Run(testname, func(t *testing.T) {
			ans := NextPayDate(data.payDay, data.from)
			if ans != data.expected {
				t.Errorf("Actual %v, Expected %v", ans, data.expected)
			}
		})
	}
}

func TestDaysTillNextSalary(t *testing.T) {
	var tests = []struct {
		payDay   int
		from     time.Time
		expected int
	}{
		{18, Date(2022, 3, 17), 1},
		{17, Date(2022, 3, 19), 27},
		{1, Date(2022, 5, 30), 2},
		{27, Date(2022, 7, 30), 27},
		{15, Date(2022, 3, 15), 31},
	}
	for _, data := range tests {
		testname := fmt.Sprintf("Pay day:%d,Current date: %v", data.payDay, data.from)
		t.Run(testname, func(t *testing.T) {
			ans := DaysTillNextSalary(data.payDay, data.from)
			if ans != data.expected {
				t.Errorf("Actual %v, Expected %d", ans, data.expected)
			}
		})
	}
}

func TestAnnualPaySchedule(t *testing.T) {
	var tests = []struct {
		payDay   int
		from     time.Time
		expected []time.Time
	}{
		{18, Date(2022, 3, 17), []time.Time{
			Date(2022, 3, 18),
			Date(2022, 4, 18),
			Date(2022, 5, 18),
			Date(2022, 6, 17),
			Date(2022, 7, 18),
			Date(2022, 8, 18),
			Date(2022, 9, 16),
			Date(2022, 10, 18),
			Date(2022, 11, 18),
			Date(2022, 12, 16),
		}},
	}
	for _, data := range tests {
		testname := fmt.Sprintf("Pay day:%d,Current date: %v", data.payDay, data.from)
		t.Run(testname, func(t *testing.T) {
			ans := AnnualPaySchedule(data.payDay, data.from)
			if !reflect.DeepEqual(ans, data.expected) {
				t.Errorf("Actual %v, Expected %v", ans, data.expected)
			}
		})
	}
}
