/* Consul dead service deregisterer */

package main

import (
	"os"

	"github.com/charmbracelet/log"
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
	return nil
}
