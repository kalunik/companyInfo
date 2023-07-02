package config

import (
	"github.com/caarlos0/env/v8"
)

type Config struct {
	Grpc *GrpcServer
	Gw   *Gateway
}

type GrpcServer struct {
	Host string `env:"HOST_SERVER"`
	Port string `env:"PORT_SERVER"`
}

type Gateway struct {
	Host string `env:"HOST_GATEWAY"`
	Port string `env:"PORT_GATEWAY"`
}

func (c *Config) ParseConfig() error {
	c.Grpc = new(GrpcServer)
	c.Gw = new(Gateway)

	err := env.Parse(c)

	if err != nil {
		return err
	}
	return nil
}
