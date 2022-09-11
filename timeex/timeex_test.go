package timeex

import (
	"fmt"
	"testing"
	"time"
)

var layout = "2006-01-25T00:00:00Z"

func time2str(t time.Time) string {
	// レシーバーtを、"YYYY-MM-DDTHH-MM-SSZZZZ"という形の文字列に変換する
	return t.Format("2006-01-02T15:04:05Z07:00")
}

func timeEx() {

	// // fmt.Println(now)

	// // str := timeToString(now)
	// // fmt.Println(str)

	// t := stringToTime("2006-01-25T00:00:00Z")
	// fmt.Println(t)

	// var timeTime = time.Date(2022, 4, 1, 9, 0, 0, 0, time.Local)
	// fmt.Println(time2str(timeTime))
	// // 2022-04-01T09:00:00+09:00

	// layout := "2006-01-02T15:04:05Z"
	layout := time.RFC3339
	str := "2022-09-10T13:44:36Z"
	t, err := time.Parse(layout, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)
}

func timeToString(t time.Time) string {
	str := t.Format(layout)
	return str
}

func stringToTime(str string) time.Time {
	loc, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(layout, str, loc)
	return t
}

func TestTimeEx(t *testing.T) {
	timeEx()
}
