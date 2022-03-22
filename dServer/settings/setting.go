package settings

import "flag"

var (
	Port    string
	Timeout int
)

func ReadFlags() {
	flag.StringVar(&Port, "port", "8099", "")
	flag.IntVar(&Timeout, "timeout", 5, "")
}
