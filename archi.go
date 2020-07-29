package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type LayerMap map[string]*Layer

type Layer struct {
	Name string
	Pkgs []string
}

type Dependencies struct {
	LayerMap     LayerMap
	Dependencies []*Dependency
}

type Dependency struct {
	From *Layer
	To   *Layer
}

func ReadArchiFromUML(umlPath string) (*Dependencies, error) {
	f, err := os.Open(umlPath)
	if err != nil {
		return nil, err
	}

	layerMap := NewLayerMap()
	dependencies := NewDependencies(layerMap)

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
			if err := layerMap.updateLayerMap(line); err != nil {
				return nil, err
			}
		}

		if isDependency(line) {
			if err := dependencies.updateDependencies(line); err != nil {
				return nil, err
			}
		}
	}

	return dependencies, nil
}

func NewLayerMap() LayerMap {
	return map[string]*Layer{}
}

func (lm LayerMap) updateLayerMap(line string) error {
	ss := strings.Split(line, ":")
	if len(ss) != 2 {
		return fmt.Errorf("invalid layer description: %s", line)
	}

	ss0 := strings.TrimSpace(ss[0])
	ss1 := strings.TrimSpace(ss[1])

	if layer, ok := lm[ss0]; !ok {
		lm[ss0] = &Layer{Name: ss0, Pkgs: []string{ss1}}
	} else {
		layer.Pkgs = append(layer.Pkgs, ss1)
	}

	return nil
}

func (lm LayerMap) Exist(pkg string) bool {
	for _, l := range lm {
		if l.Exist(pkg) {
			return true
		}
	}

	return false
}

func (lm LayerMap) GetLayer(pkg string) *Layer {
	for _, l := range lm {
		if l.Exist(pkg) {
			return l
		}
	}
	return nil
}

func (l Layer) Exist(pkg string) bool {
	for _, p := range l.Pkgs {
		if strings.HasPrefix(pkg, p) {
			return true
		}
	}
	return false
}

func isLayer(line string) bool {
	return strings.Contains(line, ":")
}

func isDependency(line string) bool {
	return strings.Contains(line, "->")
}

func NewDependencies(l LayerMap) *Dependencies {
	return &Dependencies{
		LayerMap:     l,
		Dependencies: []*Dependency{},
	}
}

func (d *Dependencies) updateDependencies(line string) error {
	ss := strings.Split(line, "->")
	if len(ss) != 2 {
		return fmt.Errorf("invalid dependency description: %s", line)
	}

	ss0 := strings.TrimSpace(ss[0])
	ss1 := strings.TrimSpace(ss[1])

	if _, ok := d.LayerMap[ss0]; !ok {
		return fmt.Errorf("unknown layer: %s", ss0)
	}
	if _, ok := d.LayerMap[ss1]; !ok {
		return fmt.Errorf("unknown layer: %s", ss1)
	}

	d.Dependencies = append(d.Dependencies, &Dependency{
		From: d.LayerMap[ss0],
		To:   d.LayerMap[ss1],
	})

	return nil
}