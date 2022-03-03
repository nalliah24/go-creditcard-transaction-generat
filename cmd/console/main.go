package main

import (
	"fmt"

	infra "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/infracture"
	ts "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service"
)


func main() {
	fmt.Println("Generating mock transactions, based on 'config.json' file")
	cfgs, err := infra.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Calling transaction service to generate")
	trans, err := ts.GenerateAllTransactions(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// calc number of total trans for file output name
	var numOfRows int
	for _, cfg := range cfgs {
		numOfRows += cfg.NumOfTrans
	}

	fmt.Println("Calling file handler to write 'output_n.json' file")
	infra.WriteJson(fmt.Sprintf("output_%d.json", numOfRows), trans)

	// prepare summary
	summaries, err := ts.PrepareSummary(trans)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Calling file handler to write 'output_summary_n.json' file")
	infra.WriteJson(fmt.Sprintf("output_summary_%d.json", numOfRows), summaries)

	fmt.Println("Done")
}
