package config

var config = map[string]string{
	"DB_PORT":     "5432",
	"DB_USERNAME": "postgres",
	"DB_PASSWORD": "postgres",
	"DB_HOST":     "localhost",
	"DB_NAME":     "postgres",
	"DB_SSL_MODE":  "disable",
	"DB_DRIVER":   "postgres",
}

func Config(key string) string {
	return config[key]
}
