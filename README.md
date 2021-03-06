# Transaction Generator

Golang command line and web api utility to generate sample credit card transactions and output as json file.
Summarizes all transactions by type and timestamp with output.

Generated transactions can be used in other application as input.

### Methods:

Transaction Service has two methods. One generates sequentially and other concurrently. When generating hundreds of transactions concurrently using channels, there is a significant amount of performance improvement. see below for benchmark test.

```code
ts = reference the service in main / caller
ts.GenerateAllTransactions(cfgs []m.Config) (m.TransactionList, error)
ts.GenerateAllTransactionsConc(cfgs []m.Config) (m.TransactionList, error)

func (trans TransactionList) PrepareSummary(cfgs []Config) (SummaryList, error)
```

### Required:

#### Console App:
config.json - This file located in 'data/' directory, contains an array of input parameters for transaction types. Number of transactions and Transaction Type are required fields. Other fields are optional and if missing in the config file, it will be replaced by the application. Transaction date is currently default to current date.

Golang executable

#### Web API:
Run the server from web folder. Create a POST request to host:5000/api/transactions

Create a JSON POST request similar to below and attached to body of the post call. Can use Postman or other clients.

Todo: currently it creates and saves the file. can be extenced to download the file.

Sample config file
```json
[
  {
	"num_of_trans": 50,
	"tran_type": "CR",
	"min_amt": 50,
	"max_amt": 200,
	"tran_code": ["Flight", "Accommodation", "Cab"]
  },
  {
	 "num_of_trans": 20,
	 "tran_type": "DR",
	 "min_amt": 100,
	 "max_amt": 500,
	 "tran_code": ["Return", "Payment"]
  }
]
```

### Run:

```code
console
\creditcard\transaction_generator\cmd\console>console.exe
\creditcard\transaction_generator\cmd\console>go run main.go

web
creditcard\transaction_generator\cmd\web\server>go run main.go

tests
\creditcard\transaction_generator\pkg\bank\service>go test
\creditcard\transaction_generator\pkg\bank\service>go test -run=XXX -bench=.
```


### Notes:
Utility uses float64 for amount. It only need to output an amount, but no need for calculations. For accurate calculations use decimal packages such as shopsprint/decimal.

### Benchmark

```code

{NumOfTrans: 10000000, TranType: "CR"}

C:\Users\major\Documents\projects2\golang117\creditcard\transaction_generator\pkg\bank\service>go test -run=XXX -bench=.
goos: windows
goarch: amd64
pkg: github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkGenerateAllTransactions-8                     1        4658107900 ns/op
BenchmarkGenerateAllTransactionsConc-8                 3         447450100 ns/op
PASS
ok      github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service       7.630s

C:\Users\major\Documents\projects2\golang117\creditcard\transaction_generator\pkg\bank\service>go test -run=XXX -bench=.
goos: windows
goarch: amd64
pkg: github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkGenerateAllTransactions-8                     1        4589963200 ns/op
BenchmarkGenerateAllTransactionsConc-8                 3         449144633 ns/op
PASS
ok      github.com/nalliah24/go-creditcard-transaction-generator/pkg/bank/service       7.561s
```
