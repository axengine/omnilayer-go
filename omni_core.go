package omnilayer

import (
	"encoding/json"
	"github.com/axengine/omnilayer-go/omnijson"
)

type futureOmniCreatePayloadSimpleSend chan *response

func (f futureOmniCreatePayloadSimpleSend) Receive() (omnijson.OmniCreatePayloadSimpleSendResult, error) {
	var result omnijson.OmniCreatePayloadSimpleSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxChange chan *response

func (f futureOmniCreateRawTxChange) Receive() (omnijson.OmniCreateRawTxChangeResult, error) {
	var result omnijson.OmniCreateRawTxChangeResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxOpReturn chan *response

func (f futureOmniCreateRawTxOpReturn) Receive() (omnijson.OmniCreateRawTxOpReturnResult, error) {
	var result omnijson.OmniCreateRawTxOpReturnResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniCreateRawTxReference chan *response

func (f futureOmniCreateRawTxReference) Receive() (omnijson.OmniCreateRawTxReferenceResult, error) {
	var result omnijson.OmniCreateRawTxReferenceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureGetInfo chan *response

func (f futureGetInfo) Receive() (omnijson.OmniGetInfoResult, error) {
	var result omnijson.OmniGetInfoResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetTransaction chan *response

func (f futureOmniGetTransaction) Receive() (omnijson.OmniGettransactionResult, error) {
	var result omnijson.OmniGettransactionResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniListBlockTransactions chan *response

func (f futureOmniListBlockTransactions) Receive() (omnijson.OmniListBlockTransactionsResult, error) {
	data, err := receive(f)
	if err != nil {
		return nil, err
	}

	result := make(omnijson.OmniListBlockTransactionsResult, 0)

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetBalance chan *response

func (f futureOmniGetBalance) Receive() (omnijson.OmniGetBalanceResult, error) {
	var result omnijson.OmniGetBalanceResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniGetNewAddress chan *response

func (f futureOmniGetNewAddress) Receive() (omnijson.OmniGetNewAddressResult, error) {
	var result omnijson.OmniGetNewAddressResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniSend chan *response

func (f futureOmniSend) Receive() (omnijson.OmniSendResult, error) {
	var result omnijson.OmniSendResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniSendWithFee chan *response

func (f futureOmniSendWithFee) Receive() (omnijson.OmniSendWithFeeResult, error) {
	var result omnijson.OmniSendWithFeeResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureValidAddress chan *response

func (f futureValidAddress) Receive() (omnijson.ValidAddressResult, error) {
	var result omnijson.ValidAddressResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}

type futureOmniDumpPrivkey chan *response

func (f futureOmniDumpPrivkey) Receive() (omnijson.OmniDumpPrivkeyResult, error) {
	var result omnijson.OmniDumpPrivkeyResult

	data, err := receive(f)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(data, &result)
	return result, err
}
