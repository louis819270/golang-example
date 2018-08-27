
make example.sol -> example.abi -> example.go

```
cd contract
solc --abi example.sol -o .
abigen --abi example.abi --pkg contract --out example.go
```

