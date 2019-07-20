package config

import (
	"github.com/spf13/cast"
)

const (
	logLevel      = "log.level"
	httpPort      = "http.port"
	apiGatewayURL = "api.gateway.url"
)

type Config struct {
	config map[string]interface{}
}

func New() *Config {
	c := new(Config)
	c.config = make(map[string]interface{})
	c.setDefaults()
	return c
}

func (c *Config) setDefaults() {
	c.SetLogLevel("debug")
	c.SetHttpPort(8080)
	c.SetAPIGatewayURL("http://api-gateway:8070/")
}

func (c *Config) GetLogLevel() string {
	return cast.ToString(c.config[logLevel])
}

func (c *Config) SetLogLevel(level string) {
	c.config[logLevel] = level
}

func (c *Config) GetHttpPort() int {
	return cast.ToInt(c.config[httpPort])
}

func (c *Config) SetHttpPort(port int) {
	c.config[httpPort] = port
}

func (c *Config) GetAPIGatewayURL() string {
	return cast.ToString(c.config[apiGatewayURL])
}

func (c *Config) SetAPIGatewayURL(url string) {
	c.config[apiGatewayURL] = url
}
