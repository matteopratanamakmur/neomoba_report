package main

import (
	"encoding/csv"
	"log"
	"os"
	"time"

	"github.com/sclevine/agouti"
)

func failOnError(err error) {
	if err != nil {
		log.Fatal("Error:", err)
	}
}

func main() {
	driver := agouti.ChromeDriver()
	defer driver.Stop()

	err := driver.Start()
	failOnError(err)

	page, err := driver.NewPage()
	failOnError(err)

	// open neomoba page
	page.Navigate("https://trade.sbineomobile.co.jp/login")

	// put user info
	u := page.FindByName("username")
	ue, err := u.Elements()
	failOnError(err)
	ue[0].Value(os.Getenv("NEO_USER"))
	p := page.FindByClass("input-password")
	pe, err := p.Elements()
	failOnError(err)
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
		failOnError(err)
		time.Sleep(time.Second * 1)
	}

	// get portfolio info
	// plm := page.AllByClass("panels")
	// fmt.Print(plm.All("table"))
	const n = 10
	const m = 5
	pfs := make([][]string, n)
	for i := 0; i < n; i++ {
		pfs[i] = make([]string, m)
	}

	// open csv
	file, err := os.OpenFile("export.csv", os.O_WRONLY|os.O_CREATE, 0600)
	failOnError(err)
	defer file.Close()

	// truncate csv
	err = file.Truncate(0)
	failOnError(err)

	// write csv
	writer := csv.NewWriter(file)
	for i := 0; i < n; i++ {
		writer.Write(pfs[i])
	}
	writer.Flush()

	// sleep
	time.Sleep(time.Second * 60)
}
