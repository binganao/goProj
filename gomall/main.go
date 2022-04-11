package main

import "./initialize"

func main() {
	initialize.LoadConfig()
	initialize.Mysql()
	initialize.Redis()
	initialize.Elastic()
	initialize.Router()
}
