package config

type Config struct {
	AuthSecret string
	BCryptCost uint
	TokenTTL   uint

	CORS map[string]struct{}

	NotTLSWhiteList map[string]struct{}
}
