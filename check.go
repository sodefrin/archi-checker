package main

import "strings"

func check(a *Architecture, ips []*Import) ([]*Import, []*Import) {
	errImport := []*Import{}
	unknownImport := []*Import{}

	for _, ip := range ips {
		if !a.Contain(ip) {
			if !isOfficialPkg(ip.To) {
				unknownImport = append(unknownImport, ip)
			}
			continue
		}

		if !a.Valid(ip) {
			errImport = append(errImport, ip)
		}
	}

	return errImport, unknownImport
}

func isOfficialPkg(pkg string) bool {
	return !strings.ContainsRune(pkg, '.')
}
