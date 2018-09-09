=================

### install chaincode
```
peer chaincode install -n education -v 1.0 -p github.com/hyperledger/fabric-samples/chaincode/app_dev/education
```

#### createSchool 
```
* args[0] name of school 
* args[1] location of shcool

peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n education -c '{"Args":["createSchool","shcool_a","loc_a"]}'  

* response: payload:"{\"Name\":\"shcool_a\",\"Location\":\"loc_a\",\"Address\":\"3fb803b7ba7e7c8a6c1ca418ce95a20d\",\"PriKey\":\"3fb803b7ba7e7c8a6c1ca418ce95a20d1\",\"PubKey\":\"3fb803b7ba7e7c8a6c1ca418ce95a20d2\",\"StudentAddress\":null}"  

```

### creatStudent 
```
* args[0] name of student

peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n education -c '{"Args":["createStudent","student_a"]}'  

* response: payload:"{\"Name\":\"student_a\",\"Address\":\"a057abb10475889ab2cb09b35ba7e459\",\"BackgroundId\":null}" 
```

### enrollStudent
```
* args[0] address of school
* args[1] sign of school
* args[2] address of student

peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n education -c '{"Args":["enrollStudent","3fb803b7ba7e7c8a6c1ca418ce95a20d","shcool_a_sign","a057abb10475889ab2cb09b35ba7e459"]}' 
response: payload:"{\"Id\":1,\"SchoolAddress\":\"3fb803b7ba7e7c8a6c1ca418ce95a20d\",\"StudentAddress\":\"a057abb10475889ab2cb09b35ba7e459\",\"SchoolSign\":\"shcool_a_sign\",\"ModifyTime\":1535892811,\"ModifyOperation\":\"2\"}" 

```

### updateDiploma
```
* # args[0] address of school
* # args[1] sign of school
* # args[2] address of student
* # args[3] modify operation 0: graduation 1: dropout
peer chaincode invoke -o orderer.example.com:7050 -C mychannel -n education -c '{"Args":["updateDiploma","3fb803b7ba7e7c8a6c1ca418ce95a20d","shcool_a_sign","a057abb10475889ab2cb09b35ba7e459","0"]}'

* response: payload:"{\"Id\":3,\"SchoolAddress\":\"3fb803b7ba7e7c8a6c1ca418ce95a20d\",\"StudentAddress\":\"a057abb10475889ab2cb09b35ba7e459\",\"SchoolSign\":\"shcool_a_sign\",\"ModifyTime\":1535893059,\"ModifyOperation\":\"0\"}" 
```

