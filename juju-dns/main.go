package main

import (
	"fmt"
	"os"

	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/coremain"
	_ "github.com/nvinuesa/juju-coredns-plugin"
)

var directives = []string{
	"juju",
	"whoami",
	"startup",
	"shutdown",
}

func init() {
	dnsserver.Directives = directives
}

func main() {
	fmt.Printf("config file path %q\n", os.Getenv("JUJU_DNS_PLUGIN_CONF_PATH"))
	coremain.Run()
}
