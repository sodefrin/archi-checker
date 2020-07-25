package check

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sodefrin/archi-checker/src/archi"
	"github.com/sodefrin/archi-checker/src/parser"
)

func TestRun(t *testing.T) {
	testCases := map[string]struct {
		haveDeps    *archi.Dependencies
		haveIps     []*parser.Import
		wantResults []*Result
	}{
		"normal": {
			haveDeps: &archi.Dependencies{
				LayerMap: archi.LayerMap{
					"a": &archi.Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &archi.Layer{
						Name: "b", Pkgs: []string{"b"},
					},
				},
				Dependencies: []*archi.Dependency{
					{
						From: &archi.Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &archi.Layer{
							Name: "a", Pkgs: []string{"b"},
						},
					},
				},
			},
			haveIps: []*parser.Import{
				{
					From: "a",
					To:   "b",
				},
			},
			wantResults: []*Result{},
		},
		"child_pkg": {
			haveDeps: &archi.Dependencies{
				LayerMap: archi.LayerMap{
					"a": &archi.Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &archi.Layer{
						Name: "b", Pkgs: []string{"b"},
					},
				},
				Dependencies: []*archi.Dependency{
					{
						From: &archi.Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &archi.Layer{
							Name: "b", Pkgs: []string{"b"},
						},
					},
				},
			},
			haveIps: []*parser.Import{
				{
					From: "a",
					To:   "b/c",
				},
			},
			wantResults: []*Result{},
		},
		"non_target_pkg": {
			haveDeps: &archi.Dependencies{
				LayerMap: archi.LayerMap{
					"a": &archi.Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &archi.Layer{
						Name: "b", Pkgs: []string{"b"},
					},
				},
				Dependencies: []*archi.Dependency{
					{
						From: &archi.Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &archi.Layer{
							Name: "b", Pkgs: []string{"c"},
						},
					},
				},
			},
			haveIps: []*parser.Import{
				{
					From: "a",
					To:   "b/c",
				},
			},
			wantResults: []*Result{},
		},
		"undefined_dependency": {
			haveDeps: &archi.Dependencies{
				LayerMap: archi.LayerMap{
					"a": &archi.Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &archi.Layer{
						Name: "b", Pkgs: []string{"b"},
					},
					"c": &archi.Layer{
						Name: "c", Pkgs: []string{"c"},
					},
				},
				Dependencies: []*archi.Dependency{
					{
						From: &archi.Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &archi.Layer{
							Name: "b", Pkgs: []string{"b"},
						},
					},
				},
			},
			haveIps: []*parser.Import{
				{
					From: "a",
					To:   "c",
				},
			},
			wantResults: []*Result{
				{},
			},
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			results := run(v.haveDeps, v.haveIps)
			if diff := cmp.Diff(v.wantResults, results); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
