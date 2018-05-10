package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Test_Init(t *testing.T) {
	simpleCC := new(SimpleChaincode)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	mockStub.MockTransactionStart(txId)
	response := simpleCC.Init(mockStub)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Init failed")
		t.FailNow()
	}
}

func Test_initMarble(t *testing.T) {
	simpleCC := new(SimpleChaincode)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"newMarble", "blue", "35", "bob"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.initMarble(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 200 {
		fmt.Println("Init failed")
		t.FailNow()
	}
}

func Test_initMarble_incorrectArgs(t *testing.T) {
	simpleCC := new(SimpleChaincode)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"newMarble", "blue", "35"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.initMarble(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 500 {
		fmt.Println("Init failed")
		t.FailNow()
	}
}
