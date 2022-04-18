package initialize

func Start() {
	LoadConfig()
	Mysql()
	Redis()
	Router()
}
