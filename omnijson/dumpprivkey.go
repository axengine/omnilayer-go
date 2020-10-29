package omnijson

type OmniDumpPrivkeyResult string

type OmniDumpPrivkeyCommand string

func (cmd OmniDumpPrivkeyCommand) Method() string {
	return "dumpprivkey"
}

func (cmd OmniDumpPrivkeyCommand) ID() string {
	return "1"
}

func (cmd OmniDumpPrivkeyCommand) Params() []interface{} {
	return []interface{}{cmd}
}
