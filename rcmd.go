package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

var version = "v0.0.1"

func showHelp() {
	const v = ` 
Usage: rcmd [options] [FILE]

Options:
  -h, --help            show this help message and exit
  --version             print the version and exit
  --config              path to the config
`
	os.Stderr.Write([]byte(v))
}

type cmdOptions struct {
	OptHelp    bool   `short:"h" long:"help" description:"show this help message"`
	OptVersion bool   `long:"version" description:"print the version"`
	OptConf    string `long:"config" description:"path to the config"`
}

func runServer(conf *Conf) error {
	server := NewServer(conf)
	err := server.Run()
	return err
}

func main() {
	opts := &cmdOptions{}
	parser := flags.NewParser(opts, flags.PrintErrors)
	args, err := parser.Parse()
	if err != nil {
		showHelp()
		return
	}

	if len(args) > 0 {
		showHelp()
		return
	}

	if opts.OptHelp {
		showHelp()
		return
	}

	if opts.OptVersion {
		fmt.Fprintf(os.Stderr, "rcmd: %s\n", version)
		return
	}

	if opts.OptConf == "" {
		fmt.Fprintln(os.Stderr, "config option is required.\n")
		showHelp()
		return
	}

	conf, err := LoadConfig(opts.OptConf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config: %s\n", err.Error())
		return
	}

	err = runServer(conf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to run server: %s\n", err.Error())
		return
	}

}
