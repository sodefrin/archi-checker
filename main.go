package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sodefrin/archi-checker/src/check"
	"golang.org/x/mod/modfile"
)

const Version = ""

const usage = `Usage: archi-checker [options...] [pkgs...]

Options: 
-version, -v
  Print version and exit.

-help, -h
  Show this usage.

-strict, -s
  Fail if not defined dependency is found.

-uml, -u
  Uml file path. By default, .archi-checker.yml.

-package, -p
  Package path of your project, By default use gomodule path.
`

const (
	exitOK    = 0
	exitError = 1
)

func run() int {
	var (
		version bool
		help    bool

		strict  bool
		umlPath string
		pkgName string
	)

	flag.BoolVar(&version, "version", false, "")
	flag.BoolVar(&version, "v", false, "")

	flag.BoolVar(&help, "help", false, "")
	flag.BoolVar(&help, "h", false, "")

	flag.BoolVar(&strict, "strict", false, "")
	flag.BoolVar(&strict, "s", false, "")

	flag.StringVar(&umlPath, "uml", ".archi-checker.uml", "")
	flag.StringVar(&umlPath, "u", ".archi-checker.uml", "")

	flag.StringVar(&pkgName, "package", "", "")
	flag.StringVar(&pkgName, "p", "", "")

	flag.Parse()

	if version {
		if Version == "" {
			fmt.Println("Unknown version")
		} else {
			fmt.Println(Version)
		}
		return exitError
	}

	if help {
		fmt.Println(usage)
		return exitError
	}

	if pkgName == "" {
		mod, err := ioutil.ReadFile("./go.mod")
		if err != nil {
			fmt.Printf("Please specify package path. By default, read package path from go.mod; %s\n", err)
			return exitError
		}
		pkgName = modfile.ModulePath(mod)
	}

	pkgs := flag.Args()

	if len(pkgs) == 0 {
		fmt.Println("Please specify packages.")
		return exitError
	}

	invalidImports, err := check.ArchiCheck(umlPath, pkgName, pkgs...)
	if err != nil {
		fmt.Printf("[ERROR] %s\n%s\n", err, usage)
		return exitError
	}

	if len(invalidImports) > 0 {
		for _, ip := range invalidImports {
			importpos := ip.FileSet.Position(ip.Import.Path.ValuePos)
			fmt.Printf("%s: cannot import %s from %s\n", importpos, ip.To, ip.From)
		}
		return exitError
	}

	return exitOK
}

func main() {
	os.Exit(run())
}
