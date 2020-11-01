package main

import (
	"log"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
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
	u := page.FindByName("username")
	ue, err := u.Elements()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	ue[0].Value(os.Getenv("NEO_USER"))
	p := page.FindByClass("input-password")
	pe, err := p.Elements()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	pe[0].Value(os.Getenv("NEO_PASS"))

	// // click
	// b := page.FindByID("neo-login-btn")
	// b.Click()

	// sleep
	time.Sleep(time.Second * 10)
}
