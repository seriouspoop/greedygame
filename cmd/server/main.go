package main

import (
	"fmt"
	"log"
	"seriouspoop/greedygame/pkg/config"
)

func main() {
	appCfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appCfg)
}
