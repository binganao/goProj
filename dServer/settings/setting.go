package settings

import "flag"

var (
	Room    string
	Debug   bool
	Port    string
	Timeout int
)

func ReadFlags() {
	flag.StringVar(&Port, "port", "8099", "")
	flag.BoolVar(&Debug, "debug", false, "")
	flag.StringVar(&Room, "room", "545068", "number")
	flag.IntVar(&Timeout, "timeout", 5, "")
}
