package main

import (
	"comparer/convert"
	"comparer/entity"
	"encoding/json"
	"fmt"
	"os"
)

const (
	LS             = `input\ЛицевыеСчета.txt`
	PU             = `input\ПостоянныеУдержания.txt`
	saveNormalized = false
)

func main() {
	accountsJson, err := convert.ToJSON(LS, saveNormalized)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	recoupmentJson, err := convert.ToJSON(PU, saveNormalized)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var accounts []*entity.Account
	var recoupments []*entity.Recoupment

	err = json.Unmarshal([]byte(*accountsJson), &accounts)
	err = json.Unmarshal([]byte(*recoupmentJson), &recoupments)

	resultJson, err := compare(accounts, recoupments)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resultFile, err := os.Create(`output\result.json`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resultJsonString := string(*resultJson)

	resultFile.WriteString(resultJsonString)

	fmt.Println(resultJsonString)
}

func compare(accounts []*entity.Account, recoupments []*entity.Recoupment) (*[]byte, error) {
	accMap := make(map[string]*entity.Account)
	for _, a := range accounts {
		accMap[a.EmplId] = a
	}

	var result []*entity.Recoupment
	for _, r := range recoupments {
		a := accMap[r.EmplId]

		diff := r.Alfa != a.Alfa
		diff = diff || r.Gazprom != a.Gazprom
		diff = diff || r.Databank != a.Databank
		diff = diff || r.Rosbank != a.Rosbank
		diff = diff || r.Rusbank != a.Rusbank
		diff = diff || r.Sber != a.Sber
		diff = diff || r.Hlyn != a.Hlyn

		if diff {
			result = append(result, r)
		}
	}

	resultJson, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &resultJson, nil
}
