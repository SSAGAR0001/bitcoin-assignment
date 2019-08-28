package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"bitbucket.org/work/test/db"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	go func() {
		router := mux.NewRouter()
		router.HandleFunc("/{blockNumber}", getData)
		log.Fatal(http.ListenAndServe(":8081", router))
	}()
	client, err := ethclient.Dial("https://kovan.infura.io/v3/6c6f87a10e12438f8fbb7fc7c762b37c")
	if err != nil {
		log.Fatal("error connection to the block chain server: ", err)
	}
	fmt.Println("we have a connection")
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("error getting the latest block number: ", err)
	}
	fmt.Println(header.Number.String())

	latestBlock, err := strconv.ParseInt(header.Number.String(), 0, 64)
	if err != nil {
		log.Fatal("error converting: ", err)
	}

	for i := latestBlock; i > latestBlock-10000; i-- {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(i))
		if err != nil {
			log.Fatal("error getting the block: ", err)
		}
		// fmt.Println(block.Number().Uint64())

		if len(block.Transactions()) == 0 {
			continue
		}

		for _, tx := range block.Transactions() {
			msg, err := tx.AsMessage(types.NewEIP155Signer(tx.ChainId()))
			if err != nil {
				log.Fatal("error getting sender's address: ", err)
			}
			fromAddr := msg.From().Hex()
			toAddr := tx.To().Hex()
			blockNumber := fmt.Sprintf("%v", i)
			tHash := tx.Hash().Hex()

			if fromAddr == "" || toAddr == "" || blockNumber == "" || tHash == "" {
				continue
			}

			go func() {
				err := db.Mgr.Insert(fromAddr, toAddr, blockNumber, tHash)
				if err != nil {
					log.Fatal("error inserting into the db: ", err)
				}
			}()
		}
	}
}

func getData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]
	var resp []db.TransactionStruct
	resp, err := db.Mgr.Fetch(blockNumber)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	if len(resp) < 1 {
		fmt.Fprint(w, "no transaction for block number: "+blockNumber)
		return
	}
	jsonData, err := json.MarshalIndent(resp, "", "")
	fmt.Fprint(w, string(jsonData))
}
