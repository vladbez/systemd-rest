package docker

type APIHistory struct {
	ID        string `json:"Id"`
	Created   int64
	CreatedBy string `json:",omitempty"`
}

type APIImages struct {
	Repository string `json:",omitempty"`
	Tag        string `json:",omitempty"`
	ID         string `json:"Id"`
	Created    int64
}

type APIInfo struct {
	Debug       bool
	Containers  int
	Images      int
	NFd         int  `json:",omitempty"`
	NGoroutines int  `json:",omitempty"`
	MemoryLimit bool `json:",omitempty"`
	SwapLimit   bool `json:",omitempty"`
}

type APIRmi struct {
	Deleted  string `json:",omitempty"`
	Untagged string `json:",omitempty"`
}

type APIContainers struct {
	ID      string `json:"Id"`
	Image   string
	Command string
	Created int64
	Status  string
	Ports   string
}

type APISearch struct {
	Name        string
	Description string
}

type APIID struct {
	ID string `json:"Id"`
}

type APIRun struct {
	ID       string   `json:"Id"`
	Warnings []string `json:",omitempty"`
}

type APIPort struct {
	Port string
}

type APIVersion struct {
	Version   string
	GitCommit string `json:",omitempty"`
	GoVersion string `json:",omitempty"`
}

type APIWait struct {
	StatusCode int
}

type APIAuth struct {
	Status string
}

type APIImageConfig struct {
	ID string `json:"Id"`
	*Config
}
