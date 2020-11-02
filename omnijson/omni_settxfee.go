package omnijson

type SetTxFeeResult bool

type SetTxFeeCommand struct {
	TxFee float64
}

func (cmd SetTxFeeCommand) Method() string {
	return "settxfee"
}

func (cmd SetTxFeeCommand) ID() string {
	return "1"
}

func (cmd SetTxFeeCommand) Params() []interface{} {
	return []interface{}{cmd.TxFee}
}
