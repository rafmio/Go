package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	tempDir, err := os.MkdirTemp(".", "varlog")
	if err != nil {
		fmt.Println(err)
		return
	}

	line := "Apr 11 23:00:15 localhost kernel: [23604659.769285] [UFW BLOCK] IN=eth0 OUT= MAC=52:54:00:7c:d8:0f:fe:54:00:7c:d8:0f:08:00 SRC=91.240.118.248 DST=194.58.102.129 LEN=40 TOS=0x00 PREC=0x00 TTL=250 ID=34638 PROTO=TCP SPT=41605 DPT=63383 WINDOW=1024 RES=0x00 SYN URGP=0 "
	fmt.Println("tempDir name:", tempDir)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	envVar := currentDir + "/" + tempDir
	fmt.Println("envVar:", envVar)

	fmt.Println("current dir: ", currentDir)
	defer os.RemoveAll(tempDir)

	file := filepath.Join(tempDir, "tmpfile")
	if err := os.WriteFile(file, []byte(line), 0666); err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second * 10)
}
