package main //读取文件内容

import (
	"fmt"
	"time"

	"github.com/hpcloud/tail"
)

func main() {
	//filename := "D:\\JAVA\\kafka_2.12-2.3.0\\config\\server.properties"
	filename := ".\\my.log"
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen: true,
		Follow: true,
		//Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}
	var msg *tail.Line
	var ok bool
	for true {
		msg, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		fmt.Println("msg:", msg)
	}
}
