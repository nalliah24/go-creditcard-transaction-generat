package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	infra "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/infracture"
	m "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
	ts "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service"
)

func main() {
	args := os.Args[1:]
	genTransactions(args)
}

func genTransactions(args []string) {
	isIndent := false
	outFileTrans := "output_trans"
	outFileSummary := "output_summary"

	if len(args) <= 0 {
		fmt.Println("1st Argument required to process trans data: 1 = Sequential, 2 = Concurrent")
		fmt.Println("2nd Argument optional: true = 'Indent' Json in output file")
		os.Exit(1)
	}

	mode := args[0]
	if mode != "1" && mode != "2" {
		fmt.Println("1St argument must be 1 or 2")
		os.Exit(1)
	}

	// check if json output as Indented?
	if len(args) == 2 {
		v, err := strconv.ParseBool(args[1])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		isIndent = v
	}

	// read config
	fmt.Println("Read configuratins from 'config.json'")
	cfgs, err := infra.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// gen transactions
	var trans m.TransactionList
	start := time.Now()

	if args[0] == "1" {
		fmt.Println("Generating mock transactions sequentially")
		trans, err = genTransactionsSeq(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		outFileTrans = fmt.Sprintf("%s_seq_%d.json", outFileTrans, len(trans))
	}

	if args[0] == "2" {
		fmt.Println("Generating mock transactions concurrently")
		trans, err = genTransactionsConc(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		outFileTrans = fmt.Sprintf("%s_conc_%d.json", outFileTrans, len(trans))
	}

	elapsed := time.Since(start)
	fmt.Println("Time took: ", elapsed)

	fmt.Println("Calling file handler to write transactions: ", outFileTrans)
	infra.WriteJson(outFileTrans, trans, isIndent)

	// prepare summary
	summaries, err := trans.PrepareSummary(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	outFileSummary = fmt.Sprintf("%s_%d.json", outFileSummary, len(trans))

	fmt.Println("Calling file handler to write summary: ", outFileSummary)
	infra.WriteJson(outFileSummary, summaries, isIndent)

	fmt.Println("Done")
}

func genTransactionsSeq(cfgs []m.Config) (m.TransactionList, error) {
	trans, err := ts.GenerateAllTransactions(cfgs)
	if err != nil {
		return nil, err
	}
	return trans, nil
}

func genTransactionsConc(cfgs []m.Config) (m.TransactionList, error) {
	trans, err := ts.GenerateAllTransactionsConc(cfgs)
	if err != nil {
		return nil, err
	}
	return trans, nil
}
