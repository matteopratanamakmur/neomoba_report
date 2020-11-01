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

	// click
	b := page.FindByID("neo-login-btn")
	b.Click()
	time.Sleep(time.Second * 1)

	// open portfolio page
	page.Navigate("https://trade.sbineomobile.co.jp/account/portfolio")
	time.Sleep(time.Second * 5)

	// page size
	page.Size(300, 1080)
	time.Sleep(time.Second * 5)

	// scroll
	for i := 0; i < 10; i++ {
		err = page.RunScript("scroll(0, document.body.scrollHeight);", nil, nil)
		if err != nil {
			log.Fatalf("%s\n", err)
		}
		time.Sleep(time.Second * 1)
	}

	// sleep
	time.Sleep(time.Second * 30)
}
