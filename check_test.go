package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheck(t *testing.T) {
	testCases := map[string]struct {
		haveDeps    *Dependencies
		haveIps     []*Import
		wantResults []*Import
	}{
		"normal": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"b"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &Layer{
							Name: "a", Pkgs: []string{"b"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "a",
					To:   "b",
				},
			},
			wantResults: []*Import{},
		},
		"child_pkg": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"b"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &Layer{
							Name: "b", Pkgs: []string{"b"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "a",
					To:   "b/c",
				},
			},
			wantResults: []*Import{},
		},
		"non_target_pkg": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"b"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &Layer{
							Name: "b", Pkgs: []string{"c"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "a",
					To:   "c",
				},
			},
			wantResults: []*Import{},
		},
		"undefined_dependency": {
			haveDeps: &Dependencies{
				LayerMap: LayerMap{
					"a": &Layer{
						Name: "a", Pkgs: []string{"a"},
					},
					"b": &Layer{
						Name: "b", Pkgs: []string{"b"},
					},
					"c": &Layer{
						Name: "c", Pkgs: []string{"c"},
					},
				},
				Dependencies: []*Dependency{
					{
						From: &Layer{
							Name: "a", Pkgs: []string{"a"},
						},
						To: &Layer{
							Name: "b", Pkgs: []string{"b"},
						},
					},
				},
			},
			haveIps: []*Import{
				{
					From: "a",
					To:   "c",
				},
			},
			wantResults: []*Import{
				{
					From: "a",
					To:   "c",
				},
			},
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			results := check(v.haveDeps, v.haveIps)
			if diff := cmp.Diff(v.wantResults, results); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
