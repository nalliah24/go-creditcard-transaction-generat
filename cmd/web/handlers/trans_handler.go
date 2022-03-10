package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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
		path := "data"
		outFileName := fmt.Sprintf("%s/%s", path, "output_web.json")
		var trans m.TransactionList
		isIndent := false

		trans, err = ts.GenerateAllTransactionsConc(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// prepare summary
		summaries, err := trans.PrepareSummary(cfgs)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := m.Result{
			Timestamp:     time.Now().Format(time.RFC3339),
			NumberOfTrans: len(trans),
			Summaries:     summaries,
			Data:          trans,
		}

		fmt.Println("Delete existing output and write result: ", outFileName)
		err = infra.DeleteFileNameStartsWith(path, "output_")
		if err != nil {
			fmt.Println("error deleting output file ", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		infra.WriteJson(outFileName, result, isIndent)

		// extract first 5 trans as sample response
		sdata := trans[:5]
		msg := fmt.Sprintf("Transactions created successfully. Num Of Trans: %d", len(trans))

		w.WriteHeader(http.StatusCreated)
		m := fmt.Sprintf(`{"message": %s, "sample": %v}`, msg, sdata)
		w.Write([]byte(m))
	}
}
