package service

import (
	"testing"

	"github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
)

func TestGenerateTransactionWithZeroTransConfig(t *testing.T) {
	cfgs := []model.Config{{NumOfTrans: 0}}
	_, err := GenerateAllTransactions(cfgs)
	if err.Error() != "config error: number of transactions required" {
		t.Errorf(`expected num of trans error %s`, err.Error())
	}
}

func TestGenerateTransactionWithEmptyTranTypeConfig(t *testing.T) {
	cfgs := []model.Config{{NumOfTrans: 10, TranType: ""}}
	_, err := GenerateAllTransactions(cfgs)
	if err.Error() != "config error: trancation code (CR|DR) required" {
		t.Errorf(`expected transaction code error %s`, err.Error())
	}
}

func TestGenerateTransactionFor10CR(t *testing.T) {
	cfgs := []model.Config{{NumOfTrans: 10, TranType: "CR"}}
	trans, _ := GenerateAllTransactions(cfgs)
	if len(trans) != 10 {
		t.Errorf(`want %d transactions, but got %d`, cfgs[0].NumOfTrans, len(trans))
	}
	// validate amount. when config not provided default to 100-5000
	// for _, tr := range trans {
	// 	if tr.Amount > 5000 {
	// 		t.Errorf(`transaction amount is > 5K default amount`)
	// 	}
	// }
}

func TestGenerateTransactionFor5CRWith2kTo20k(t *testing.T) {
	cfgs := []model.Config{{NumOfTrans: 5, TranType: "CR", MinAmount: 2000.00, MaxAmount: 20000.00}}
	trans, _ := GenerateAllTransactions(cfgs)
	if len(trans) != 5 {
		t.Errorf(`want %d transactions, but got %d`, cfgs[0].NumOfTrans, len(trans))
	}
}

func TestGenerateTransactionFor5CRWithTranCodes(t *testing.T) {
	cfgs := []model.Config{{NumOfTrans: 5, TranType: "CR", MinAmount: 2000.00, MaxAmount: 20000.00, TranCode: []string{"Hotel", "Flight", "Cab"}}}
	trans, _ := GenerateAllTransactions(cfgs)
	if len(trans) != 5 {
		t.Errorf(`want %d transactions, but got %d`, cfgs[0].NumOfTrans, len(trans))
	}

	for _, tr := range trans {
		if tr.TranCode == "Other" {
			t.Errorf(`transaction code 'Other' found. not expected`)
		}
	}
}

func TestGenerateTransactionWithMultipleTypesAndValidateIdOrderedSequentially(t *testing.T) {
	cfgs := []model.Config{
		{NumOfTrans: 5, TranType: "CR", MinAmount: 2000.00, MaxAmount: 20000.00},
		{NumOfTrans: 3, TranType: "DR", MinAmount: 500.00, MaxAmount: 1000.00}}
	trans, _ := GenerateAllTransactions(cfgs)
	if len(trans) != 8 {
		t.Errorf(`want %d transactions, but got %d`, cfgs[0].NumOfTrans, len(trans))
	}
	for i, tr := range trans {
		if tr.Id != i+1 {
			t.Errorf(`want %d transactions, but got %d`, i+1, tr.Id)
		}
	}
}

func TestGenerateTransactionWithMultipleTypesAndValidateSummary(t *testing.T) {
	cfgs := []model.Config{
		{NumOfTrans: 5, TranType: "CR", MinAmount: 2000.00, MaxAmount: 20000.00},
		{NumOfTrans: 3, TranType: "DR", MinAmount: 500.00, MaxAmount: 1000.00}}
	trans, _ := GenerateAllTransactions(cfgs)
	summaries, _ := trans.PrepareSummary(cfgs)

	for _, cfg := range cfgs {
		var trAmt float64
		var sumAmt float64
		for _, tr := range trans {
			if tr.TranType == cfg.TranType {
				trAmt += tr.Amount
			}
		}

		for _, s := range summaries {
			if s.TranType == cfg.TranType {
				sumAmt += s.TotalAmount
			}
		}

		if trAmt != sumAmt {
			t.Errorf(`summary does not match for %s: got tran total: %f got summary total: %f`, cfg.TranType, trAmt, sumAmt)
		}
	}
}