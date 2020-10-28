package omnijson

type OmniSendResult string

type OmniSendCommand struct {
	FromAddress string `json:"fromaddress"`
	ToAddress   string `json:"toaddress"`
	PropertyId  int    `json:"propertyid"`
	Amount      string `json:"amount"`
}

func (cmd OmniSendCommand) Method() string {
	return "omni_send"
}

func (cmd OmniSendCommand) ID() string {
	return "1"
}

func (cmd OmniSendCommand) Params() []interface{} {
	return []interface{}{cmd.FromAddress, cmd.ToAddress, cmd.PropertyId, cmd.Amount}
}
