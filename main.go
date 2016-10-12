package main

import (
	"bytes"
	"fmt"
	"os"
	"os/signal"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/gigawattio/go-commons/pkg/upstart"
	"github.com/jaytaylor/tesseract-web/interfaces"
	"gopkg.in/urfave/cli.v2"
)

const (
	AppName = "tesseract-web"
)

var (
	webService *interfaces.WebService
)

func runWeb(addr string) error {
	webService = interfaces.NewWebService(addr)
	if err := webService.Start(); err != nil {
		return err
	}
	log.Printf("Successfully started web service on addr=%v\n", webService.Addr())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	<-sig // Wait for ^C signal
	fmt.Fprintln(os.Stderr, "\nInterrupt signal detected, shutting down..")

	if err := webService.Stop(); err != nil {
		return err
	}

	return nil
}

func errorExit(err error, exitStatusCode int) {
	os.Stderr.WriteString("error: " + err.Error() + "\n")
	os.Exit(exitStatusCode)
}

func getVersion() string {
	buf := &bytes.Buffer{}
	DisplayVersion(buf, "\n")
	v := strings.TrimSpace(buf.String())
	return v
}

func main() {
	var (
		install     bool
		uninstall   bool
		serviceUser string = os.Getenv("USER")
		web         bool
		bindAddr    string = "0.0.0.0:8080"
	)

	app := &cli.App{
		Name:    AppName,
		Version: getVersion(),
		Usage:   "Exposes tesseract image OCR as a set of web APIs",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "install",
				Usage:       fmt.Sprintf("Install %s as a system service", AppName),
				Destination: &install,
			},
			&cli.BoolFlag{
				Name:        "uninstall",
				Usage:       fmt.Sprintf("Uninstall %s as a system service", AppName),
				Destination: &uninstall,
			},
			&cli.StringFlag{
				Name:        "user",
				Aliases:     []string{"u"},
				Usage:       fmt.Sprintf("Specifies the user to run the %s system service as (will default to %q)", AppName, serviceUser),
				Value:       serviceUser,
				Destination: &serviceUser,
			},
			&cli.BoolFlag{
				Name:        "web",
				Aliases:     []string{"w"},
				Usage:       "Starts in web-server mode",
				Destination: &web,
			},
			&cli.StringFlag{
				Name:        "bind",
				Aliases:     []string{"b"},
				Usage:       "Set the web-server bind-address and (optionally) port",
				Value:       bindAddr,
				Destination: &bindAddr,
			},
		},
		Action: func(c *cli.Context) error {
			log.Infof("c.Args=%v c.FlagNames=%v", c.Args(), c.FlagNames())
			if uninstall || install {
				if uninstall {
					config := upstart.DefaultConfig(c.App.Name)
					if err := upstart.UninstallService(config); err != nil {
						return err
					}
				}
				if install {
					config := upstart.DefaultConfig(c.App.Name)
					config.User = serviceUser
					if err := upstart.InstallService(config); err != nil {
						return err
					}
				}
				return nil
			}

			if err := runWeb(bindAddr); err != nil {
				return err
			}
			return nil
		},
	}
	if err := app.Run(os.Args); err != nil {
		errorExit(err, 1)
	}
}
