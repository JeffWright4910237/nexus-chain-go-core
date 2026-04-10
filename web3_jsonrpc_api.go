package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RPCResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func RPCHandler(w http.ResponseWriter, r *http.Request) {
	var req RPCRequest
	json.NewDecoder(r.Body).Decode(&req)
	resp := RPCResponse{Code: 200, Data: "NEXUS_HEIGHT_3000", Message: "success"}
	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/web3/rpc", RPCHandler)
	fmt.Println("Web3 JSON-RPC API Running on :9876")
	http.ListenAndServe(":9876", nil)
}
