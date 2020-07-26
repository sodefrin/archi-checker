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
		if !isTarget(deps, ip) {
			continue
		}

		if !isValidDependency(deps, ip) {
			ret = append(ret, ip)
		}
	}

	return ret
}

func isTarget(deps *archi.Dependencies, ip *parser.Import) bool {
	if deps.LayerMap.Exist(ip.From) && deps.LayerMap.Exist(ip.To) {
		return deps.LayerMap.GetLayer(ip.From).Name != deps.LayerMap.GetLayer(ip.To).Name
	}

	return false
}

func isValidDependency(deps *archi.Dependencies, ip *parser.Import) bool {
	for _, dep := range deps.Dependencies {
		if dep.From.Exist(ip.From) && dep.To.Exist(ip.To) {
			return true
		}
	}

	return false
}
