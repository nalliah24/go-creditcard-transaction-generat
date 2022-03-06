package model

type Summary struct {
	NumOfTrans  int     `json:"num_of_trans"`
	TranType    string  `json:"tran_type"`
	TotalAmount float64 `json:"total_amt"`
}

type SummaryList []Summary
