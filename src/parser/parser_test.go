package parser

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
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
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "fmt"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "go/parser"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "go/token"},
		{From: "github.com/sodefrin/archi-checker/src/parser", To: "strings"},
	}

	if diff := cmp.Diff(want, ret); diff != "" {
		t.Fatal(diff)
	}
}
