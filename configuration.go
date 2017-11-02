package gonfig

// IConfiguration configuration instance
type IConfiguration interface {
	Get(key string) string
}

type configuration struct {
}
