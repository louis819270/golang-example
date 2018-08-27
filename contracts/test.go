package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	Contract "golang-example/golang-contract-example/contract"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var conn *ethclient.Client
var contract *Contract.Contract
var err error

func main() {
	contractAddress := `0xe7655f6d95b2bdfff4a5f6369e994ecacd34f7be`

	conn, err = ethclient.Dial(`ws://localhost:8546`)
	if err != nil {
		log.Fatalf(`Failed to connect to the Ethereum client: %v`, err)
		return
	}

	contract, err = Contract.NewContract(common.HexToAddress(contractAddress), conn)
	if err != nil {
		log.Fatalf(`Failed to instantiate a Token contract: %v`, err)
		return
	}

	var ch = make(chan *Contract.ContractTest)
	sub, err := contract.WatchTest(nil, ch)

	if err != nil {
		log.Println(`Subscribe:`, err)
		return
	}

	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case log := <-ch:
				fmt.Println(`Sender:`, `0x`+hex.EncodeToString(log.Sender[:]))
				fmt.Println(`Times:`, log.Times)
			}
		}
	}()

	fmt.Println("Event watching...")
	fmt.Println("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
