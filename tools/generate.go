package main

import (
	"os"
	"os/exec"
	"text/template"
)

func main() {
	fmap := template.FuncMap{
		"calcTime": calcTime,
	}
	t := template.Must(template.New("README.md.tmpl").Funcs(fmap).ParseFiles("README.md.tmpl"))
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}

func calcTime(command string) string {
	cmd := exec.Command("/bin/bash", "tools/calcTime.sh", command)

	stdout, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	return string(stdout)
}
