package globals

type Config struct {
	DB struct {
		Type   string      `yaml:"type"`
		Config MysqlConfig `yaml:"config"`
	}
	Redis struct {
		Config RedisConfig `yaml:"redis"`
	}
	App App `yaml:"app"`
	Jwt Jwt `yaml:"jwt"`
}

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	MaxCon   int    `yaml:"maxCon"`
	MaxIdle  int    `yaml:"maxIdle"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Db       int    `yaml:"db"`
	PoolSize int    `yaml:"poolSize"`
}

type App struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Jwt struct {
	Secret string `yaml:"secret"`
}
