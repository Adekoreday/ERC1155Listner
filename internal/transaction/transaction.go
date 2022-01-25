package transaction 

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"
    "time"
    
	"go.mongodb.org/mongo-driver/mongo"
    "unchain/configs"
	"unchain/internal/contracts"
    "unchain/internal/models"
    "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

type LogTransferSingle struct {
	Operator common.Address
	From   common.Address
    To     common.Address
	Id     *big.Int
	Tokens  *big.Int
}

type LogTransferBatch struct {
	Operator common.Address
	From   common.Address
    To     common.Address
	Ids     [] *big.Int
	Tokens  [] *big.Int
}

var transactionCollection *mongo.Collection = configs.GetCollection(configs.DB, "transaction")

func Listen () {

	client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/56ab5afaf4d9451da8a2a72225d02aba")
    if err != nil {
        log.Fatal(err)
    }

	contractAddress := common.HexToAddress(configs.SmartContractAdd())
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

	logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
    if err != nil {
        log.Fatal(err)
    }

	logTransferSig := []byte("TransferSingle(address,address,address,uint256,uint256)")
    logTransferBatchSig := []byte("TransferBatch(address,address,address,uint256[],uint256[])")
    logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
    logTransferBatchSigHash := crypto.Keccak256Hash(logTransferBatchSig)

	fmt.Println(logTransferSigHash.Hex())
	fmt.Println(logTransferBatchSigHash.Hex())
	for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
        fmt.Println(vLog) // pointer to event log
		fmt.Println(vLog.Topics[0].Hex())
		switch vLog.Topics[0].Hex() {
        case logTransferSigHash.Hex():
            fmt.Printf("Log Name: Transfer\n")

            var transferSingleEvent LogTransferSingle

             err := contractAbi.UnpackIntoInterface(&transferSingleEvent, "TransferSingle", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }
            
            transferSingleEvent.Operator = common.HexToAddress(vLog.Topics[1].Hex())
            transferSingleEvent.From = common.HexToAddress(vLog.Topics[2].Hex())
            transferSingleEvent.To = common.HexToAddress(vLog.Topics[3].Hex())

            newTransaction := models.Transaction {
                TokenId:  transferSingleEvent.Id.String(),
                Operator: transferSingleEvent.Operator.Hex(),
                Sender:   transferSingleEvent.From.Hex(),
                Receiver: transferSingleEvent.To.Hex(),
                Token:    transferSingleEvent.Tokens.String(), 
                CreatedAt: time.Now(),
            }

            ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

            defer cancel()
            result, err := transactionCollection.InsertOne(ctx, newTransaction);

            if err != nil {
                fmt.Printf("Err: %s\n", err)
            }

            fmt.Printf("Operator: %s\n", transferSingleEvent.Operator.Hex())
            fmt.Printf("From: %s\n", transferSingleEvent.From.Hex())
			fmt.Printf("To: %s\n", transferSingleEvent.To.Hex())
            fmt.Printf("Tokens: %s\n", transferSingleEvent.Tokens.String())
            fmt.Printf("Result: %s\n", result)
        case logTransferBatchSigHash.Hex():
            fmt.Printf("Log Name: Batch Transfer\n")

            var transferBatchEvent LogTransferBatch

             err := contractAbi.UnpackIntoInterface(&transferBatchEvent, "TransferBatch", vLog.Data)
            if err != nil {
                log.Fatal(err)
            }
            //transferBatchEvent
            transferBatchEvent.Operator = common.HexToAddress(vLog.Topics[1].Hex())
            transferBatchEvent.From = common.HexToAddress(vLog.Topics[2].Hex())
            transferBatchEvent.To = common.HexToAddress(vLog.Topics[3].Hex())
            fmt.Printf("From: %s\n", transferBatchEvent.From)
            fmt.Printf("To: %s\n", transferBatchEvent.To)
        }
      }
    }
}