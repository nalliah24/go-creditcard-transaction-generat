package main

import (
	"fmt"
	"time"

	infra "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/infracture"
	ts "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service"
)

const (
	isIndent = false
)

func main() {
	genTransSequentially()
	genTransConcurrently()
}

func genTransSequentially() {
	fmt.Println("Generating mock transactions, based on 'config.json' file (Sequential)")
	cfgs, err := infra.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	start := time.Now()
	fmt.Println("Calling transaction service to generate")
	trans, err := ts.GenerateAllTransactions(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// total trans for file output name
	numOfTrans := len(trans)

	fmt.Println("Calling file handler to write 'output_n.json' file")
	infra.WriteJson(fmt.Sprintf("output_s_%d.json", numOfTrans), trans, isIndent)

	// prepare summary
	summaries, err := trans.PrepareSummary(cfgs)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Calling file handler to write 'output_summary_n.json' file")
	infra.WriteJson(fmt.Sprintf("output_s_summary_%d.json", numOfTrans), summaries, isIndent)

	fmt.Println("Done")
}

func genTransConcurrently() {
	fmt.Println("Generating mock transactions, based on 'config.json' file (Concurrent)")
	cfgs, err := infra.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	start := time.Now()
	fmt.Println("Calling transaction service to generate")
	trans, err := ts.GenerateAllTransactionsConc(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

	// total trans for file output name
	numOfTrans := len(trans)

	fmt.Println("Calling file handler to write 'output_n.json' file")
	infra.WriteJson(fmt.Sprintf("output_c_%d.json", numOfTrans), trans, isIndent)

	// prepare summary
	summaries, err := trans.PrepareSummary(cfgs)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Calling file handler to write 'output_summary_n.json' file")
	infra.WriteJson(fmt.Sprintf("output_c_summary_%d.json", numOfTrans), summaries, isIndent)

	fmt.Println("Done")
}

