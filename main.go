package main

import (
	"flag"
	"fmt"

	"github.com/NaturalSelectionLabs/RSS3-Node/node"
	"github.com/ipfs/go-log/v2"
)

func setLogger(level string) error {
	lvl, err := log.LevelFromString(level)
	if err != nil {
		return err
	}
	log.SetAllLoggers(lvl)
	return nil
}

func main() {
	err := setLogger("info")
	help := flag.Bool("h", false, "Display Help")
	cfg := ParseFlags()

	if *help {
		fmt.Println("Run distributed node for RSS3.")
		fmt.Println()
		fmt.Println("Usage: Run './rss3node")
		flag.PrintDefaults()
		return
	}

	if err != nil {
		panic(err)
	}

	rss3node, err := node.NewRSS3Node(cfg)
	if err != nil {
		panic(err)
	}
	rss3node.Run()
}
