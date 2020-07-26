package check

import (
	"github.com/sodefrin/archi-checker/src/archi"
	"github.com/sodefrin/archi-checker/src/parser"
)

func ArchiCheck(umlPath, modulePath string, pkgs ...string) ([]*parser.Import, error) {
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

func run(deps *archi.Dependencies, ips []*parser.Import) []*parser.Import {
	ret := []*parser.Import{}

	for _, ip := range ips {
		if !deps.LayerMap.Exist(ip.From) || !deps.LayerMap.Exist(ip.To) {
			continue
		}
		for _, dep := range deps.Dependencies {
			if xor(dep.From.Exist(ip.From), dep.To.Exist(ip.To)) {
				ret = append(ret, ip)
			}
		}
	}

	return ret
}

func xor(a, b bool) bool {
	return (a || b) && !(a && b)
}
