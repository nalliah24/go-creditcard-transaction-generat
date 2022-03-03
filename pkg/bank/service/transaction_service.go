// Transaction Serive is tool generate sample credit card transactions and output to json file.
// The sample input is fetched from a config file and user can alter the file for desired output
// This tool uses float64 instead of Decimal for sample data, but real life calculation use Decimal pkg
package service

import (
	"errors"
	"math/rand"
	"time"

	m "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
	// d "github.com/shopspring/decimal"
)

const (
	min_amt = 100
	max_amt = 5000
)

// Generate transaction based on config file and returns transactions or error
func GenerateAllTransactions(cfgs []m.Config) ([]m.Transaction, error) {
	// Gen transactions based on config input. Add all trans to a slice, so we can commit to json in one go
	// slice with length is better formant, we know the max rows based on the config file
	// but we CANNOT use it. appends at the end of array
	// var maxRows int
	// for _, cfg := range cfgs {
	// 	maxRows += cfg.NumOfTrans
	// }

	finTrans := []m.Transaction{}

	for _, cfg := range cfgs {
		trans, err := generateTransaction(cfg)
		if err != nil {
			return nil, errors.New("error generating transactions. " + err.Error())
		}
		finTrans = append(finTrans, trans...)
	}

	// Since trans are created based on config array. when more than one array found,
	// trans are indexed starting with one.
	// so need re-index the transactions
	for i := 0; i < len(finTrans); i++ {
		finTrans[i].Id = i + 1
	}

	return finTrans, nil
}

// truncate a float to two levels of precision
func truncate(num float64) float64 {
	return float64(int(num*100)) / 100
}

func generateTransaction(c m.Config) ([]m.Transaction, error) {
	rand.Seed(time.Now().UnixNano()) // In Windows, when run time does not update rightaway. takes 10ms. So will not be accurate in loops
	if c.NumOfTrans <= 0 {
		return nil, errors.New("config error: number of transactions required")
	}
	if c.TranType == "" {
		return nil, errors.New("config error: trancation code (CR|DR) required")
	}

	// set default values, if not submitted
	if c.MinAmount <= 0 {
		c.MinAmount = min_amt
	}
	if c.MaxAmount <= 0 {
		c.MaxAmount = max_amt
	}
	if len(c.TranCode) <= 0 {
		c.TranCode = append(c.TranCode, "Other")
	}

	trs := []m.Transaction{}
	for i := 1; i <= c.NumOfTrans; i++ {
		tcode := c.TranCode[rand.Intn(len(c.TranCode))]
		amt := rand.Float64() * (c.MaxAmount - c.MinAmount)
		t := m.Transaction{
			Id:         i,
			CustomerId: rand.Intn(100) + 1,
			TranType:   c.TranType,
			TranCode:   tcode,
			TranDate:   time.Now().Format("2006-01-02 15:04:05"),
			Amount:     truncate(amt),
		}
		trs = append(trs, t)
	}

	return trs, nil
}

// Prepare summary receives all transactions and returns the summary or error
func PrepareSummary(trs []m.Transaction) ([]m.Summary, error) {
	typs := []string{"CR", "DR"}
	var summaries = []m.Summary{}

	for _, typ := range typs {
		summaries = append(summaries, prepareSummaryForType(typ, trs))
	}

	return summaries, nil
}

func prepareSummaryForType(typ string, trs []m.Transaction) (m.Summary) {
	var s = m.Summary{}
	s.TranType = typ
	for _, tr := range trs {
		if tr.TranType == typ {
			s.TotalAmount = s.TotalAmount + tr.Amount
			s.NumOfTrans++
		}
	}
	return s
}
