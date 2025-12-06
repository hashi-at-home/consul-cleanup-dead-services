/* Consul dead service deregisterer */

package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/hashicorp/consul/api"
)

func main() {
	if err := run(os.Args[:]); err != nil {
		log.Fatalf("%s", err)
	} else {
		log.Info("Everything cool homie")
	}

}

// run is an function of type void which takes a list of strings as argument
// and returns an error
// run should parse environment variables and start a Consul client
func run(args []string) error {

	// required
	requiredConsulEnvVars := []string{"CONSUL_HTTP_ADDR", "CONSUL_TOKEN"}
	for _, env := range requiredConsulEnvVars {
		if os.Getenv(env) == "" {
			log.Fatalf("%s environment variable is not set", env)
		}
	}

	consulClient, err := api.NewClient(&api.Config{
		Address: os.Getenv("CONSUL_HTTP_ADDR"),
		Token:   os.Getenv("CONSUL_TOKEN"),
	})
	if err != nil {
		log.Fatalf("Failed to create Consul API client: %s", err)
		return err
	}

	// Construct a query that will get all nodes, healthy and not healthy
	nodeQuery := &api.QueryOptions{}
	nodes, _, err := consulClient.Catalog().Nodes(nodeQuery)
	if len(nodes) == 0 || err != nil {
		log.Fatal("No nodes found")
		return err
	}
	for _, node := range nodes {

		log.Infof("Node found: %s", node.Node)
	}

	return nil
}
