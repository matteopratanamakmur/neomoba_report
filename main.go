package main

import (
	"fmt"
	"os"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver()
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page.Navigate("https://trade.sbineomobile.co.jp/login")
	fmt.Println(page.Title())
}
