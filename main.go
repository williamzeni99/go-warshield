package main

import (
	"fmt"
	"ghoji/encryptor"
	"ghoji/graphic"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "ghoji",
		Usage:    "A CLI tool to encrypt and decrypt files",
		Version:  "v1.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "William Zeni",
				Email: "williamzeni56@gmail.com",
			},
		},
		EnableBashCompletion: true,
		Description:          "This is a super fast program for encrypting big files. It implements AES 256 with GCM. Because of the parallelism, the file is deleted after an encrypted copy is made. So, be sure to have enough space in the hard drive when performing an encryption or decryption. In addition, no limit has been set for the power of parallelism, you can set the number of goroutines that can go in parallel. If the size of the file is big enough all of them will be loaded in the ram. IMPORTANT: Do not use too high values or you will have a crash.",
		Commands: []*cli.Command{
			{
				Name:  "encrypt",
				Usage: "Encrypt a file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Aliases:  []string{"p"},
						Usage:    "Path to the file to encrypt",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "numCpu",
						Aliases: []string{"n"},
						Usage:   "Number of CPU cores to use",
						Value:   encryptor.MaxCPUs,
					},
					&cli.IntFlag{
						Name:    "goroutines",
						Aliases: []string{"g"},
						Usage:   "Maximum number of goroutines to use. High values can cause a crash. Try at your own risk",
						Value:   encryptor.DefaultGoRoutines,
					},
				},
				Action: func(c *cli.Context) error {
					path := c.String("path")
					numCpu := c.Int("numCpu")
					goroutines := c.Int("goroutines")

					if numCpu > encryptor.MaxCPUs || numCpu < 0 {
						numCpu = encryptor.MaxCPUs
					}

					if goroutines <= 0 {
						goroutines = encryptor.DefaultGoRoutines
					}

					graphic.DoEncryption(path, numCpu, goroutines)
					return nil
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a file",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Aliases:  []string{"p"},
						Usage:    "Path to the file to decrypt",
						Required: true,
					},
					&cli.IntFlag{
						Name:    "numCpu",
						Aliases: []string{"n"},
						Usage:   "Number of CPU cores to use",
						Value:   encryptor.MaxCPUs,
					},
					&cli.IntFlag{
						Name:    "goroutines",
						Aliases: []string{"g"},
						Usage:   "Maximum number of goroutines to use. High values can cause a crash. Try at your own risk",
						Value:   encryptor.DefaultGoRoutines,
					},
				},
				Action: func(c *cli.Context) error {
					path := c.String("path")
					numCpu := c.Int("numCpu")
					goroutines := c.Int("goroutines")

					if numCpu > encryptor.MaxCPUs || numCpu < 0 {
						numCpu = encryptor.MaxCPUs
					}

					if goroutines <= 0 {
						goroutines = encryptor.DefaultGoRoutines
					}

					graphic.DoDecryption(path, numCpu, goroutines)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
