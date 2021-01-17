package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Build will return our application
func Build() *cli.App {
	app := cli.NewApp()
	app.Name = "Find Ip's"
	app.Usage = "This application will find Ip's"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "127.0.0.1",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find ip by host",
			Flags:  flags,
			Action: findIPByHost,
		},

		{
			Name:   "sv",
			Usage:  "Find server name by host",
			Flags:  flags,
			Action: findSVByHost,
		},
	}

	return app
}

func findIPByHost(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println("IP:", ip)
	}
}

func findSVByHost(c *cli.Context) {
	host := c.String("host")

	svs, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, sv := range svs {
		fmt.Println("SERVER:", sv.Host)
	}
}
