package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
)

type context struct {
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
}

type contextinfo struct {
	Name    string  `json:"name"`
	Context context `json:"context"`
}

type config struct {
	Clusters []contextinfo `json:"contexts"`
	Current  string        `json:"current-context"`
}

func main() {
	cmd := exec.Command("kubectl", "config", "view", "-o", "json")
	out, err := cmd.Output()
	cfg := config{}
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(out, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	context := cfg.Current
	ns := "default"
	for _, x := range cfg.Clusters {
		if x.Name == context {
			ns = x.Context.Namespace
			if ns == "" {
				ns = "default"
			}
		}
	}
	fmt.Printf("%s/%s", context, ns)
}
