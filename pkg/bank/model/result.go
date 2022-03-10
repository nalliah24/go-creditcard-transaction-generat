package model

type Result struct {
	Timestamp     string
	NumberOfTrans int
	Summaries     SummaryList
	Data          TransactionList
}
