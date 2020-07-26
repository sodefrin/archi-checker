package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

type Import struct {
	FileSet *token.FileSet
	File    *ast.File
	Import  *ast.ImportSpec
	From    string
	To      string
}

func ParsePkgs(modulePath string, paths ...string) ([]*Import, error) {
	ips := []*Import{}
	for _, path := range paths {
		tmp, err := parsePkg(modulePath, path)
		if err != nil {
			return nil, err
		}
		ips = append(ips, tmp...)
	}
	return ips, nil
}

func parsePkg(modulePath, path string) ([]*Import, error) {
	fset := token.NewFileSet()
	rel, err := filepath.Rel(modulePath, path)
	if err != nil {
		return nil, err
	}
	ast, err := parser.ParseDir(fset, rel, nil, parser.Mode(0))
	if err != nil {
		return nil, err
	}

	ips := []*Import{}
	for _, pkg := range ast {
		for _, file := range pkg.Files {
			for _, ip := range file.Imports {
				ips = append(ips, &Import{
					FileSet: fset,
					File:    file,
					Import:  ip,
					From:    path,
					To:      strings.Trim(ip.Path.Value, "\""),
				})
			}
		}
	}

	return ips, nil
}
