package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func App() {
	var paths struct{ tempDir, homeDir, workingDir string }
	var httpPort int
	var logLevel int

	app := &cli.App{
		Name:    "dstack-runner",
		Usage:   "configure and start dstack-runner",
		Version: Version,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "log-level",
				Value:       2,
				DefaultText: "4 (Info)",
				Usage:       "log verbosity level: 2 (Error), 3 (Warning), 4 (Info), 5 (Debug), 6 (Trace)",
				Destination: &logLevel,
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "Start dstack-runner",
				Flags: []cli.Flag{
					&cli.PathFlag{
						Name:        "temp-dir",
						Usage:       "Temporary directory for logs and other files",
						Required:    true,
						Destination: &paths.tempDir,
					},
					&cli.PathFlag{
						Name:        "home-dir",
						Usage:       "Home directory for credentials and $HOME",
						Required:    true,
						Destination: &paths.homeDir,
					},
					&cli.PathFlag{
						Name:        "working-dir",
						Usage:       "Base path for the job",
						Required:    true,
						Destination: &paths.workingDir,
					},
					&cli.IntFlag{
						Name:        "http-port",
						Usage:       "Set a http port",
						Value:       10999,
						Destination: &httpPort,
					},
				},
				Action: func(c *cli.Context) error {
					start(paths.tempDir, paths.homeDir, paths.workingDir, httpPort, logLevel)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}