package main

func filterByName(records []record, pillName string) []record {
	filteredRecords := make([]record, 0)
	for _, a := range records {
		if a.PillName == pillName {
			filteredRecords = append(filteredRecords, a)
		}
	}
	return filteredRecords
}

func filterByDate(records []record, dueDate string) []record {
	filteredRecords := make([]record, 0)
	for _, a := range records {
		if a.DueDate == dueDate {
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
