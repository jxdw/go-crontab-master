package common

import (
	"fmt"
	"time"
	"net/http"
)

func FmtErr(err error)  {
	fmt.Println(time.Now().Format("2006-01-02 15-04-05 "),err)
}

func ResponseErr(resp http.ResponseWriter, errno int,errInfo  string,data interface{})  {
	if bytes, err := BuildResponse(errno,errInfo,data); err == nil{
		resp.Write(bytes)
	}
	return
}