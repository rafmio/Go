package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type proccessInfo struct {
	pid     uint64
	ppid    uint64
	tgid    uint64
	name    string
	cmdline string
}

func main() {
	procDirName := "/proc"

	processInfoList := make([]*proccessInfo, 0)

	procDir, err := os.Open(procDirName)
	if err != nil {
		fmt.Println("opening file:", err)
	}

	defer procDir.Close()

	files, err := procDir.Readdir(0)
	if err != nil {
		fmt.Println("error reading file contents")
	}

	procInfoChan := make(chan proccessInfo, 32)
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)

		go func() {
			defer wg.Done()
			pi := new(proccessInfo)
			if file.IsDir() { // check if the file is a directory
				pid, err := strconv.ParseUint(file.Name(), 10, 32)
				if err != nil {
					// fmt.Println("error converting directory name to int:", err)
					// continue
				} else {
					pi.pid = pid // assigning a pid value to a corresponding field
					cmdlineFile := filepath.Join("/proc", file.Name(), "cmdline")
					cmdlineBytes, err := os.ReadFile(cmdlineFile)
					if err != nil {
						fmt.Println("reading file:", err)
					} else {
						pi.cmdline = string(cmdlineBytes)
					}

					statusFile := filepath.Join("/proc", file.Name(), "status")
					// statusBytes, err := os.ReadFile(statusFile)
					status, err := os.Open(statusFile)
					if err != nil {
						fmt.Println("reading file:", err)
					} else {
						scanner := bufio.NewScanner(status)
						for scanner.Scan() {
							line := scanner.Text()
							parts := strings.Split(line, ":")
							if len(parts) == 2 {
								key := strings.TrimSpace(parts[0])
								value := strings.TrimSpace(parts[1])

								switch key {
								case "Name":
									pi.name = value
								case "PPid":
									pi.ppid, _ = strconv.ParseUint(value, 10, 64)
								case "Tgid":
									pi.tgid, _ = strconv.ParseUint(value, 10, 64)
								}
							}
						}
					}
				}
			}
			// else {
			// 	continue
			// }
			// processInfoList = append(processInfoList, pi)
			procInfoChan <- *pi
		}()
	}

	go func() {
		for pi := range procInfoChan {
			processInfoList = append(processInfoList, &pi)
		}
	}()
	wg.Wait()
	close(procInfoChan)
	fmt.Println("len of the processInfoList:", len(processInfoList))

	for _, pi := range processInfoList {
		if strings.Contains(pi.name, "yandex") {
			// fmt.Printf("name:\t%s, pid:\t%d, ppid:\t%d,cmdline:\t%s\n", pi.name, pi.pid, pi.ppid, pi.cmdline)
			// fmt.Printf("name:\t%s, pid:\t%d, ppid:\t%d\n", pi.name, pi.pid, pi.ppid)
			fmt.Printf("name:\t%s\tpid:\t%d\tppid:\t%d\ttgid:\t%d\n", pi.name, pi.pid, pi.ppid, pi.tgid)
		}
	}

	os.Exit(0)
}
