package base64ex

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"testing"
)

type Data struct {
	Id   uint
	Date string
}

func base64EncodeDecode() {

	data := Data{
		Id:   1,
		Date: "2022-09-09",
	}

	dEnc := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(uint64(data.Id), 10) + "@" + data.Date))
	fmt.Printf("%s\n", dEnc)

	dDec, _ := base64.StdEncoding.DecodeString(dEnc)
	fmt.Printf("%s\n", string(dDec))

}

func TestBase64EncodeDecode(t *testing.T) {
	base64EncodeDecode()
}
