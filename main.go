package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var hexString string
	var filename string

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "hex",
			Usage:       "hex string",
			Destination: &hexString,
		},
		cli.StringFlag{
			Name:        "filename",
			Usage:       "output filename",
			Destination: &filename,
		},
	}

	app.Action = func(c *cli.Context) error {
		fmt.Printf("hexString:%s\n", hexString)
		fmt.Printf("filename:%s\n", filename)

		hexByte, err := hex.DecodeString(hexString)
		check(err)

		f, err := os.Create(filename)
		check(err)

		defer f.Close()

		n, err := f.Write(hexByte)
		check(err)
		fmt.Printf("wrote %d bytes\n", n)

		return nil
	}

	app.Run(os.Args)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
