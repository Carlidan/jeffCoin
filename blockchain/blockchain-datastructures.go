// jeffCoin 1. BLOCKCHAIN blockchain-datastructures.go

package blockchain

// BLOCK & BLOCKCHAIN ****************************************************************************************************

// blockStruct is your block
type blockStruct struct {
	BlockID      int64               `json:"blockID"`
	Timestamp    string              `json:"timestamp"`
	Transactions []transactionStruct `json:"transactions"`
	Hash         string              `json:"hash"`
	PrevHash     string              `json:"prevhash"`
	Difficulty   int                 `json:"difficulty"`
	Nonce        string              `json:"nonce"`
}

// pendingBlock - Receiving transactions and not part of chain
var pendingBlock = blockStruct{}

// lockedBlock - Going to be added to the chain (No more transactions)
var lockedBlock = blockStruct{}

// blockchainSlice is my block type
type blockchainSlice []blockStruct

// blockchain is the blockchain
var blockchain = blockchainSlice{}

// TRANSACTIONS **********************************************************************************************************

// transactionStruct is your transaction - To be placed in a block
type transactionStruct struct {
	TxID    int64           `json:"txID"`
	Inputs  []inputsStruct  `json:"inputs"`
	Outputs []outputsStruct `json:"outputs"`
}

// transactions is a transaction
var transaction = transactionStruct{}

// inputsStruct is a transaction input
type inputsStruct struct {
	RefTxID   int64  `json:"refTxID"`
	InPubKey  string `json:"inPubKey"`
	Signature string `json:"signature"`
}

// outputsStruct is a transaction output
type outputsStruct struct {
	OutPubKey string `json:"outPubKey"`
	Value     int64  `json:"value"`
}

// UNSPENT OUTPUT ********************************************************************************************************

type unspentOutputStruct struct {
	TxID  int64 `json:"txID"`
	Value int64 `json:"value"`
}

// TRANSACTION REQUESTS **************************************************************************************************

// txRequestMessageSignedStruct - Signed
type txRequestMessageSignedStruct struct {
	TxRequestMessage txRequestMessageStruct `json:"txRequestMessage"`
	Signature        string                 `json:"signature"`
}

// txRequestMessageStruct - Unsigned
type txRequestMessageStruct struct {
	SourceAddress string              `json:"sourceAddress"`
	Destinations  []destinationStruct `json:"destinations"`
}

// destinationStruct for multiple destinations
type destinationStruct struct {
	DestinationAddress string `json:"destinationAddress"`
	Value              int64  `json:"value"`
}
