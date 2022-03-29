package ews_test

import (
	"fmt"
	. "github.com/mhewedy/ews"
	"github.com/mhewedy/ews/ewsutil"
	"log"
	"testing"
)

func Test_Example(t *testing.T) {

	c := NewClient(
		"https://outlook.office365.com/EWS/Exchange.asmx",
		"username",
		"password",
		&Config{Dump: true, NTLM: true, SkipTLS: true},
	)

	err := testSendEmail(c, "Test Subject")

	if err != nil {
		log.Fatal("err>: ", err.Error())
	}

	_ = c

	fmt.Println("--- success ---")
}

func testSendEmail(c Client, subject string) error {
	return ewsutil.SendEmail(c,
		[]string{"ihsan.ugur@biltekin.onmicrosoft.com"},
		subject,
		"<HTML><Body><p><b>The email body, as plain text</b></p></body></HTML>", true, nil)
}
