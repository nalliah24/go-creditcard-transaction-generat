package service

import (
	"testing"

	"github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
)

func BenchmarkGenerateAllTransactions(b *testing.B) {
	c := []model.Config{{NumOfTrans: 10000000, TranType: "CR"}}
	for n := 0; n < b.N; n++ {
		GenerateAllTransactions(c)
	}
}

func BenchmarkGenerateAllTransactionsConc(b *testing.B) {
	c := []model.Config{{NumOfTrans: 1000000, TranType: "CR"}}
	for n := 0; n < b.N; n++ {
		GenerateAllTransactionsConc(c)
	}
}