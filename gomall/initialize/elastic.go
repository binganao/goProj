package initialize

import (
	"log"

	"./global"
	"github.com/olivere/elastic/v7"
)

func Elastic() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200/"))
	if err != nil {
		log.Panicln(err)
	}
	global.Es = client
}
