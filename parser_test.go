package main

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParsePkgs(t *testing.T) {
	if err := os.Chdir("./testdata/src/xxx.xxx"); err != nil {
		t.Fatal(err)
	}

	ret, err := ParsePkgs("xxxx.xxx", true, "xxx.xxx/a", "xxx.xxx/b")
	if err != nil {
		t.Fatal(err)
	}

	want := []*Import{
		{From: "xxx.xxx/a", To: "fmt"},
		{From: "xxx.xxx/b", To: "xxx.xxx/a"},
	}

	sortOpt := cmpopts.SortSlices(func(a, b *Import) bool {
		if a.From > b.From {
			return true
		} else if a.From < b.From {
			return false
		}

		if a.To > b.To {
			return true
		}
		return false
	})

	opt := cmpopts.IgnoreFields(Import{}, "FileSet", "File", "Import")
	if diff := cmp.Diff(want, ret, sortOpt, opt); diff != "" {
		t.Fatal(diff)
	}
}
