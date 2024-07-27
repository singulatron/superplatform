package router

type AppConfig struct {
	App       AppMetadata `yaml:"app"`
	Endpoints []Endpoint  `yaml:"endpoints"`
	Hooks     Hooks       `yaml:"hooks"`
}

type AppMetadata struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Author  string `yaml:"author"`
}

type Endpoint struct {
	Path    string `yaml:"path"`
	Method  string `yaml:"method"`
	Handler string `yaml:"handler"`
}

type Hooks struct {
	PreRequest  []Hook `yaml:"preRequest"`
	PostRequest []Hook `yaml:"postRequest"`
}

type Hook struct {
	Event   string `yaml:"event"`
	Handler string `yaml:"handler"`
}
