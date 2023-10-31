package api

import (
	"testing"
	"time"
)

func TestGetBlock(t *testing.T) {

	r, err := GetBlock(time.Now().Unix(), "8VZRWD329TBQGGATEHUA57Z5E7GHQJ7P9Y")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestGetTx(t *testing.T) {

	r, err := GetTx("0x19436f072822DA4d6c8D007D15Aaf0cbc6205224", 18427226, 18427226, "8VZRWD329TBQGGATEHUA57Z5E7GHQJ7P9Y")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}

func TestSendTxTask(t *testing.T) {

	r, err := SendTxTask("http://47.245.118.1:9001/api/task/tx", "0xd353de80e0f757d18e7d3ffb5e251ca8f533ed9336bd7a76bde0561f1c8b8a56", 200)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(r)
	}
}
