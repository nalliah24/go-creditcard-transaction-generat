package service

import (
	"testing"

	"github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
)

func TestGenerateTransactionWithZeroTransConfig(t *testing.T) {
	c := []model.Config{{NumOfTrans: 0}}
	_, err := GenerateAllTransactions(c)
	if err.Error() != "config error: number of transactions required" {
		t.Errorf(`expected num of trans error`)
	}
}

func TestGenerateTransactionWithEmptyTranTypeConfig(t *testing.T) {
	c := []model.Config{{NumOfTrans: 10, TranType: ""}}
	_, err := GenerateAllTransactions(c)
	if err.Error() != "config error: trancation code (CR|DR) required" {
		t.Errorf(`expected transaction code error`)
	}
}

func TestGenerateTransactionFor10CR(t *testing.T) {
	c := []model.Config{{NumOfTrans: 10, TranType: "CR"}}
	trans, _ := GenerateAllTransactions(c)
	if len(trans) != 10 {
		t.Errorf(`want %d transactions, but got %d`, c[0].NumOfTrans, len(trans))
	}
	// validate amount. when config not provided default to 100-5000
	// for _, tr := range trans {
	// 	if tr.Amount > 5000 {
	// 		t.Errorf(`transaction amount is > 5K default amount`)
	// 	}
	// }
}

func TestGenerateTransactionFor5CRWith2kTo20k(t *testing.T) {
	c := []model.Config{{NumOfTrans: 5, TranType: "CR", MinAmount: 2000.00, MaxAmount: 20000.00}}
	trans, _ := GenerateAllTransactions(c)
	if len(trans) != 5 {
		t.Errorf(`want %d transactions, but got %d`, c[0].NumOfTrans, len(trans))
	}
}

func TestGenerateTransactionFor5CRWithTranCodes(t *testing.T) {
	c := []model.Config{{NumOfTrans: 5, TranType: "CR", MinAmount: 2000.00, MaxAmount: 20000.00, TranCode: []string{"Hotel", "Flight", "Cab"}}}
	trans, _ := GenerateAllTransactions(c)
	if len(trans) != 5 {
		t.Errorf(`want %d transactions, but got %d`, c[0].NumOfTrans, len(trans))
	}

	for _, tr := range trans {
		if tr.TranCode == "Other" {
			t.Errorf(`transaction code 'Other' found. not expected`)
		}
	}
}
