package repository

var (
	GetColumnUniqueValuesQuery = "SELECT DISTINCT %s AS value FROM %s;"
)
