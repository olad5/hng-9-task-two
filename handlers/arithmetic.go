package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ArithemticOperationResponse struct {
	SlackUserName string `json:"slackUsername"`
	Result        int64  `json:"result"`
	OperationType string `json:"operation_type"`
}

func PerformArithemticOperation(w http.ResponseWriter, r *http.Request) {
	var operationType = r.FormValue("operation_type")
	x, _ := strconv.ParseInt(r.FormValue("x"), 10, 64)
	y, _ := strconv.ParseInt(r.FormValue("y"), 10, 64)

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
