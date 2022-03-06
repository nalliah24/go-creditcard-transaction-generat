package model

type TransactionList []Transaction

func (trans TransactionList) ReIndex() {
	for i := 0; i < len(trans); i++ {
		trans[i].Id = i + 1
	}
}

func (trans TransactionList) PrepareSummary(cfgs []Config) (SummaryList, error) {
	types := map[string]string{}
	for _, cfg := range cfgs {
		types[cfg.TranType] = cfg.TranType
	}

	var summaries = SummaryList{}
	for _, typ := range types {
		summaries = append(summaries, prepareSummaryForType(typ, trans))
	}

	return summaries, nil
}

func prepareSummaryForType(typ string, trs TransactionList) Summary {
	var s = Summary{}
	s.TranType = typ
	for _, tr := range trs {
		if tr.TranType == typ {
			s.TotalAmount = s.TotalAmount + tr.Amount
			s.NumOfTrans++
		}
	}
	return s
}
