package supported

import "go.uber.org/fx"

var allMarkets = []string{"gb", "ie", "it", "fr", "be", "ae", "kw", "qa", "hk", "sg"}

type Config struct {
	AllMarkets []string
}

func NewConfig() *Config {
	return &Config{
		AllMarkets: allMarkets,
	}
}

func (c *Config) GetAllSupportedMarkets() []string {
	return c.AllMarkets
}

var Module = fx.Options(
	fx.Provide(NewConfig),
)
