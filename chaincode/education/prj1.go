package main

import (
    "fmt"
)

type SimpleChaincode struct {
}

var BackGroundNo int = 0
var RecordNo    int = 0


type School struct {
    Name        string
    Location    string
    Address     string
    PriKey      string
    PubKey      string
    StudentAddress []string
}

type Student struct {
    Name        string
    Address     string
    BackgroundId []int
}

type BackgroundId struct {
    Id          int
    ExitTime    int64
    Status      string //0: graduation 1:dropout
}

type Record Struct {
    Id              int
    SchoolAddress   string
    StudentAddress  string
    SchoolSign      string
    ModifyTime      int64
    ModifyOperation string
}


func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    if function == "createSchool" {
        return t.createSchool(stub, args)
    }

    return shim.Success(nil)
}

func (t *SimpleChaincode) createSchool(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expectiong 2")
    }

    var school School
    var schoolBytes []byte
    var stuAddress []string
    var address, priKey, pubKey string

    address, priKey, pubKey = GetAddress()

    school = School{Name: args[0], Location: args[1], Address:address, PriKey: priKey, PubKey: pubKey, StudentAddress: stuAddress}
    err := wirteSchool(stub, school)
    if err != nil {
        shim.Error("Error write school")
    }

    schoolBytes, err = json.Marshal(&school)
    if err != nil {
        return shim.Error("Error retrieving schoolBytes")
    }
    return shim.Success(schoolBytes)
}
    
