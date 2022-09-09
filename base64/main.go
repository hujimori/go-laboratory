package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

type Data struct {
	Id   uint
	Date string
}

func main() {

	data := Data{
		Id:   1,
		Date: "2022-09-09",
	}

	dEnc := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(data.Id), 10) + "@" + data.Date))
	fmt.Printf("%s\n", dEnc)

	dDec, _ := base64.StdEncoding.DecodeString(dEnc)
	fmt.Printf("%s\n", string(dDec))

}
