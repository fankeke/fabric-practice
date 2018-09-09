fabric-practice
===============
A detail play-by-play of some major Fabric operations without docker


### 创建工作目录 
```
mkdir -p /exapp/fabric_test/
```

### 生成证书
```
cd /exapp/fabric-test/ && mkdir fabricconfig && cd fabricconfig

cryptogen showtemplate > crypto-config.yaml  ##然后修改 为自己需要的 top结构

cryptogen generate --config=crypto-config.yaml --output ./crypto-config/

```

### 修改hosts（各组织的域名）指向本机
`vim /etc/hosts/`
```
127.0.0.1 orderer.example.com
127.0.0.1 peer0.org1.example.com
127.0.0.1 peer1.org1.example.com
127.0.0.1 peer0.org2.example.com
127.0.0.1 peer1.org2.example.com
```

### 创世区块
```
cd /exapp/fabric-test/ && mkdir order && cd order

cp your-path-to/github.com/hyperledger/fabric/sampleconfig/configtx.yaml ./configtx.yaml #需要修改

configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./geneis.block
```


### 创建账本创世区块交易
```
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel.tx -channelID mychannel

```

### 创建锚节点更新交易
```
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./Org1MSPanchors.tx -channelID mychanel -asOrg Org1MSP
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./Org2MSPanchors.tx -channelID mychanel -asOrg Org2MSP
```


### 启动order节点
```
cd /exapp/fabric-test/order

cp your-path-to/github.com/hyperledger/fabric/sampleconfig/orderer.yaml ./orderer.yaml #需要修改

orderer start 

```


### 启动peer节点（一个peer进程充当了多个节点(端口区分))

* make sure in `peer` dir or errors occur`

```
cd /exapp/fabric-test/ && mkdir peer && cd peer

cp your-path-to/github.com/hyperledger/fafbric/sampleconfig/core.yaml ./orderer.yaml #需要修改

export set FABRIC_CFG_PATH=/exapp/fabric_test/peer

peer node start 
```




### 创建channel 
```
export CORE_PEER_LOCALMSPID=Org1MSP
export CORE_PEER_ADDRESS=peer0.org1.example.com:7051
export CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt

peer channel create -o  orderer.example.com:7050 -c mychannel -f ../order/channel.tx --tls true --cafile /exapp/fabric_test/fabricconfig/crypto-config/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

```

### 创建锚节点交易
```
configtxgen -profile TestTwoOrgsChannel -outputAnchorPeersUpdate ./Org1MSPanchors.tx -channelID mychannl -asOrg Org1MSP

configtxgen -profile TestTwoOrgsChannel -outputAnchorPeersUpdate ./Org2MSPanchors.tx -channelID mychannl -asOrg Org2MSP
```



###  peer0.org1 加入channel
```
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.example.com:7051
export set CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

peer channel join -b ./mychannel.block 
```


###  更新org1锚节点
```
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.example.com:7051 
export set CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

peer channel update -o orderer.example.com:7050 -c mychannel -f ../order/Org1MSPanchors.tx
```


### 部署chaincode
```
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.example.com:7051 
export set CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

peer chaincode install -n mycc -v 1.0 -p github.com/hyperledger/fabric-samples/chaincode/chaincode_example02/go
```


### 实例化
```
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.example.com:7051 
export set CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n mycc -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "OR ('Org1MSP.member', 'Org2MSP.member')"
```


### 发起查询
```
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.example.com:7051 
export set CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

peer chaincode query -C mychannel -n mycc -c '{"Args":["query","a"]}'
```


### 发起invoke
```
export set CORE_PEER_LOCALMSPID=Org1MSP
export set CORE_PEER_ADDRESS=peer0.org1.example.com:7051 
export set CORE_PEER_MSPCONFIGPATH=/exapp/fabric_test/fabricconfig/crypto-config/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp

peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n mycc -c '{"Args":["invoke", "a","b","10"]}'
```
