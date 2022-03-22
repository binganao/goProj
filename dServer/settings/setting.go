package settings

import "flag"

var (
	Port string
)

func ReadFlags() {
	flag.StringVar(&Port, "port", "8080", "")
}
