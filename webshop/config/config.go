package config

type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Upload Upload `mapstructure:"upload"`
	Jwt    Jwt    `mapstructure:"jwt"`
}

type Server struct {
	Port string `mapstructure:"port"`
}

type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

type Upload struct {
	SavePath  string `mapstructure:"savePath"`
	AccessUrl string `mapstructure:"accessUrl"`
}

type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}
