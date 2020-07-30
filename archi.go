package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Dependency struct {
	FromLayer string
	ToLayer   string
}

type Architecture struct {
	Packages     map[string]string
	Dependencies []*Dependency
}

func (a *Architecture) GetLayer(pkg string) (string, bool) {
	for p, l := range a.Packages {
		if pkg == p {
			return l, true
		}
		if strings.HasPrefix(pkg, p+"/") {
			return l, true
		}
	}
	return "", false
}

func (a *Architecture) Contain(ip *Import) bool {
	_, ok := a.GetLayer(ip.From)
	if !ok {
		return false
	}
	_, ok = a.GetLayer(ip.To)
	return ok
}

func (a *Architecture) Valid(ip *Import) bool {
	fromLayer, ok := a.GetLayer(ip.From)
	if !ok {
		return false
	}
	toLayer, ok := a.GetLayer(ip.To)
	if !ok {
		return false
	}

	if fromLayer == toLayer {
		return true
	}

	for _, dep := range a.Dependencies {
		if fromLayer == dep.FromLayer && toLayer == dep.ToLayer {
			return true
		}
	}
	return false
}

func ReadArchitectureFromUML(umlPath string) (*Architecture, error) {
	f, err := os.Open(umlPath)
	if err != nil {
		return nil, err
	}

	pkgs := map[string]string{}
	deps := []*Dependency{}

	r := bufio.NewReader(f)
	for {
		b, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		line := string(b)
		if isLayer(line) {
			l, p, err := parseLayer(line)
			if err != nil {
				return nil, err
			}
			if _, ok := pkgs[p]; ok {
				return nil, fmt.Errorf("%s belongs to 2 layer. One package must belongs to only one package.", p)
			}
			pkgs[p] = l
		}

		if isDependency(line) {
			d, err := parseDependency(line)
			if err != nil {
				return nil, err
			}
			deps = append(deps, d)
		}
	}

	return &Architecture{
		Packages:     pkgs,
		Dependencies: deps,
	}, nil
}

func isLayer(line string) bool {
	return strings.Contains(line, ":")
}

func parseLayer(line string) (string, string, error) {
	ss := strings.Split(line, ":")
	if len(ss) != 2 {
		return "", "", fmt.Errorf("invalid layer description: %s", line)
	}

	return strings.TrimSpace(ss[0]), strings.TrimSpace(ss[1]), nil
}

func isDependency(line string) bool {
	return strings.Contains(line, "->")
}

func parseDependency(line string) (*Dependency, error) {
	ss := strings.Split(line, "->")
	if len(ss) != 2 {
		return nil, fmt.Errorf("invalid dependency description: %s", line)
	}

	return &Dependency{FromLayer: strings.TrimSpace(ss[0]), ToLayer: strings.TrimSpace(ss[1])}, nil
}
