package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ArithemticOperationResponse struct {
	SlackUserName string `json:"slackUsername"`
	Result        int64  `json:"result"`
	OperationType string `json:"operation_type"`
}

type RequestBody struct {
	Operation_type string `json:"operation_type"`
	X              int    `json:"x"`
	Y              int    `json:"y"`
}

func PerformArithemticOperation(w http.ResponseWriter, r *http.Request) {

	var requestBody RequestBody

	json.NewDecoder(r.Body).Decode(&requestBody)

	var operationType = requestBody.Operation_type
	x := int64(requestBody.X)
	y := int64(requestBody.Y)

	switch operationType {
	case "addition":
		{
			jsonObject := add(x, y)
			w.Write(jsonObject)
		}
	case "subtraction":
		{
			jsonObject := subtract(x, y)
			w.Write(jsonObject)
		}
	case "multiplication":
		{
			jsonObject := multiply(x, y)
			w.Write(jsonObject)
		}
	}
}

func constructJsonObject(result int64, OperationType string) []byte {
	response := ArithemticOperationResponse{
		SlackUserName: UserSlackUsername,
		Result:        result,
		OperationType: OperationType,
	}

	jsonObject, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	return jsonObject
}

func add(x int64, y int64) []byte {
	result := x + y
	return constructJsonObject(result, "addition")
}

func subtract(x int64, y int64) []byte {
	result := x - y
	return constructJsonObject(result, "subtraction")
}

func multiply(x int64, y int64) []byte {
	result := x * y
	return constructJsonObject(result, "multiplication")
}
