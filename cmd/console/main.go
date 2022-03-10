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
	run(args)
}

func run(args []string) {
	path := "../1data_out"
	confName := "data/config.json"
	outFileName := fmt.Sprintf("%s/%s", path, "output.json")

	v := valArgs(args)
	if v != true {
		os.Exit(1)
	}

	isIndent, err := isIndent(args)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Read configuratins from ", confName)
	cfgs, err := readConfig(confName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// all good. start gen transactions
	var trans m.TransactionList
	start := time.Now()

	if args[0] == "1" {
		fmt.Println("Generating mock transactions sequentially")
		trans, err = genTransactionsSeq(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	if args[0] == "2" {
		fmt.Println("Generating mock transactions concurrently")
		trans, err = genTransactionsConc(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	elapsed := time.Since(start)
	fmt.Println("Time took: ", elapsed)

	// prepare summary
	summaries, err := trans.PrepareSummary(cfgs)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// prepare result
	result := m.Result{
		Timestamp:     time.Now().Format(time.RFC3339),
		NumberOfTrans: len(trans),
		Summaries:     summaries,
		Data:          trans,
	}

	fmt.Println("Delete existing output and write result: ", outFileName)
	err = infra.DeleteFileNameStartsWith(path, "output")
	if err != nil {
		fmt.Println("error deleting output file ", err.Error())
		return
	}

	infra.WriteJson(outFileName, result, isIndent)

	fmt.Println("Done")
}

func valArgs(args []string) bool {
	if len(args) <= 0 {
		fmt.Println("1st Argument required to process trans data: 1 = Sequential, 2 = Concurrent")
		fmt.Println("2nd Argument optional: true = 'Indent' Json in output file")
		return false
	}

	mode := args[0]
	if mode != "1" && mode != "2" {
		fmt.Println("1St argument must be 1 or 2")
		return false
	}
	return true
}

func isIndent(args []string) (bool, error) {
	if len(args) == 2 {
		v, err := strconv.ParseBool(args[1])
		if err != nil {
			return false, err
		}
		return v, nil
	}
	return false, nil
}

func readConfig(file string) ([]m.Config, error) {
	cfgs, err := infra.ReadConfig(file)
	if err != nil {
		return nil, err
	}
	return cfgs, nil
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
