package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestReadArcguFromUML(t *testing.T) {
	ret, err := ReadArchiFromUML("testdata/sample.uml")
	if err != nil {
		t.Fatal(err)
	}

	want := &Dependencies{
		LayerMap: LayerMap{
			"domains": &Layer{
				Name: "domains", Pkgs: []string{"github.com/sodefrin/test/src/domains"},
			},
			"infrastructures": &Layer{
				Name: "infrastructures", Pkgs: []string{"github.com/sodefrin/test/src/infrastructures/proto", "google.golang.org/grpc"},
			},
			"interfaces": &Layer{
				Name: "interfaces", Pkgs: []string{"github.com/sodefrin/test/src/interfaces"},
			},
			"usecases": &Layer{
				Name: "usecases", Pkgs: []string{"github.com/sodefrin/test/src/usecases"},
			},
			"utils": &Layer{
				Name: "utils", Pkgs: []string{"log"},
			},
		},
		Dependencies: []*Dependency{
			{
				From: &Layer{
					Name: "usecases", Pkgs: []string{"github.com/sodefrin/test/src/usecases"},
				},
				To: &Layer{
					Name: "domains", Pkgs: []string{"github.com/sodefrin/test/src/domains"},
				},
			},
			{
				From: &Layer{
					Name: "interfaces", Pkgs: []string{"github.com/sodefrin/test/src/interfaces"},
				},
				To: &Layer{
					Name: "usecases", Pkgs: []string{"github.com/sodefrin/test/src/usecases"},
				},
			},
			{
				From: &Layer{
					Name: "interfaces", Pkgs: []string{"github.com/sodefrin/test/src/interfaces"},
				},
				To: &Layer{
					Name: "domains", Pkgs: []string{"github.com/sodefrin/test/src/domains"},
				},
			},
			{
				From: &Layer{
					Name: "interfaces", Pkgs: []string{"github.com/sodefrin/test/src/interfaces"},
				},
				To: &Layer{
					Name: "infrastructures", Pkgs: []string{"github.com/sodefrin/test/src/infrastructures/proto", "google.golang.org/grpc"},
				},
			},
			{
				From: &Layer{
					Name: "usecases", Pkgs: []string{"github.com/sodefrin/test/src/usecases"},
				},
				To: &Layer{
					Name: "utils", Pkgs: []string{"log"},
				},
			},
			{
				From: &Layer{
					Name: "infrastructures", Pkgs: []string{"github.com/sodefrin/test/src/infrastructures/proto", "google.golang.org/grpc"},
				},
				To: &Layer{
					Name: "utils", Pkgs: []string{"log"},
				},
			},
			{
				From: &Layer{
					Name: "interfaces", Pkgs: []string{"github.com/sodefrin/test/src/interfaces"},
				},
				To: &Layer{
					Name: "utils", Pkgs: []string{"log"},
				},
			},
		},
	}

	if diff := cmp.Diff(want, ret); diff != "" {
		t.Fatal(diff)
	}
}
