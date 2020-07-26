package main

import (
	"errors"
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
		return cfg, errors.New(usage)
	}

	return cfg, nil
}

func run(cfg config) int {
	invalidImports, err := check.ArchiCheck(cfg.umlPath, cfg.modulePath, cfg.pkgs...)
	if err != nil {
		fmt.Printf("%s\n", err)
		fmt.Println(usage)
		return 1
	}

	for _, ip := range invalidImports {
		importpos := ip.FileSet.Position(ip.Import.Path.ValuePos)
		fmt.Printf("%s: cannot import %s from %s\n", importpos, ip.To, ip.From)
	}

	if len(invalidImports) > 0 {
		return 1
	}

	return 0
}

func main() {
	cfg, err := parseConfig()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	os.Exit(run(cfg))
}
