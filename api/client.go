package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tidwall/gjson"
)

/**

 {
    "status": "1",
    "message": "OK-Missing/Invalid API Key, rate limit of 1/5sec applied",
    "result": "9251482"
}
*/

type BlockResult struct {
	Result  string `json:"result" gorm:"column:result"`
	Message string `json:"message" gorm:"column:message"`
	Status  string `json:"status" gorm:"column:status"`
}

/**
{
    "status": "1",
    "message": "OK",
    "result": [
        {
            "blockNumber": "18427226",
            "timeStamp": "1698236639",
            "hash": "0x2bc3f5bca76dc7948643b0aedb54e2a43efd63d82a1aef208a71f87920b577f5",
            "nonce": "0",
            "blockHash": "0xc49ee9f56da9a0478961168d14b562b1d40b2de82ce179ac0c5ec3bf441d52f0",
            "transactionIndex": "143",
            "from": "0x19436f072822da4d6c8d007d15aaf0cbc6205224",
            "to": "0x000000f20032b9e171844b00ea507e11960bd94a",
            "value": "0",
            "gas": "338767",
            "gasPrice": "16193707989",
            "isError": "0",
            "txreceipt_status": "1",
            "input": "0x7e734c5a000000000000000000000000000000000000000000000000000000000000006000000000000000000000000000000000000000000000000000000000000000a055b7c347d57950f7e0130a57f3e1d8129eb75c2bd907c8b85f722d515a802276000000000000000000000000000000000000000000000000000000000000000c4a494f2043617073756c6573000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000034a494f0000000000000000000000000000000000000000000000000000000000029f64e3",
            "contractAddress": "",
            "cumulativeGasUsed": "23984569",
            "gasUsed": "253605",
            "confirmations": "6",
            "methodId": "0x7e734c5a",
            "functionName": "createClone(string name,string symbol,bytes32 salt)"
        }
    ]
}
*/

type TxResult struct {
	Result  []*Tx  `json:"result" gorm:"column:result"`
	Message string `json:"message" gorm:"column:message"`
	Status  string `json:"status" gorm:"column:status"`
}

type Tx struct {
	BlockHash         string `json:"blockHash" gorm:"column:blockHash"`
	FunctionName      string `json:"functionName" gorm:"column:functionName"`
	ContractAddress   string `json:"contractAddress" gorm:"column:contractAddress"`
	MethodID          string `json:"methodId" gorm:"column:methodId"`
	TransactionIndex  string `json:"transactionIndex" gorm:"column:transactionIndex"`
	Confirmations     string `json:"confirmations" gorm:"column:confirmations"`
	Nonce             string `json:"nonce" gorm:"column:nonce"`
	TimeStamp         string `json:"timeStamp" gorm:"column:timeStamp"`
	Input             string `json:"input" gorm:"column:input"`
	GasUsed           string `json:"gasUsed" gorm:"column:gasUsed"`
	IsError           string `json:"isError" gorm:"column:isError"`
	TxreceiptStatus   string `json:"txreceipt_status" gorm:"column:txreceipt_status"`
	BlockNumber       string `json:"blockNumber" gorm:"column:blockNumber"`
	Gas               string `json:"gas" gorm:"column:gas"`
	CumulativeGasUsed string `json:"cumulativeGasUsed" gorm:"column:cumulativeGasUsed"`
	From              string `json:"from" gorm:"column:from"`
	To                string `json:"to" gorm:"column:to"`
	Value             string `json:"value" gorm:"column:value"`
	Hash              string `json:"hash" gorm:"column:hash"`
	GasPrice          string `json:"gasPrice" gorm:"column:gasPrice"`
}

func SendTxTask(url string, hash string, chainCode int64) (int64, error) {
	payload := `{
    "blockChain": %v,
    "txHash": "%v"
}`

	payload = fmt.Sprintf(payload, chainCode, hash)

	//{
	//    "code": 0,
	//    "data": null
	//}
	resp, err := send2(url, "POST", payload)
	if err != nil {
		return 1, err
	}

	if gjson.Parse(resp).Get("code").Int() == 0 {
		return 0, nil
	}
	return 1, fmt.Errorf("%v", resp)
}

func GetBlock(sed int64, key string) (*BlockResult, error) {
	url := "https://api.etherscan.io/api?"
	url = fmt.Sprintf("%vmodule=block&action=getblocknobytime&timestamp=%v&closest=before&apikey=%v", url, sed, key)
	resp, err := send(url, "GET", "")
	if err != nil {
		return nil, err
	}
	var b BlockResult

	_ = json.Unmarshal([]byte(resp), &b)
	return &b, nil
}

func GetTx(address string, start, end int64, key string) (*TxResult, error) {
	url := "https://api.etherscan.io/api?"
	url = fmt.Sprintf("%vmodule=account&action=txlist&address=%v&startblock=%v&endblock=%v&page=0&offset=10000&sort=asc&apikey=%v", url, address, start, end, key)
	resp, err := send(url, "GET", "")
	if err != nil {
		return nil, err
	}
	var b TxResult

	_ = json.Unmarshal([]byte(resp), &b)
	return &b, nil
}

func send(url string, method string, payload string) (string, error) {
	req, _ := http.NewRequest(method, url, strings.NewReader(payload))
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	if gjson.ParseBytes(body).Get("status").Exists() {
		if gjson.ParseBytes(body).Get("status").String() == "1" {
			return string(body), nil
		}
	}

	return "", fmt.Errorf("%v", string(body))
}

func send2(url string, method string, payload string) (string, error) {
	req, _ := http.NewRequest(method, url, strings.NewReader(payload))
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
