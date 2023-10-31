package util

import (
	"errors"
	"os"
)

const (
	LatestBlockPath = "./data"
	LatestBlockFile = "./data/tx.json"
)

func DeleteFile(path string) error {
	return os.Remove(path)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func WriteLatestBlock(content string) error {
	if ok, _ := PathExists(LatestBlockFile); !ok {
		//不存在
		err := os.MkdirAll(LatestBlockPath, os.ModePerm)
		if err != nil {
			//log.Println(err.Error())
			return err
		}
		err = os.Chmod(LatestBlockPath, 0777)
		if err != nil {
			return err
		}
	}

	return os.WriteFile(LatestBlockFile, []byte(content), 0777)
}

func ReadLatestBlock() ([]byte, error) {
	if ok, _ := PathExists(LatestBlockFile); ok {
		bs, err := os.ReadFile(LatestBlockFile)
		if err != nil {
			return nil, err
		}
		return bs, nil
	} else {
		return nil, errors.New("not found blockchain.json")
	}
}
