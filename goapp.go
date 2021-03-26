package goapp

type GoApp struct {
	Name string
	Version string
}

// NewGoApp returns a new CLI instance with sensible defaults.
func NewGoApp(app, version string) *GoApp{
	return &GoApp{
		Name:         app,
		Version:      version,
	}

}
