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

	app.Commands = []cli.Command {
		{
			Name: "ip",
			Usage: "Gets IPs from the web",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "devbook.com.br",
				},
			},
			Action: getIps,
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