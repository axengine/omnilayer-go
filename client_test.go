package omnilayer

import (
	"fmt"
	"github.com/axengine/omnilayer-go/omnijson"
	"testing"
)

//

func createOmni() *Client {
	omni := New(&ConnConfig{
		Host:                 "192.168.10.10:18332",
		Endpoint:             "",
		User:                 "usdtuser_cs",
		Pass:                 "usdtuser_cs",
		Proxy:                "",
		ProxyUser:            "",
		ProxyPass:            "",
		Certificates:         nil,
		DisableAutoReconnect: false,
		DisableConnectOnNew:  false,
		EnableBCInfoHacks:    false,
	})
	return omni
}

func TestGetNewAddress(t *testing.T) {
	omni := createOmni()
	ret, err := omni.OmniGetNewAddress(omnijson.OmniGetNewAddressCommand{})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
	// 2N1URX4Yn6y4NtHq7tSyyusYEEGLS7PFQDN
}

func TestValidAddress(t *testing.T) {
	omni := createOmni()
	ret, err := omni.ValidAddress(omnijson.ValidAddressCommand("moneyqMan7uh8FqdCA2BV5yZ8qVrc9ikLP"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}

func TestGetBalance(t *testing.T) {
	omni := createOmni()
	ret, err := omni.OmniGetBalance(omnijson.OmniGetBalanceCommand{"2N9RjPfvmYx9ykhRvkz4DPHRPRFMEAyQqtY", 2})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}

func TestFetchBlock(t *testing.T) {
	omni := createOmni()
	ret, err := omni.GetBlockChainInfo()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)

	txs, err := omni.OmniListBlockTransactions(1865230)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(txs)

	for _, v := range txs {
		tx, err := omni.OmniGetTransaction(v)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tx)
	}
}

func TestSend(t *testing.T) {
	omni := createOmni()
	ret, err := omni.OmniSend(omnijson.OmniSendCommand{"2N9RjPfvmYx9ykhRvkz4DPHRPRFMEAyQqtY", "2N1URX4Yn6y4NtHq7tSyyusYEEGLS7PFQDN", 2, "0.01"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(ret)
}
