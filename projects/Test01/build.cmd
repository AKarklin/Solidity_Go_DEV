docker run -v c:/Klins/Ethereum/Solidity_Go_DEV:/root ethereum/solc:stable --optimize --abi /root/projects/Test01/MySmartContract.sol -o /root/projects/Test01/build
docker run -v c:/Klins/Ethereum/Solidity_Go_DEV:/root ethereum/solc:stable --optimize --bin /root/projects/Test01/MySmartContract.sol -o /root/projects/Test01/build
..\..\abigen\abigen.exe --abi=build/MySmartContract.abi --bin=build/MySmartContract.bin --pkg=api --out=../../contracts_go_api/Test01/api/MySmartContract.go
