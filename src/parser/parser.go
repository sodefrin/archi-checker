package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path"
	"strings"
)

type Import struct {
	FileSet *token.FileSet
	File    *ast.File
	Import  *ast.ImportSpec
	From    string
	To      string
}

func ParsePkgs(modulePath string, pkgs ...string) ([]*Import, error) {
	ips := []*Import{}
	for _, pkg := range pkgs {
		tmp, err := parsePkg(modulePath, pkg)
		if err != nil {
			return nil, err
		}
		ips = append(ips, tmp...)
	}
	return ips, nil
}

func parsePkg(modulePath string, pkg string) ([]*Import, error) {
	path := path.Join(modulePath, pkg)
	fset := token.NewFileSet()
	ast, err := parser.ParseDir(fset, pkg, nil, parser.Mode(0))
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
