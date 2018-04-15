package main

import (
	"github.com/armon/go-socks5"
	"os"
	"log"
)

func main() {
	creds := socks5.StaticCredentials{
		"user":"pass",
	}
	cator := socks5.UserPassAuthenticator{Credentials: creds}
	// Create a SOCKS5 server
	conf := &socks5.Config{
		AuthMethods: []socks5.Authenticator{cator},
		Logger:      log.New(os.Stdout, "", log.LstdFlags),
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost:port dynamicly create by heroki
	port := ":" + os.Getenv("PORT")
	log.Println("used port:" + port)
	if err := server.ListenAndServe("tcp", "0.0.0.0" + port); err != nil {
		panic(err)
	}
}