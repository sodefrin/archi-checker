package main

import "testing"

func TestMain(t *testing.T) {
	testCases := map[string]struct {
		haveConfig config
		wantRet    int
	}{
		"this_package": {
			haveConfig: config{
				umlPath:    "uml/this_package.uml",
				modulePath: "github.com/sodefrin/archi-checker",
				pkgs:       []string{".", "src/archi", "src/check", "src/parser"},
			},
			wantRet: 0,
		},
		"dip_this_package": {
			haveConfig: config{
				umlPath:    "uml/this_package_dip.uml",
				modulePath: "github.com/sodefrin/archi-checker",
				pkgs:       []string{".", "src/archi", "src/check", "src/parser"},
			},
			wantRet: 1,
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			ret := run(v.haveConfig)
			if ret != v.wantRet {
				t.Fatalf("invalid exit code: want %d but have %d", v.wantRet, ret)
			}
		})
	}
}
