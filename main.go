package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/rakyll/statik/fs"
	_ "github.com/shouth/studio-aquatan-golang-hands-on/statik"
)

var flagVersion bool
var flagVerbose bool

func main() {
	rootCmd := flag.NewFlagSet("Root", flag.ContinueOnError)
	rootCmd.BoolVar(&flagVersion, "version", false, "print version")
	rootCmd.BoolVar(&flagVersion, "v", false, "print version")
	rootCmd.BoolVar(&flagVerbose, "verbose", false, "print log")

	err := rootCmd.Parse(os.Args[1:])
	if err != nil {
		if err == flag.ErrHelp {
			os.Exit(0)
		}
		log.Fatal(err)
	}
	if flagVersion {
		fmt.Println("golang hands-on 0.0.1")
	}

	addCmd := flag.NewFlagSet("add", flag.ContinueOnError)
	var fileName string
	addCmd.StringVar(&fileName, "name", time.Now().Format("2006-01-02")+".md", "filename")

	args := rootCmd.Args()
	if len(args) > 0 {
		switch args[0] {
		case "add":
			_ = addCmd.Parse(args[1:])
			handleAddCmd(fileName)
		}
	}
}

func handleAddCmd(filename string) error {
	statikFs, _ := fs.New()
	tplFile, _ := statikFs.Open("/report.md.tmpl")
	btpl, _ := ioutil.ReadAll(tplFile)
	stpl := string(btpl)

	tpl := template.Must(template.New("report").Parse(stpl))

	rptFile, _ := os.Create(filename)

	rptData := struct {
		Today string
	}{
		Today: time.Now().Format("2006-01-02"),
	}

	_ = tpl.Execute(rptFile, rptData)
	return nil
}
