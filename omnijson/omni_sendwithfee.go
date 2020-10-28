package omnijson

type OmniSendWithFeeResult string

type OmniSendWithFeeCommand struct {
	FromAddress string `json:"fromaddress"`
	ToAddress   string `json:"toaddress"`
	PropertyId  int    `json:"propertyid"`
	Amount      string `json:"amount"`
	FeeAddr     string `json:"feeaddress"`
}

func (cmd OmniSendWithFeeCommand) Method() string {
	return "omni_funded_send"
}

func (cmd OmniSendWithFeeCommand) ID() string {
	return "1"
}

func (cmd OmniSendWithFeeCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyId, cmd.Amount, cmd.FeeAddr}
}
