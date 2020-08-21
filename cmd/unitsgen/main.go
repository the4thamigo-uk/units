package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"io/ioutil"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run() error {
	var cfgFile string
	pflag.StringVarP(&cfgFile, "cfg", "c", "", `configuration file`)

	var outFile string
	pflag.StringVarP(&outFile, "out", "o", "", `filename where code will be written`)

	pflag.Parse()

	cfg, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		panic(cfgFile)
		return err
	}
	ast, err := parse(string(cfg))
	if err != nil {
		return err
	}
	s, err := buildModel(ast)
	if err != nil {
		return err
	}
	code, err := generate(s)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(outFile, []byte(code), 0644)
}
