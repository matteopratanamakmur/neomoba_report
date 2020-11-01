package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	// user / pass
	const username = "username"
	const password = "password"

	driver := agouti.ChromeDriver()
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		log.Fatalf("%s\n", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	// open neomoba page
	page.Navigate("https://trade.sbineomobile.co.jp/login")

	// put user info
	u := page.FindByName(username)
	ue, err := u.Elements()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	ue[0].Value(username)
	p := page.FindByName(password)
	pe, err := p.Elements()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	pe[0].Value(password)

	// click
	time.Sleep(time.Second * 1)
}
