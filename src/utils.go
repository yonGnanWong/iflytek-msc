package src

import (
	"os"
	"strconv"
	"time"
)

var ch = "aaaaaaaaa`"

/**
获取字符串类型的当前时间戳
*/
func getTime() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

func getFileSize(path string) string {
	fInfo, err := os.Stat(path)
	assert1(err)
	return strconv.Itoa(int(fInfo.Size()))
}

func getNextSliceId() string {
	j := len(ch) - 1
	for j >= 0 {
		cj := ch[j : j+1]
		if cj != "z" {
			ch = ch[:j] + string([]rune(cj)[0]+1) + ch[j+1:]
			break
		} else {
			ch = ch[:j] + "a" + ch[j+1:]
			j = j - 1
		}
	}
	return ch
}
