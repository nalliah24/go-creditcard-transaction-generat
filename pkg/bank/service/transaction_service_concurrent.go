package service

import (
	m "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
)

var GenerateAllTransactionsConc = func(cfgs []m.Config) (m.TransactionList, error) {
	rstream := make(chan m.TransactionResult)
	for _, cfg := range cfgs {
		go func(c m.Config) {
			rstream <- generateTransaction(c)
		}(cfg)
	}

	trans := m.TransactionList{}
	for i := 0; i < len(cfgs); i++ {
		v := <-rstream
		trans = append(trans, v.Transactions...)
	}
	
	// When more than one type of (CR|DR) found, each array starts with index 1
	// So need re-index the id to start 1 to total num of transactions
	trans.ReIndex()

	return trans, nil
}
