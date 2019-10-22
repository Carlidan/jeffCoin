# jeffCoin

```text
*** THE REPO IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

[![Go Report Card](https://goreportcard.com/badge/github.com/JeffDeCola/jeffCoin)](https://goreportcard.com/report/github.com/JeffDeCola/jeffCoin)
[![GoDoc](https://godoc.org/github.com/JeffDeCola/jeffCoin?status.svg)](https://godoc.org/github.com/JeffDeCola/jeffCoin)
[![Maintainability](https://api.codeclimate.com/v1/badges/0c7cf619a01dd65fc06b/maintainability)](https://codeclimate.com/github/JeffDeCola/jeffCoin/maintainability)
[![Issue Count](https://codeclimate.com/github/JeffDeCola/jeffCoin/badges/issue_count.svg)](https://codeclimate.com/github/JeffDeCola/jeffCoin/issues)
[![License](http://img.shields.io/:license-mit-blue.svg)](http://jeffdecola.mit-license.org)

_A cryptocurrency (transaction based data) built on decentralized
multi node P2P open network using a sha256 Proof of Work blockchain
with a REST JSON API and a TCP Server to communicate between
the nodes over IP._

Or more simply a distributed, decentralized, public ledger.

Table of Contents,

* [OVERVIEW](https://github.com/JeffDeCola/jeffCoin#overview)
* [1. BLOCKCHAIN](https://github.com/JeffDeCola/jeffCoin#1-blockchain)
* [2. MINER](https://github.com/JeffDeCola/jeffCoin#2-miner)
* [3. ROUTING NODE](https://github.com/JeffDeCola/jeffCoin#3-routing-node)
* [4. WALLET](https://github.com/JeffDeCola/jeffCoin#4-wallet)
* [5. WEBSERVER](https://github.com/JeffDeCola/jeffCoin#5-webserver)
* [RUN](https://github.com/JeffDeCola/jeffCoin#run)
  * [GENESIS NODE](https://github.com/JeffDeCola/jeffCoin#genesis-node)
  * [NEW NODE](https://github.com/JeffDeCola/jeffCoin#new-node)
  * [WEBSERVER AND API](https://github.com/JeffDeCola/jeffCoin#webserver-and-api)
  * [ROUTING NODE](https://github.com/JeffDeCola/jeffCoin#routing-node)

Documentation and reference,

* The Blockchain is built from my
  [single-node-blockchain-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/single-node-blockchain-with-REST)
* The Webserver (GUI & REST JSON API) is built from my
  [simple-webserver-with-REST](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-webserver-with-REST)
* The Routing Node (TCP Server) is built from my
  [simple-tcp-ip-server](https://github.com/JeffDeCola/my-go-examples/tree/master/api/simple-tcp-ip-server)
* Generating Keys and creating the jeffCoin address was built from
  [create-bitcoin-address-from-ecdsa-publickey](https://github.com/JeffDeCola/my-go-examples/tree/master/blockchain/create-bitcoin-address-from-ecdsa-publickey)
* ECDSA digital signature verification was built from
  [ecdsa-digital-signature](https://github.com/JeffDeCola/my-go-examples/tree/master/cryptography/asymmetric-cryptography/ecdsa-digital-signature)
* Refer to my
  [cheat sheet on blockchains](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/blockchain/blockchain-cheat-sheet)
* I got a lot of inspiration from
  [here](https://github.com/nosequeldeebee/blockchain-tutorial)

[GitHub Webpage](https://jeffdecola.github.io/my-go-examples/)

## OVERVIEW

This code is broken up into four main parts,

* [1. BLOCKCHAIN](https://github.com/JeffDeCola/jeffCoin/tree/master/blockchain)
  The blockchain code
* [2. MINER](https://github.com/JeffDeCola/jeffCoin/tree/master/miner)
  To mine the cryptocurrency
* [3. ROUTING NODE (TCP Server)](https://github.com/JeffDeCola/jeffCoin/tree/master/routingnode)
  To communicate between the P2P nodes (network)
* [4. WALLET](https://github.com/JeffDeCola/jeffCoin/tree/master/wallet)
  To hold the cryptocurrency

I also added a WebServer for a GUI and an API,

* [5. WEBSERVER](https://github.com/JeffDeCola/jeffCoin/tree/master/webserver)
  The API and GUI

jeffCoin will,

* Implement a blockchain using a sha256 hash
* Decentralized multi nodes with an open P2P Architecture
* Maintain a network of nodes
* A GUI and API

This illustration may help,

![IMAGE - jeffCoin-overview - IMAGE](docs/pics/jeffCoin-overview.jpg)

## 1. BLOCKCHAIN

The blockchain section is the core of the entire design. It will keep the
transactions secure. A transaction is a transfer of value between Bitcoin wallets.

A block in the blockchain is the following struct,

```go
type blockStruct struct {
    Index      int      `json:"index"`
    Timestamp  string   `json:"timestamp"`
    Data       []string `json:"data"`
    Hash       string   `json:"hash"`
    PrevHash   string   `json:"prevhash"`
    Difficulty int      `json:"difficulty"`
    Nonce      string   `json:"nonce"`
}
```

The states of a block are,

* **currentBlock** Receiving transactions and not part of blockchain
* **lockedBlock** To be mined and added to the blockchain.
* **Part of Chain** Already in the **blockchain**

**[BLOCKCHAIN-INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/blockchain-interface.go)**

* BLOCKCHAIN
  * **GenesisBlockchain()** Creates the blockchain
  * **LoadBlockchain()** Receives the blockchain and the currentBlock
    from a Network Node
    * **SENDBLOCKCHAIN** Request
  * **GetBlockchain()** Gets the blockchain
* BLOCK  
  * **GetBlock()** Gets a block (via Index number) from the blockchain
* LOCKED BLOCK
  * **GetLockedBlock** Gets the lockedBlock
* CURRENT BLOCK
  * **GetCurrentBlock** Gets the currentBlock
  * **AddTransactionToCurrentBlock()** Adds a transaction to the currentBlock
* COINS
  * **GetAddressBalance()** Gets jeffCoin Address balance
* TRANSACTIONS
  * **TransactionRequest()** Request to Transfer Coins to a jeffCoin Address

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/guts.go)**

* BLOCKCHAIN
  * **loadBlockchain()** Loads the entire blockchain
  * **getBlockchain()** Gets the blockchain
  * **replaceBlockchain()** Replaces chain with the longer one
* BLOCK
  * **getBlock()** Gets a block in the blockchain
  * **calculateBlockHash()** Calculates SHA256 hash on a block
  * **isBlockValid()** Checks if block is valid
* LOCKED BLOCK
  * **getLockedBlock** Gets the lockedBlock
  * **appendLockedBlock()** Appends the lockedBlock to the blockchain
* CURRENT BLOCK
  * **loadCurrentBlock()** Loads the currentBlock
  * **resetCurrentBlock()** Resets the currentBlock
  * **getCurrentBlock** Gets the currentBlock
  * **addTransactionToCurrentBlock()** Adds a transaction to the currentBlock
  * **lockCurrentBlock()** Moves the currentBlock to the lockedBlock
    and resets the currentBlock
* COINS
  * **getAddressBalance()** Gets jeffCoin Address balance

**[TRANSACTIONS](https://github.com/JeffDeCola/jeffCoin/blob/master/blockchain/transactions.go)**

* **transactionRequest()** Request to Transfer Coins to a jeffCoin Address

This may help understand transactions,

![IMAGE - transactions-and-ledger - IMAGE](docs/pics/transactions-and-ledger.jpg)

## 2. MINER

The miner. tbd.

## 3. ROUTING NODE

The Routing Node has two main parts, the nodeList
and the ability to handling Node Requests (TCP Server).

A node in the nodelist is the following struct,

```go
type nodeStruct struct {
    Index     int    `json:"index"`
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
    IP        string `json:"ip"`
    Port      string `json:"port"`
}
```

**[ROUTINGNODE-INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/routingnode-interface.go)**

* NODELIST
  * **GenesisNodeList()** Creates the nodeList
  * **LoadNodeList()** Receives the nodeList from a Network Node
    * **SENDNODELIST** Request
  * **GetNodeList()** Gets the nodeList
* NODE
  * **GetNode()** Gets a Node (via Index number) from the nodeList
  * **AppendNewNode()** Appends a New Node to the nodeList  
* THIS NODE
  * **LoadThisNode()** Loads thisNode
  * **GetThisNode()** Gets thisNode  
  * **AppendThisNode()** Appends thisNode to the nodeList  
  * **BroadcastThisNode()** Broadcasts thisNode to the Network
    * **ADDNEWNODE** Request

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/guts.go)**

* NODELIST
  * **loadNodeList()** Loads the entire nodeList
  * **getNodeList()** Gets the nodeList
* NODE
  * **getNode()** Gets a node in the nodeList
  * **appendNewNode()** Appends a node to the nodeList
* THIS NODE
  * **loadThisNode()** Loads thisNode
  * **getThisNode()** Gets thisNode
  * **appendThisNode()** Appends thisNode to the nodeList
  * **checkIfThisNodeinNodeList** - Check if thisNode is already in the nodeList

**[REQUESTS HANDLERS](https://github.com/JeffDeCola/jeffCoin/blob/master/routingnode/handlers.go)**

* BLOCKCHAIN
  * **SENDBLOCKCHAIN (SB)** Sends the blockchain & currentBlock to another node
* ROUTINGNODE
  * **ADDNEWNODE (NN)** Adds a node to the nodeList
  * **SENDNODELIST (GN)** Sends the nodeList to another node
* WALLET
  * **SENDADDRESSBALANCE (SAB)** Sends the coin balance for a jeffCoin Address
  * **TRANSACTIONREQUEST (TR)** Request to Transfer Coins to a jeffCoin Address

* **EOF**

## 4. WALLET

To make things simpler,

* All wallets will have addresses to keep the jeffCoin
* A transaction is a transfer of value (coins) between jeffCoin wallets
* Can only do a 1 to 1 transaction
* All transactions are broadcast to the network and if valid added to blockchain
* Receiver must wait until in blockchain to see transaction completed.

A wallet is the following struct,

```go
type walletStruct struct {
    PrivateKeyHex   string `json:"privateKeyHex"`
    PublicKeyHex    string `json:"publicKeyHex"`
    JeffCoinAddress string `json:"jeffCoinAddress"`
}
```

**[WALLET INTERFACE](https://github.com/JeffDeCola/jeffCoin/blob/master/wallet/wallet-interface.go)**

* WALLET
  * **GenesisWallet()** Creates the wallet
  * **GetWallet()** Gets the wallet
* COINS
  * **GetAddressBalance()** Gets the coin balance for a jeffCoin Address
    * **SENDADDRESSBALANCE** Request
  * **TransactionRequest()** Request to Transfer Coins to a jeffCoin Address
    * **TRANSACTIONREQUEST** Request

**[GUTS](https://github.com/JeffDeCola/jeffCoin/blob/master/wallet/guts.go)**

* WALLET
  * **getWallet()** Gets the wallet
  * **makeWallet()** Creates wallet with Keys and jeffCoin address
* KEYS
  * **generateECDSASKeys()** - Generate privateKeyHex and publicKeyHex
  * **encodeKeys()** - Encodes privateKeyRaw and publicKeyRaw
  * **decodeKeys()** - Decodes privateKeyRaw and publicKeyRaw
* JEFFCOIN ADDRESS
  * **generateBitcoinAddress()** - Creates jeffCoinAddress
  * **hashPublicKey()** - Hashes publicKeyHex
  * **checksumKeyHash()** - Checksums verPublicKeyHash
  * **encodeKeyHash()** - Encodes verPublicKeyHash & checkSum

## 5. WEBSERVER

**GUI**

* [192.168.20.100:1234](http://localhost:1234/)

**[API COMMANDS](https://github.com/JeffDeCola/jeffCoin/blob/master/webserver/handlers.go)**

* BLOCKCHAIN
  * /showBlockchain
  * /showBlock/{blockID}
  * /showlockedblock
  * /showcurrentblock
* NODELIST
  * /shownodelist
  * /shownode/{nodeID}
  * /showthisnode
* WALLET
  * /showwallet
  * /showjeffcoinaddress
  * /showaddressbalance
  * /transactionrequest/{address}/{value}

## RUN

If this is you first time running, you need to create the first node.
You only do this once.

### GENESIS NODE

```bash
go run jeffCoin.go \
       -genesis \
       -ip 192.168.20.100 \
       -webport 1234 \
       -tcpport 3334
```

### NEW NODES

Then all other nodes, you do something like this to hook
up to the network.  You need the ip of any working network node,

```bash
go run jeffCoin.go \
       -ip 192.168.20.100 \
       -webport 1235 \
       -tcpport 3335 \
       -netip 192.168.20.100 \
       -netport 3334
```

```bash
go run jeffCoin.go \
       -ip 192.168.20.100 \
       -webport 1236 \
       -tcpport 3336 \
       -netip 192.168.20.100 \
       -netport 3335
```

```bash
go run jeffCoin.go \
       -ip 192.168.20.100 \
       -webport 1237 \
       -tcpport 3337 \
       -netip 192.168.20.100 \
       -netport 3336
```

This will,

* Add your Node to the node list
  * That list will be updated with the network nodes
* Receive the blockchain

### WEBSERVER AND API

The user GUI,

[192.168.20.100:1234/](http://localhost:1234/)

You could also use curl from the command line,

```go
curl 192.168.20.100:1234
```

The main page will list the various API command.

For example, show a Particular Block,

[192.168.20.100:1234//showblock/0](http://192.168.20.100:1234/showblock/0)

### ROUTING NODE

Since their is no security setup yet, you can open a connection to the TCP server,

```txt
netcat -q -1 192.168.20.100 3334
```

And request commands such as,

```txt
ADDTRANSACTION
```

## UPDATE GITHUB WEBPAGE USING CONCOURSE (OPTIONAL)

For fun, I use concourse to update
[jeffCoin GitHub Webpage](https://jeffdecola.github.io/jeffCoin/)
and alert me of the changes via repo status and slack.

A pipeline file [pipeline.yml](https://github.com/JeffDeCola/jeffCoin/tree/master/ci/pipeline.yml)
shows the entire ci flow. Visually, it looks like,

![IMAGE - jeffCoin concourse ci pipeline - IMAGE](docs/pics/jeffCoin-pipeline.jpg)

The `jobs` and `tasks` are,

* `job-readme-github-pages` runs task
  [readme-github-pages.sh](https://github.com/JeffDeCola/jeffCoin/tree/master/ci/scripts/readme-github-pages.sh).

The concourse `resources types` are,

* `jeffCoin` uses a resource type
  [docker-image](https://hub.docker.com/r/concourse/git-resource/)
  to PULL a repo from github.
* `resource-slack-alert` uses a resource type
  [docker image](https://hub.docker.com/r/cfcommunity/slack-notification-resource)
  that will notify slack on your progress.
* `resource-repo-status` uses a resource type
  [docker image](https://hub.docker.com/r/dpb587/github-status-resource)
  that will update your git status for that particular commit.

For more information on using concourse for continuous integration,
refer to my cheat sheet on [concourse](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/operations-tools/continuous-integration-continuous-deployment/concourse-cheat-sheet).
