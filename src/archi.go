package src

type Layer struct {
	Name string
	Pkgs []string
}

type Dependency struct {
	From *Layer
	To   *Layer
}
