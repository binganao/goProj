package settings

import (
	"flag"
	"fmt"
	"os"
)

var (
	Room    string
	Debug   bool
	Path    string
	Port    string
	Store   string
	Timeout int
)

func ReadFlags() {
	flag.StringVar(&Path, "path", "/blive", "base path")
	flag.StringVar(&Port, "port", "8099", "")
	flag.BoolVar(&Debug, "debug", false, "")
	flag.StringVar(&Room, "room", "545068", "number")
	flag.StringVar(&Store, "store", "", "string stored in /?store")
	flag.IntVar(&Timeout, "timeout", 5, "")
}

func StringFlags() string {
	return fmt.Sprintf(`%s -path "%v" -port %v -debug=%v -room %v -store "%v" -timeout %v`,
		os.Args[0], Path, Port, Debug, Room, Store, Timeout)
}
