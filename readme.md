
### ERC1155 LISTENER
This codebase consist of the following functionalities

1. Listening to the Ethereum test or mainnet blockchain transfer event logs. The logs should be are into a unified format and persisted to a local database.

2. dashboard rest API that returns a list of JSON objects having the following fields
- sender address
- receiver address
- smart-contract address
- balance of sender
- balance of receiver
- amount of token sent

### To make changes
Solidity version 
To setup this project kindly downgrade your solidity compiler version by 
```
brew uninstall solidity
pip3 install solc-select
solc-select install 0.4.24
solc-select use 0.4.24

```

Generate contract abi

```
solc --abi erc1155.sol | awk '/JSON ABI/{x=1;next}x' > erc1155.abi 
abigen --abi=erc1155.abi --pkg=token --out=erc1155.go

```
Setup mongodb using the compose file in this repo

```
    docker-compose up
```

Url of the Mongo Instance is 
```
mongodb://localhost:27018

```

Run the application

```
go run cmd/main.go 

```

Pending Features
- Rate limit and throtlling
- unit test of core logic
- ci pipelines config files