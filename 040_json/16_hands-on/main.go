package main

import (
	"fmt"
	"encoding/json"
)

type data struct {
	Code        int    `json:"Code"`
	Description string `json:"Description"`
}

func main() {
	received := `[{"Code":200,"Description":"StatusOK"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":302,"Description":"StatusFound"},{"Code":303,"Description":"StatusSeeOther"},{"Code":307,"Description":"StatusTemporaryRedirect"},{"Code":400,"Description":"StatusBadRequest"},{"Code":401,"Description":"StatusUnauthorized"},{"Code":402,"Description":"StatusPaymentRequired"},{"Code":403,"Description":"StatusForbidden"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":301,"Description":"StatusMovedPermanently"},{"Code":404,"Description":"StatusNotFound"},{"Code":405,"Description":"StatusMethodNotAllowed"},{"Code":418,"Description":"StatusTeapot"},{"Code":500,"Description":"StatusInternalServerError"}]`

	var receivedData []data
	err := 	json.Unmarshal([]byte(received),&receivedData)
	if err != nil{
		panic(err)
	}
	for _,dataItem := range receivedData{
		fmt.Printf("%d : %s\n",dataItem.Code,dataItem.Description)
	}
}