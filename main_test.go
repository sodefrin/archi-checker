package main

import "testing"

func TestMain(t *testing.T) {
	run(config{
		umlPath:    "testdata/sample.uml",
		modulePath: "github.com/sodefrin/archi-checker",
		pkgs:       []string{".", "src/archi", "src/check", "src/parser"},
	})
}
