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

	valArgs(args)

	isIndent := isIndent(args)

	fmt.Println("Read configuratins from 'config.json'")
	cfgs := readConfig()

	// all good. start gen transactions
	var trans m.TransactionList
	outFileTrans := "output_trans"
	outFileSummary := "output_summary"

	start := time.Now()

	if args[0] == "1" {
		fmt.Println("Generating mock transactions sequentially")
		trans = genTransactionsSeq(cfgs)
		infra.WriteJson(fmt.Sprintf("%s_seq_%d.json", outFileTrans, len(trans)), trans, isIndent)
	}

	if args[0] == "2" {
		fmt.Println("Generating mock transactions concurrently")
		trans = genTransactionsConc(cfgs)
		infra.WriteJson(fmt.Sprintf("%s_conc_%d.json", outFileTrans, len(trans)), trans, isIndent)
	}

	elapsed := time.Since(start)
	fmt.Println("Time took: ", elapsed)

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

func valArgs(args []string) {
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
}

func isIndent(args []string) bool {
	if len(args) == 2 {
		v, err := strconv.ParseBool(args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		return v
	}
	return false
}

func readConfig() []m.Config {
	cfgs, err := infra.ReadConfig("config.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfgs
}

func genTransactionsSeq(cfgs []m.Config) (m.TransactionList) {
	trans, err := ts.GenerateAllTransactions(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return trans
}

func genTransactionsConc(cfgs []m.Config) (m.TransactionList) {
	trans, err := ts.GenerateAllTransactionsConc(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return trans
}
