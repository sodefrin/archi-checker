package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadArchitectureFromUML(t *testing.T) {
	ret, err := ReadArchitectureFromUML("testdata/sample.uml")
	if err != nil {
		t.Fatal(err)
	}

	want := &Architecture{
		Packages: map[string]string{
			"github.com/sodefrin/test/src/domains":               "domains",
			"github.com/sodefrin/test/src/infrastructures/proto": "infrastructures",
			"google.golang.org/grpc":                             "infrastructures",
			"github.com/sodefrin/test/src/interfaces":            "interfaces",
			"github.com/sodefrin/test/src/usecases":              "usecases",
			"log":                                                "utils",
		},
		Dependencies: []*Dependency{
			{
				FromLayer: "usecases",
				ToLayer:   "domains",
			},
			{
				FromLayer: "interfaces",
				ToLayer:   "usecases",
			},
			{
				FromLayer: "interfaces",
				ToLayer:   "domains",
			},
			{
				FromLayer: "interfaces",
				ToLayer:   "infrastructures",
			},
			{
				FromLayer: "usecases",
				ToLayer:   "utils",
			},
			{
				FromLayer: "infrastructures",
				ToLayer:   "utils",
			},
			{
				FromLayer: "interfaces",
				ToLayer:   "utils",
			},
		},
	}

	if diff := cmp.Diff(want, ret); diff != "" {
		t.Fatal(diff)
	}
}

func TestArchitectureGetLayer(t *testing.T) {
	testCases := map[string]struct {
		pkgs      map[string]string
		pkg       string
		wantLayer string
		wantOK    bool
	}{
		"equal": {
			pkgs: map[string]string{
				"a": "la",
			},
			pkg:       "a",
			wantLayer: "la",
			wantOK:    true,
		},
		"not found": {
			pkgs:      map[string]string{},
			pkg:       "a",
			wantLayer: "",
			wantOK:    false,
		},
		"not contains": {
			pkgs: map[string]string{
				"a": "la",
			},
			pkg:       "ab",
			wantLayer: "",
			wantOK:    false,
		},
		"contains": {
			pkgs: map[string]string{
				"a": "la",
			},
			pkg:       "a/b",
			wantLayer: "la",
			wantOK:    true,
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			a := &Architecture{
				Packages: v.pkgs,
			}

			l, ok := a.GetLayer(v.pkg)
			if l != v.wantLayer {
				t.Fatalf("want layer %v but have %v", v.wantLayer, l)
			}
			if ok != v.wantOK {
				t.Fatalf("want OK %v but have %v", v.wantOK, ok)
			}
		})
	}
}

func TestArchitectureContain(t *testing.T) {
	testCases := map[string]struct {
		pkgs   map[string]string
		ip     *Import
		wantOK bool
	}{
		"not found": {
			pkgs: map[string]string{},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: false,
		},
		"ok": {
			pkgs: map[string]string{
				"a": "la",
				"b": "lb",
			},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: true,
		},
		"a not found": {
			pkgs: map[string]string{
				"b": "lb",
			},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: false,
		},
		"b not found": {
			pkgs: map[string]string{
				"a": "la",
			},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: false,
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			a := &Architecture{
				Packages: v.pkgs,
			}

			ok := a.Contain(v.ip)
			if ok != v.wantOK {
				t.Fatalf("want OK %v but have %v", v.wantOK, ok)
			}
		})
	}
}

func TestArchitectureValid(t *testing.T) {
	testCases := map[string]struct {
		pkgs   map[string]string
		deps   []*Dependency
		ip     *Import
		wantOK bool
	}{
		"invalid": {
			pkgs: map[string]string{
				"a": "la",
				"b": "lb",
			},
			deps: []*Dependency{
				{
					FromLayer: "lb",
					ToLayer:   "la",
				},
			},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: false,
		},
		"ok": {
			pkgs: map[string]string{
				"a": "la",
				"b": "lb",
			},
			deps: []*Dependency{
				{
					FromLayer: "la",
					ToLayer:   "lb",
				},
			},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: true,
		},
		"a not found": {
			pkgs: map[string]string{
				"b": "lb",
			},
			deps: []*Dependency{},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: false,
		},
		"b not found": {
			pkgs: map[string]string{
				"a": "la",
			},
			deps: []*Dependency{},
			ip: &Import{
				From: "a",
				To:   "b",
			},
			wantOK: false,
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			a := &Architecture{
				Packages:     v.pkgs,
				Dependencies: v.deps,
			}

			ok := a.Valid(v.ip)
			if ok != v.wantOK {
				t.Fatalf("want OK %v but have %v", v.wantOK, ok)
			}
		})
	}
}

func TestReadArchitectureFromUMLDuplicatedDefinition(t *testing.T) {
	_, err := ReadArchitectureFromUML("testdata/duplicated.uml")
	if err == nil {
		t.Fatal("expected duplicated error")
	}
}
