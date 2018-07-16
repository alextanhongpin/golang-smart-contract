# Ethereum with go

- Deploying the contract
- Reading data from contract
- Interacting with contract
- Instantiating the deployed contract via address

https://zupzup.org/smart-contract-solidity/
https://zupzup.org/eth-smart-contracts-go/




# Generating the abi and bin

```bash
$ solc -o target --bin --abi Contract.sol
```

Output:

```
\target
    \WinnerTakesAll.abi
    \WinnerTakesAll.bin

```

## Compile

```bash
$ abigen --sol=Contract.sol --pkg=main --out=contract.go
```

Same as:

```bash
$ abigen --bin=./target/WinnerTakesAll.abi --abi=./target/WinnerTakesAll.bin --pkg=main --out=main.go
```