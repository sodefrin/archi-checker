package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheck(t *testing.T) {
	testCases := map[string]struct {
		haveDeps           *Dependencies
		haveIps            []*Import
		wantInvalidImports []*Import
		wantUnknownImports []*Import
	}{
		"normal": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"xxx.xxx/a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"xxx.xxx/b"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"xxx.xxx/a"},
						},
						To: &Layer{
							Name: "a", Pkgs: []string{"xxx.xxx/b"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/b",
				},
			},
			wantInvalidImports: []*Import{},
			wantUnknownImports: []*Import{},
		},
		"child_pkg": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"xxx.xxx/a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"xxx.xxx/b"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"xxx.xxx/a"},
						},
						To: &Layer{
							Name: "b", Pkgs: []string{"xxx.xxx/b"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/b/c",
				},
			},
			wantInvalidImports: []*Import{},
			wantUnknownImports: []*Import{},
		},
		"unknown": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"xxx.xxx/a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"xxx.xxx/b"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"xxx.xxx/a"},
						},
						To: &Layer{
							Name: "b", Pkgs: []string{"xxx.xxx/c"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/c",
				},
			},
			wantInvalidImports: []*Import{},
			wantUnknownImports: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/c",
				},
			},
		},
		"invalid": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"xxx.xxx/a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"xxx.xxx/b"},
					},
					"c": &Layer{
						Name: "c", Pkgs: []string{"xxx.xxx/c"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"xxx.xxx/a"},
						},
						To: &Layer{
							Name: "b", Pkgs: []string{"xxx.xxx/b"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/c",
				},
			},
			wantInvalidImports: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/c",
				},
			},
			wantUnknownImports: []*Import{},
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			invalids, unknowns := check(v.haveDeps, v.haveIps)
			if diff := cmp.Diff(v.wantInvalidImports, invalids); diff != "" {
				t.Fatal(diff)
			}
			if diff := cmp.Diff(v.wantUnknownImports, unknowns); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
