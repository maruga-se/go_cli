package main

import "flag"

func parseArgs() (string, string, []string) {
	logDir := flag.String("logdir", "", "Log output directory(default)")
	flag.Parse()
	return *logDir, flag.Arg(0), flag.Args()[1:]
}
