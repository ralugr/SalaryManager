package model

import (
	"time"
)

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func NextPayDate(payDay int, currentDate time.Time) time.Time {
	salaryDate := currentDate
	if currentDate.Day() >= payDay {
		salaryDate = Date(salaryDate.Year(), int(salaryDate.Month())+1, payDay)
	} else {
		salaryDate = Date(currentDate.Year(), int(currentDate.Month()), payDay)
	}
	salaryDate = weekendHandler(salaryDate)

	if salaryDate.Before(currentDate) {
		salaryDate = Date(salaryDate.Year(), int(salaryDate.Month())+1, payDay)
		salaryDate = weekendHandler(salaryDate)
	}

	return salaryDate
}

func weekendHandler(salaryDate time.Time) time.Time {
	switch salaryDate.Weekday() {
	case time.Saturday:
		salaryDate = salaryDate.AddDate(0, 0, -1)
		return salaryDate
	case time.Sunday:
		salaryDate = salaryDate.AddDate(0, 0, -2)
		return salaryDate
	default:
		return salaryDate
	}
}

func DaysTillNextSalary(payDay int, currentDate time.Time) int {
	return int(NextPayDate(payDay, currentDate).Sub(currentDate).Hours()) / 24
}

func AnnualPaySchedule(payDay int, currentDate time.Time) []time.Time {
	var scd []time.Time
	nxt := NextPayDate(payDay, currentDate)
	scd = append(scd, nxt)

	for nxt.Month() < 12 && currentDate.Year() == nxt.Year() {
		nxt = NextPayDate(payDay, Date(nxt.Year(), int(nxt.Month()), nxt.Day()+1))
		scd = append(scd, nxt)
	}

	return scd
}
