package config

type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Upload Upload `mapstructure:"upload"`
	Jwt    Jwt    `mapstructure:"jwt"`
}

type Server struct {
	Port        string `mapstructure:"port"`
	ReleaseMode bool   `mapstructure:"release_mode"`
}

type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

type Redis struct {
	Addr string `mapstructure:"host"`
}

type Upload struct {
	SavePath  string `mapstructure:"savePath"`
	AccessUrl string `mapstructure:"accessUrl"`
}

type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}
