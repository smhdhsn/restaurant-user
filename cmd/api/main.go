package main

import (
	"fmt"
	"log"

	"github.com/smhdhsn/bookstore-user/internal/config"
)

func main() {
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)
}
