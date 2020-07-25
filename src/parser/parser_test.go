package parser

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestParsePkgs(t *testing.T) {
	// testing from pkg root
	if err := os.Chdir("../../"); err != nil {
		t.Fatal(err)
	}

	ret, err := ParsePkgs("github.com/sodefrin/archi-checker", "src/parser")
	if err != nil {
		t.Fatal(err)
	}

	want := []*Import{
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "os"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "testing"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "github.com/google/go-cmp/cmp"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "github.com/google/go-cmp/cmp/cmpopts"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "fmt"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "go/parser"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "go/token"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "strings"},
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

	if diff := cmp.Diff(want, ret, sortOpt); diff != "" {
		t.Fatal(diff)
	}
}
