package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/yggdrasil-network/yggdrasil-go/src/config"
	"github.com/yggdrasil-network/yggdrasil-go/src/core"
)

// LoggerAdapter adapts *log.Logger to core.Logger interface
type LoggerAdapter struct {
	*log.Logger
}

// Implement additional methods required by core.Logger
func (la *LoggerAdapter) Debugf(format string, v ...interface{}) {
	la.Printf("DEBUG: "+format, v...)
}

func (la *LoggerAdapter) Debugln(v ...interface{}) {
	la.Println(append([]interface{}{"DEBUG:"}, v...)...)
}

func (la *LoggerAdapter) Infof(format string, v ...interface{}) {
	la.Printf("INFO: "+format, v...)
}

func (la *LoggerAdapter) Infoln(v ...interface{}) {
	la.Println(append([]interface{}{"INFO:"}, v...)...)
}

func (la *LoggerAdapter) Warnf(format string, v ...interface{}) {
	la.Printf("WARN: "+format, v...)
}

func (la *LoggerAdapter) Warnln(v ...interface{}) {
	la.Println(append([]interface{}{"WARN:"}, v...)...)
}

func (la *LoggerAdapter) Errorf(format string, v ...interface{}) {
	la.Printf("ERROR: "+format, v...)
}

func (la *LoggerAdapter) Errorln(v ...interface{}) {
	la.Println(append([]interface{}{"ERROR:"}, v...)...)
}

func (la *LoggerAdapter) Traceln(v ...interface{}) {
	la.Println(append([]interface{}{"TRACE:"}, v...)...)
}

func main() {
	// Generate a new configuration
	cfg := config.GenerateConfig()

	// Use the generated certificate from the configuration
	cert := cfg.Certificate

	// Create a new Yggdrasil node with the adapted logger
	node, err := core.New(cert, &LoggerAdapter{log.Default()})
	u, _ := url.Parse("tls://192.9.143.104:443")
	node.AddPeer(u, "")

	if err != nil {
		log.Fatalf("Failed to create Yggdrasil node: %v", err)
	}

	defer node.Stop()

	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Example response using the node's address, assuming a method to fetch it
		fmt.Fprintf(w, "Hello from Yggdrasil: %s", node.Address())
	})

	// Start the HTTP server on localhost:8383
	fmt.Println("Server starting on http://localhost:8383/")
	if err := http.ListenAndServe(":8383", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
