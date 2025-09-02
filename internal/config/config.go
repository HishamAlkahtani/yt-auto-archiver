package config

type Config struct {
	dbPath string
}

func (c *Config) GetDbPath() string {
	return c.dbPath
}
