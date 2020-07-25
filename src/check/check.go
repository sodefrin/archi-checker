package check

import (
	"github.com/sodefrin/archi-checker/src/archi"
	"github.com/sodefrin/archi-checker/src/parser"
)

type Result struct {
}

func ArchiCheck(umlPath, modulePath string, pkgs ...string) ([]*Result, error) {
	ips, err := parser.ParsePkgs(modulePath, pkgs...)
	if err != nil {
		return nil, err
	}

	deps, err := archi.ReadArchiFromUML(umlPath)
	if err != nil {
		return nil, err
	}

	return run(deps, ips), nil
}

func run(dep *archi.Dependencies, ips []*parser.Import) []*Result {
	return nil
}
