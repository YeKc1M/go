package model

type App struct {
	Cluster   Cluster  `json:"cluster"`
	Namespace string   `json:"namespace"`
	Name      string   `json:"name"`
	Database  []string `json:"database"`
}

type Cluster struct {
	Name       string `json:"name"`
	Kubeconfig string `json:"kubeconfig"`
}
