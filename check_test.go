package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCheck(t *testing.T) {
	testCases := map[string]struct {
		haveArchi          *Architecture
		haveIps            []*Import
		wantInvalidImports []*Import
		wantUnknownImports []*Import
	}{
		"normal": {
			haveArchi: &Architecture{
				Packages: map[string]string{
					"xxx.xxx/a": "a",
					"xxx.xxx/b": "b",
				},
				Dependencies: []*Dependency{
					{
						FromLayer: "a",
						ToLayer:   "b",
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
			haveArchi: &Architecture{
				Packages: map[string]string{
					"xxx.xxx/a": "a",
					"xxx.xxx/b": "b",
				},
				Dependencies: []*Dependency{
					{
						FromLayer: "a",
						ToLayer:   "b",
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
			haveArchi: &Architecture{
				Packages: map[string]string{
					"xxx.xxx/a": "a",
					"xxx.xxx/b": "b",
				},
				Dependencies: []*Dependency{
					{
						FromLayer: "a",
						ToLayer:   "b",
					},
				},
			},
			haveIps: []*Import{
				{
					From: "xxx.xxx/a",
					To:   "xxx.xxx/c",
				},
				{
					From: "xxx.xxx/a",
					To:   "fmt",
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
			haveArchi: &Architecture{
				Packages: map[string]string{
					"xxx.xxx/a": "a",
					"xxx.xxx/b": "b",
				},
				Dependencies: []*Dependency{
					{
						FromLayer: "a",
						ToLayer:   "b",
					},
				},
			},
			haveIps: []*Import{
				{
					From: "xxx.xxx/b",
					To:   "xxx.xxx/a",
				},
			},
			wantInvalidImports: []*Import{
				{
					From: "xxx.xxx/b",
					To:   "xxx.xxx/a",
				},
			},
			wantUnknownImports: []*Import{},
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			invalids, unknowns := check(v.haveArchi, v.haveIps)
			if diff := cmp.Diff(v.wantInvalidImports, invalids); diff != "" {
				t.Fatal(diff)
			}
			if diff := cmp.Diff(v.wantUnknownImports, unknowns); diff != "" {
				t.Fatal(diff)
			}
		})
	}
}
