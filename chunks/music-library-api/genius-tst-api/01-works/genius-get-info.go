package main

import (
	"fmt"

	"github.com/gabyshev/genius-api/genius"
)

func main() {
	accessToken := "ycPo8Ic25mjQO6-Kcka0gUeeZDNNDmhCFh1o19YNPLGXY4W95_xXGuPwCDVjBeJA"
	client := genius.NewClient(nil, accessToken)

	response, err := client.GetArtistHTML(2029)
	if err != nil {
		fmt.Println("Error getting artist HTML:", err)
		return
	}

	fmt.Println(response.Response.Artist)

	fmt.Println("--------------")

}
