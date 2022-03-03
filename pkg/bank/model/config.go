package model

type Config struct {
	NumOfTrans int      `json:"num_of_trans"`
	TranType   string   `json:"tran_type"`
	MinAmount  float64  `json:"min_amt"`
	MaxAmount  float64  `json:"max_amt"`
	TranCode   []string `json:"tran_code"`
}
