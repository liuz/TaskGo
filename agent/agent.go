package agent

import (
	"context"
	"net/http"
)

type Agent struct {
	endPoint    *http.Server
	contentType string
	//statsdServer StatsdServer
}

func NewAgent(config *Config) (*Agent, error) {
	return NewAgentContext(context.Background(), config)
}
