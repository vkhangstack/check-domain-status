package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/vkhangstack/check-domain-status/mail"
)

func main() {
	godotenv.Load()

	data := make([]string, 0)

	// open file
	f, err := os.Open("domains.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Println("***********************************************")
		domain := scanner.Text()
		fmt.Printf("domain: %s\n", domain)
		// isDomain := utils.IsDomain(domain)
		// fmt.Println(isDomain)
		// if !isDomain {
		// 	fmt.Fprintf(os.Stdout, "not isDomain: %s\n", []any{isDomain}...)
		// }
		http.DefaultClient.Timeout = 2 * time.Second
		res, err := http.Get(domain)

		if err != nil || res.StatusCode != 200 {
			data = append(data, domain)
		}

	}
	fmt.Printf("List domain failed: %s\r\n", data)

	if len(data) != 0 {
		errEmail := mail.SendNotificationMail(strings.Join(data, "\n"))

		if errEmail != nil {
			fmt.Println(err.Error())
		}
	}
}
