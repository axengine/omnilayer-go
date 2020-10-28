package omnijson

type OmniGetNewAddressResult string

type OmniGetNewAddressCommand struct{}

func (OmniGetNewAddressCommand) Method() string {
	return "getnewaddress"
}

func (OmniGetNewAddressCommand) ID() string {
	return "1"
}

func (OmniGetNewAddressCommand) Params() []interface{} {
	return nil
}
