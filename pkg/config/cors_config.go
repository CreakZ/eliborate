package config

type CorsConfig struct {
	AccessControlAllowOrigin  string
	AccessControlAllowMethods string
	AccessControlAllowHeaders string
}

func NewCorsConfig(origin, methods, headers string) *CorsConfig {
	return &CorsConfig{
		AccessControlAllowOrigin:  origin,
		AccessControlAllowMethods: methods,
		AccessControlAllowHeaders: headers,
	}
}
