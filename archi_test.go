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
