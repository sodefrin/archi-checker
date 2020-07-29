package main

func check(deps *Dependencies, ips []*Import) []*Import {
	ret := []*Import{}

	for _, ip := range ips {
		if !isTarget(deps, ip) {
			continue
		}

		if !isValidDependency(deps, ip) {
			ret = append(ret, ip)
		}
	}

	return ret
}

func isTarget(deps *Dependencies, ip *Import) bool {
	if deps.LayerMap.Exist(ip.From) && deps.LayerMap.Exist(ip.To) {
		return deps.LayerMap.GetLayer(ip.From).Name != deps.LayerMap.GetLayer(ip.To).Name
	}

	return false
}

func isValidDependency(deps *Dependencies, ip *Import) bool {
	for _, dep := range deps.Dependencies {
		if dep.From.Exist(ip.From) && dep.To.Exist(ip.To) {
			return true
		}
	}

	return false
}
