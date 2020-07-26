package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sodefrin/archi-checker/src/check"
)

type config struct {
	umlPath    string
	modulePath string
	pkgs       []string
}

const usage = "archi-checker -uml [path to uml] -module [repo url]"

func parseConfig() (config, error) {
	cfg := config{}

	flag.StringVar(&cfg.umlPath, "uml", "", "uml path")
	flag.StringVar(&cfg.modulePath, "module", "", "repo url")

	cfg.pkgs = flag.Args()

	if cfg.umlPath == "" || cfg.modulePath == "" || len(cfg.pkgs) == 0 {
		fmt.Println(usage)
		os.Exit(1)
	}
	return config{}, nil
}

func run(cfg config) {
	ret, err := check.ArchiCheck(cfg.umlPath, cfg.modulePath, cfg.pkgs...)
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		fmt.Println(usage)
		os.Exit(1)
	}

	for _, v := range ret {
		fmt.Println(v.From)
		fmt.Println(v.To)
	}
	os.Exit(0)
}

func main() {
	cfg, err := parseConfig()
	if err != nil {
		fmt.Printf("[ERROR] %s\n", err)
		fmt.Println(usage)
		os.Exit(1)
	}

	run(cfg)
}
