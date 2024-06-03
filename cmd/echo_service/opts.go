package main

import (
	// Built-in/core modules.
	"flag"
	"fmt"
)

type Config struct {
	Listen string
}

func parse_cmd_opts() (*Config, error) {
	conf := &Config{}

	flag.StringVar(&conf.Listen, "listen", "127.0.0.1:9000",
		"IP address and port for the gRPC/REST to listen on")

	flag.Parse()

	if conf.Listen == "" {
		return nil, fmt.Errorf("`restlisten` parameter required")
	}

	return conf, nil
}
