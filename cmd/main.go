package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pirobtumen/k8v/pkg"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Missing kubectl version")
	}

	version := os.Args[1]

	fmt.Print("Downloading kubectl v", version, "\n\n")
	pkg.DownloadKubectl(version)

	fmt.Print("Setting /usr/local/bin/kubectl as v", version, "\n\n")
	pkg.SetKubectlVersion(version)
}
