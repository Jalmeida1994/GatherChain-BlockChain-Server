/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleTracker example simple Chaincode implementation
type SimpleTracker struct {
}

type group struct {
	Commit string `json:"commit"`
	Author int    `json:"author"`
}

// NOT WORKING
/*
type history struct {
	Commit    string
	Author    int
	Timestamp string
	TxID      string
}
*/

func (t *SimpleTracker) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")

	_, args := stub.GetFunctionAndParameters()
	var groupID, commitID string // Hash of the commit and group ID
	var studentID int
	var err error

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	// Initialize the chaincode
	groupID = args[0]
	commitID = args[1]
	studentID, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	fmt.Printf("studentID = %d, commitID = %d\n", studentID, commitID)

	// ==== Check if group already exists ====
	groupAsBytes, err := stub.GetState(groupID)
	if err != nil {
		return shim.Error("Failed to get group: " + err.Error())
	} else if groupAsBytes != nil {
		fmt.Println("This group already exists: " + groupID)
		return shim.Error("This group already exists: " + groupID)
	}

	group := &group{commitID, studentID}
	groupJSONasBytes, err := json.Marshal(group)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save group to state ===
	err = stub.PutState(groupID, groupJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)

}

func (t *SimpleTracker) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Push changes
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// Queries ledger
		return t.query(stub, args)
	} else if function == "history" {
		// Queries ledger
		return t.getHistory(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Pushes changes
func (t *SimpleTracker) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	groupID := args[0]
	commitID := args[1]
	studentID, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a numeric string")
	}

	groupAsBytes, err := stub.GetState(groupID)
	if err != nil {
		return shim.Error("Failed to get group:" + err.Error())
	} else if groupAsBytes == nil {
		return shim.Error("Group does not exist")
	}

	group := &group{commitID, studentID}
	groupJSONasBytes, err := json.Marshal(group)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save group to state ===
	err = stub.PutState(groupID, groupJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("- end commit (success)")
	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleTracker) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var groupID string // group ID

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	groupID = args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(groupID)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleTracker) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var groupID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the group to query")
	}

	groupID = args[0]
	valAsbytes, err := stub.GetState(groupID) //get the group from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + groupID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Group does not exist: " + groupID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

func (t *SimpleTracker) getHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	groupID := args[0]

	fmt.Printf("- start getHistory: %s\n", groupID)

	resultsIterator, err := stub.GetHistoryForKey(groupID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the group
	var buffer bytes.Buffer
	buffer.WriteString("[")

	// NOT WORKING
	// var h1 history

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON group)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true

		// NOT WORKING
		/*
			var g group
			erro := json.Unmarshal(response.Value, &g)
			if erro != nil {
				return shim.Error(erro.Error())
			}
			h1 = history{g.Commit, g.Author, time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String(), response.TxId}
		*/
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistory returning:\n%s\n", buffer.String())

	// NOT WORKING
	/*
		historyJSONbytes, e := json.Marshal(h1)
		if e != nil {
			return shim.Error(e.Error())
		}
		ef := ioutil.WriteFile("testdata.json", historyJSONbytes, 0644)
		if ef != nil {
			return shim.Error(ef.Error())
		}
	*/

	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(SimpleTracker))
	if err != nil {
		fmt.Printf("Error starting Simple tracker: %s", err)
	}
}
