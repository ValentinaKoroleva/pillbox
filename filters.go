package main

import (
	"time"
)

func filterByName(records []record, pillName string) []record {
	filteredRecords := make([]record, 0)
	for _, a := range records {
		if a.PillName == pillName {
			filteredRecords = append(filteredRecords, a)
		}
	}
	return filteredRecords
}

func filterByDate(records []record, dueDate time.Time) []record {
	filteredRecords := make([]record, 0)
	for _, a := range records {
		if a.DueDate == dueDate {
			filteredRecords = append(filteredRecords, a)
		}
	}
	return filteredRecords
}

func filterByInterval(records []record, fromDate time.Time, toDate time.Time) []record {
	filteredRecords := make([]record, 0)
	for _, a := range records {
		if a.DueDate.Before(toDate.AddDate(0, 0, 1)) && a.DueDate.After(fromDate.AddDate(0, 0, -1)) {
			filteredRecords = append(filteredRecords, a)
		}
	}
	return filteredRecords
}

func filterByStatus(records []record, status bool) []record {
	filteredRecords := make([]record, 0)
	for _, a := range records {
		if a.Status == status {
			filteredRecords = append(filteredRecords, a)
		}
	}
	return filteredRecords
}
