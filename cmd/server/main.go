package main

import (
	"context"
	"log"
	"seriouspoop/greedygame/pkg/config"
	"seriouspoop/greedygame/pkg/transport"
)

func main() {
	ctx := context.Background()
	appCfg, err := config.FromEnv()
	if err != nil {
		log.Fatal(err)
	}

	server, err := transport.NewServer(ctx, appCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer server.Shutdown(ctx)
	server.Initialize(ctx)
	server.Run(ctx)
}
