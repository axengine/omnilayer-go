package omnijson

type ValidAddressResult struct {
	Isvalid bool `json:"isvalid"`
}

type ValidAddressCommand string

func (cmd ValidAddressCommand) Method() string {
	return "validateaddress"
}

func (cmd ValidAddressCommand) ID() string {
	return "1"
}

func (cmd ValidAddressCommand) Params() []interface{} {
	return []interface{}{cmd}
}
