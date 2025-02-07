// Author: elvis (crypto man)
// Date: 2025
// Description: Code for transfer XRP
// Version: 1.0.0

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"github.com/ripple-go-sdk/crypto"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RippleTestSuite struct {
	suite.Suite
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *RippleTestSuite) SetupTest() {

}

func TestRippleTestSuite(t *testing.T) {
	suite.Run(t, new(RippleTestSuite))
}

func (suite *RippleTestSuite) TestBalanceWallet() {

	rippleClient := NewXRPLClient(true)

	balance, err := rippleClient.GetAccountBalance("rsHmyjLEp9GR3yi2LzThKkKWqJ1FwmyRxU")
	fmt.Println("balance, err", balance, err)

}


func (suite *RippleTestSuite) TestCreateWallet() {

	//
	key, err := crypto.NewECDSAKey(nil)
	if err != nil {
		return
	}
	address3, err := crypto.AccountId(key, nil)
	pri3 := hex.EncodeToString(key.Private(nil))
	pub3 := hex.EncodeToString(key.Public(nil))
	fmt.Println("address3", address3)
	fmt.Println("pri3", pri3)
	fmt.Println("pub3", pub3)

	// check address again:
	bt, _ := hex.DecodeString(pri3)
	key2 := crypto.LoadECDSKey(bt)
	fmt.Printf("Private2: %x\n", key2.Private(nil))

	fmt.Println("key2", key.Key.String())
	fmt.Println("key2", key.PrivateKey.Key.String())

	// var sequenceZero uint32
	account2, err := crypto.AccountId(key2, nil)
	if err != nil {
		err = errors.Wrap(err, "AccountId")
		return
	}

	fmt.Println("address2", account2.String())
	if account2.String() == address3.String() {
		fmt.Println("compare address okkkkkk")
	}

	assert.NotEmpty(suite.T(), "")

}

func (suite *RippleTestSuite) TestGetAddressFromPrivateKy() {

	priKeyStr := "..."
	// address: rsi4TyeNiSGbsTzXSmrAaz3QNWCkkFEvyk
	// rippleClient := NewXRPLClient(true)
	// var seq uint32
	bt, _ := hex.DecodeString(priKeyStr)
	key := crypto.LoadECDSKey(bt)
	fmt.Printf("pub: %x\n", key.Private(nil))

	fmt.Println("key", key.Key.String())
	fmt.Println("key", key.PrivateKey.Key.String())

	// var sequenceZero uint32
	account, err := crypto.AccountId(key, nil)
	if err != nil {
		err = errors.Wrap(err, "AccountId")
		return
	}

	fmt.Println("address1", account.String()) //=====> parse again ok
	assert.NotEmpty(suite.T(), "")

}

func (suite *RippleTestSuite) TestTransfer() {

	rippleClient := NewXRPLClient(true)

	fromAddress := "rJqqJzpivdmRwHQ5a3t3WbizJHjQrZYtG6"
	fromPrivateKey := "..."

	toAddress := "rHNvvvUCJpnJJbbvHaW5EBMP5dhQmieKEv"

	amount := big.NewInt(1 * 1e6)
	tx, err := rippleClient.Transfer(fromPrivateKey, fromAddress, toAddress, amount)

	fmt.Println("tx, err: ", tx, err)
}
