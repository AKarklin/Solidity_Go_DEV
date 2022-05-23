docker run -v c:/Klins/Ethereum/Solidity_Go_DEV:/root --name solc_abi ethereum/solc:stable --optimize --abi /root/projects/Faucet/sol/Faucet.sol -o /root/projects/Faucet/build
docker rm solc_abi
docker run -v c:/Klins/Ethereum/Solidity_Go_DEV:/root --name solc_bin ethereum/solc:stable --optimize --bin /root/projects/Faucet/sol/Faucet.sol -o /root/projects/Faucet/build
docker rm solc_bin
..\..\abigen\abigen.exe --abi=build/Faucet.abi --bin=build/Faucet.bin --pkg=api --out=api/Faucet.go
