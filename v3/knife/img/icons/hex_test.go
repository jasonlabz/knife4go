package icons

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	file, err := os.Open("./file.js")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	all, _ := io.ReadAll(file)
	rs := hex.EncodeToString(all)
	//rs, err := hex.DecodeString(js.APP_2FAB4AC5_JS_HEX_CONTENT)
	//if nil != err {
	//	fmt.Println("err:", err)
	//	return
	//}
	fmt.Println(rs)
}

var hexContent = ``
