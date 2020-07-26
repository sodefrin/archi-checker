package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sodefrin/archi-checker/src/check"
	"golang.org/x/mod/modfile"
)

type config struct {
	umlPath string
	pkgname string
	pkgs    []string
}

const usage = "archi-checker -uml [path to uml] [target pacakges]"

func parseConfig() (config, error) {
	cfg := config{}

	flag.StringVar(&cfg.umlPath, "uml", "", "uml path")
	mod, err := ioutil.ReadFile("./go.mod")
	if err != nil {
		return cfg, err
	}

	flag.Parse()
	cfg.pkgs = flag.Args()
	cfg.pkgname = modfile.ModulePath(mod)

	if cfg.umlPath == "" || cfg.pkgname == "" || len(cfg.pkgs) == 0 {
		fmt.Println(cfg)
		return cfg, errors.New(usage)
	}

	return cfg, nil
}

func run(cfg config) int {
	invalidImports, err := check.ArchiCheck(cfg.umlPath, cfg.pkgname, cfg.pkgs...)
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
