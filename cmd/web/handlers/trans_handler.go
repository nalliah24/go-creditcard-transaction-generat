package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	infra "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/infracture"
	m "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/model"
	ts "github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service"
)

func TransHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "GET not supported. use POST /api/transactions"}`))

	case "POST":
		var cfgs []m.Config
		err := json.NewDecoder(r.Body).Decode(&cfgs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// all good. start gen transactions
		var trans m.TransactionList
		isIndent := false
		outFileTrans := "output_trans_web"
		outFileSummary := "output_summary_web"
		trans, err = ts.GenerateAllTransactionsConc(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		infra.WriteJson(fmt.Sprintf("%s_conc_%d.json", outFileTrans, len(trans)), trans, isIndent)

		// prepare summary
		summaries, err := trans.PrepareSummary(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		outFileSummary = fmt.Sprintf("%s_%d.json", outFileSummary, len(trans))
		infra.WriteJson(outFileSummary, summaries, isIndent)
		// endof trans and summary

		// extract first 5 trans as sample response
		sdata := trans[:5]
		msg := fmt.Sprintf("Transactions created successfully. Num Of Trans: %d", len(trans))

		w.WriteHeader(http.StatusCreated)
		m := fmt.Sprintf(`{"message": %s, "sample": %v}`, msg, sdata)
		w.Write([]byte(m))
	}
}
