# go-mail - Simple and easy way to send mails in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/wneessen/go-mail)](https://goreportcard.com/report/github.com/wneessen/go-mail) [![Build Status](https://api.cirrus-ci.com/github/wneessen/go-mail.svg)](https://cirrus-ci.com/github/wneessen/go-mail) <a href="https://ko-fi.com/D1D24V9IX"><img src="https://uploads-ssl.webflow.com/5c14e387dab576fe667689cf/5cbed8a4ae2b88347c06c923_BuyMeACoffee_blue.png" height="20" alt="buy ma a coffee"></a>

The main idea of this library was to provide a simple interface to sending mails for
my [JS-Mailer](https://github.com/wneessen/js-mailer) project. It quickly evolved into a 
full-fledged mail library.

**This library is "WIP" an should not be considered "production ready", yet.**

go-mail follows idiomatic Go style and best practice. It's only dependency is the Go Standard Library.
It combines a lot of functionality from the standard library to give easy and convenient access to
mail and SMTP related tasks.

## Features
Some of the features of this library:
* [X] Only Standard Library dependant
* [X] Modern, idiotmatic Go
* [X] Sane and secure defaults
* [X] SSL/TLS support
* [X] StartTLS support with different policies
* [X] Makes use of contexts for a better control flow and timeout/cancelation handling
* [X] SMTP Auth support (LOGIN, PLAIN, CRAM-MD5, DIGEST-MD5)
* [X] RFC5322 compliant mail address validation
* [X] Support for common mail header field generation (Message-ID, Date, Bulk-Precedence, etc.)
* [X] Reusing the same SMTP connection to send multiple mails
* [ ] Support for different encodings
* [ ] Support for attachments
* [ ] Go template support

## Example
```go
package main

import (
	"context"
	"fmt"
	"github.com/wneessen/go-mail"
	"os"
	"time"
)

func main() {
	c, err := mail.NewClient("mail.example.com", mail.WithTimeout(time.Millisecond*500),
		mail.WithTLSPolicy(mail.TLSMandatory), mail.WithSMTPAuth(mail.SMTPAuthDigestMD5),
		mail.WithUsername("tester@example.com"), mail.WithPassword("secureP4ssW0rd!"))
	if err != nil {
		fmt.Printf("failed to create new client: %s\n", err)
		os.Exit(1)
	}
	defer c.Close()

	if err := c.DialAndSend(); err != nil {
		fmt.Printf("failed to dial: %s\n", err)
		os.Exit(1)
	}
}
```