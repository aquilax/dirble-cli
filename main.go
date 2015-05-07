package main

import (
	"encoding/json"
	"fmt"
	"github.com/aquilax/go-dirble"
	"github.com/codegangsta/cli"
	"net/http"
	"os"
)

const (
	appName    = "dirble-cli"
	appVersion = "0.0.1"
	defaultInt = -1
)

func getDirble(token string) *dirble.Dirble {
	tr := http.Transport{}
	return dirble.New(&tr, token)
}

func processResult(d interface{}, err error) {
	var res []byte
	if err != nil {
		panic(err)
	}
	if res, err = json.MarshalIndent(d, "", "	"); err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", res)
}

func intToParam(c *cli.Context, name string) *int {
	if c.IsSet(name) {
		result := c.Int(name)
		return &result
	}
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = appName
	app.Version = appVersion
	app.Usage = "Fetches information from dirble.com"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "API Token",
			EnvVar: "DIRBLE_API_TOKEN",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "stations",
			Aliases: []string{"st"},
			Usage:   "Get List of stations",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "page",
					Usage: "page to fetch",
				},
				cli.IntFlag{
					Name:  "ipp",
					Usage: "items per page",
				},
				cli.IntFlag{
					Name:  "offset",
					Usage: "offset",
				},
			},
			Action: func(c *cli.Context) {
				processResult(getDirble(c.GlobalString("token")).Stations(intToParam(c, "page"),
					intToParam(c, "ipp"), intToParam(c, "offset")))
			},
		},
	}
	app.Run(os.Args)
}