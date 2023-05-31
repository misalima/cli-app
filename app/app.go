package app

import (
	"log"
	"net"
	"fmt"
	"github.com/urfave/cli"
)

//This function will return the CLI app ready to be executed
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command Line Application"
	app.Usage = "Gets IP numbers and server names from the web"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "devbook.com.br",
		},
	}
	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Gets IPs from the web",
			Flags: flags,
			Action: getIps,
		},
		{
			Name: "servers",
			Usage: "Gets server names from the web",
			Flags: flags,
			Action: getServers,

		},
	}

	return app
}

func getIps(c *cli.Context) {
			host := c.String("host")
			
			ips, err := net.LookupIP(host)
			if err != nil {
				log.Fatal(err)
			}
			for _, ip := range ips {
				fmt.Println(ip)
			}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}