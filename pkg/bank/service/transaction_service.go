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
// Config.json is an array and should contain number of rows for each type of transaction. CR|DR
func GenerateAllTransactions(cfgs []m.Config) (m.TransactionList, error) {
	// Add all trans to a slice, so we can write to json in one go
	// Using unbuffered slice. It is possible to use buffered since we know the total rows.
	// But then we need to manipulate each pos in array.
	// Slice append is better, it fills from the 0 pos
	trans := m.TransactionList{}

	for _, cfg := range cfgs {
		ts := generateTransaction(cfg)
		if ts.Err != nil {
			return nil, ts.Err
		}
		trans = append(trans, ts.Transactions...)
	}

	// When more than one type of (CR|DR) found, each array starts with index 1
	// So need re-index the id to start 1 to total num of transactions
	trans.ReIndex()

	return trans, nil
}

// truncate a float to two levels of precision
func truncate(num float64) float64 {
	return float64(int(num*100)) / 100
}

// core method to create transactions based on config file
func generateTransaction(c m.Config) m.TransactionResult {
	rand.Seed(time.Now().UnixNano()) // In Windows, when run time does not update rightaway. takes 10ms. So will not be accurate in loops

	if c.NumOfTrans <= 0 {
		return m.TransactionResult{ Transactions: nil, Err: errors.New("config error: number of transactions required")}
	}
	if c.TranType == "" {
		return m.TransactionResult{ Transactions: nil, Err: errors.New("config error: trancation code (CR|DR) required")}
	}

	// Set default values, if not submitted
	if c.MinAmount <= 0 {
		c.MinAmount = min_amt
	}
	if c.MaxAmount <= 0 {
		c.MaxAmount = max_amt
	}
	if len(c.TranCode) <= 0 {
		c.TranCode = append(c.TranCode, "Other")
	}

	trs := m.TransactionList{}
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

	// time.Sleep(1 * time.Second) // used to test for concurrency and benchmark test
	return m.TransactionResult{ Transactions: trs, Err: nil}
}
