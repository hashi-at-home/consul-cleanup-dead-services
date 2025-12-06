/* Consul dead service deregisterer */

package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

	// Start a context to keep the service alive
	ctx := context.Background()
	defer ctx.Done()

	// create a Consul API consumer to handle agents
	if err != nil {
		log.Fatalf("Failed to create Consul agent consumer: %s", err)
		return err
	}

	// Create a channel to receive OS signals which will shut the service down
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine to handle signals
	go func() {
		sig := <-signalChan
		log.Infof("Received signal: %s", sig)
		os.Exit(0)
	}()

	return nil
}
