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
		"success": {
			haveDeps:    &archi.Dependencies{},
			haveIps:     []*parser.Import{},
			wantResults: []*Result{},
		},
	}

	for _, v := range testCases {
		results := run(v.haveDeps, v.haveIps)
		if diff := cmp.Diff(v.wantResults, results); diff != "" {
			t.Fatal(diff)
		}
	}
}
