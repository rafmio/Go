package main

import (
	"fmt"
	"net/url"
)

func main() {
	path := `https://link.alfabank.ru/statement/?query=&queryType=all&period%5BfromDate%5D=2024-01-01&period%5BtoDate%5D=2024-01-16&periodType=month&shouldMergeCommissionsByDay=false&shouldMergeCharityByDay=false&rangeSum%5B0%5D=&rangeSum%5B1%5D=`
	u, err := url.Parse(path)
	if err != nil {
		fmt.Println("url.Parse: parsing error")
	}
	fmt.Println("scheme:", u.Scheme)
	fmt.Println("opaque:", u.Opaque)
	fmt.Println("user  :", u.User)
	fmt.Println("host  :", u.Host)
	fmt.Println("path  :", u.Path)
	fmt.Println("raw path   :", u.RawPath)
	fmt.Println("omit host  :", u.OmitHost)
	fmt.Println("force query:", u.ForceQuery)
	fmt.Println("raw query  :", u.RawQuery)
	fmt.Println("fragment   :", u.Fragment)
	fmt.Println("raw fragment:", u.RawFragment)
}
