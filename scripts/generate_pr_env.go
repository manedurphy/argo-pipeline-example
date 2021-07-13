package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var (
	commit = flag.String("commit", "", "defines the commit hash for the pull request")
	ns     string
	name   string
	app    string
)

type Config struct {
	ApiVersion string `yaml:"apiVersion"`
	Kind       string
	Metadata   map[string]interface{}
	Spec       map[string]interface{}
}

func main() {
	flag.Parse()

	name = "ping-pong-deployment-" + *commit
	app = "ping-pong-" + *commit
	ns = "development"
	filename := "k8s/development/development.yaml"

	namespace, _ := generateNamespace()
	deployment, _ := generateDeployment()

	err := os.WriteFile(filename, namespace, 0644)

	if err != nil {
		log.Fatalf("could not write development yaml file: %v", err)
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		log.Fatalf("could not open development yaml file: %v", err)
	}

	defer f.Close()

	if _, err := f.WriteString("---\n"); err != nil {
		fmt.Println("Panic on line 49")
		panic(err)
	}

	if _, err := f.WriteString(string(deployment)); err != nil {
		fmt.Println("Panic on line 54")
		panic(err)
	}

}

func generateNamespace() ([]byte, error) {

	namespace := Config{
		ApiVersion: "v1",
		Kind:       "Namespace",
		Metadata: map[string]interface{}{
			"name": ns,
		},
	}

	yml, err := yaml.Marshal(namespace)

	if err != nil {
		return nil, err
	}

	return yml, nil
}

func generateDeployment() ([]byte, error) {

	d := Config{
		ApiVersion: "apps/v1",
		Kind:       "Deployment",
		Metadata: map[string]interface{}{
			"name":      name,
			"namespace": ns,
			"labels": map[string]string{
				"app": app,
			},
		},
		Spec: map[string]interface{}{
			"replicas": 1,
			"selector": map[string]interface{}{
				"matchLabels": map[string]string{
					"app": app,
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]string{
						"app": app,
					},
				},
				"spec": map[string]interface{}{
					"containers": []map[string]interface{}{
						{
							"name":            app,
							"image":           "manedurphy/ping-pong:development",
							"imagePullPolicy": "Always",
							"ports": []map[string]int{
								{
									"containerPort": 8080,
								},
							},
							"livenessProbe": map[string]interface{}{
								"httpGet": map[string]interface{}{
									"path": "/healthz",
									"port": 8080,
								},
								"initialDelaySeconds": 3,
								"periodSeconds":       3,
							},
							"startupProbe": map[string]interface{}{
								"httpGet": map[string]interface{}{
									"path": "/healthz",
									"port": 8080,
								},
								"failureThreshold": 30,
								"periodSeconds":    10,
							},
						},
					},
				},
			},
		},
	}

	yml, err := yaml.Marshal(d)

	if err != nil {
		return nil, err
	}

	return yml, nil
}
