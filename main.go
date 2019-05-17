package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

type amxSys struct {
	Host     string `json:"Host"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type config struct {
	ModDir string
	EXE    string
	Mask   string   `json:"Mask"`
	AmxSys []amxSys `json:"Systems"`
}

func main() {

	// Set ConfigFile Variable
	cf := config{}

	// Get Command Line Variables
	flag.StringVar(&cf.ModDir, "ModDir", `.`, "Folder containing modules to be compiled")
	flag.StringVar(&cf.EXE, "Compiler", `C:\Program Files (x86)\Crestron\Simpl\SPlusCC.exe`, "Location of CrossCompiler .exe")
	flag.Parse()

	// Grab the folder
	log.Println(`Scanning ` + cf.ModDir)
	files, err := ioutil.ReadDir(cf.ModDir)
	if err != nil {
		log.Fatal(err)
	}
	// Loop the files
	for _, file := range files {
		fmt.Println(file.Name())
		abs, err := filepath.Abs(file.Name())
		if err != nil {
		}
		if filepath.Ext(abs) == `.usp` {
			log.Println(`Compiling ` + abs)
			// Prepare a compiler instance
			c := exec.Command(cf.EXE, `/rebuild "`+abs+`"`, "/target series2 series3")
			// Run the compiler
			c.Run()

			log.Println(`Compiled ` + abs)
		}
	}

}
