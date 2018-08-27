
make example.sol -> example.abi -> example.go

```
  cd contract
  solc --abi example.sol -o .
  abigen --abi example.abi --pkg contract --out example.go
```

set contract address

```
	contractAddress := `0xe7655f6d95b2bdfff4a5f6369e994ecacd34f7be`
```

set eth node (websocket)

```
	conn, err = ethclient.Dial(`ws://localhost:8546`)
```

run test

```
  go run test.go
```

call contract and get event

```
Event watching...
Press 'Enter' to exit...
Sender: 0xf46b6b9c7cb552829c1d3dfd8ffb11aabae782f6
Times: 15
```