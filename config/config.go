package config

type Config struct {
	grpc *GrpcServer
	gw   *Gateway
}

type GrpcServer struct {
	Host string `env:"host"`
	Port string `env:"port"`
}

type Gateway struct {
	Host string `env:"host"`
	Port string `env:"port"`
}

func (c *Config) ParseConfig(filePath string) error {
	c = new(Config)

}
