package main

import "mall.com/initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Redis()
	initialize.Elastic()
	initialize.Router()
}
