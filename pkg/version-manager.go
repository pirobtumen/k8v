package pkg

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	kubectl_file    = "kubectl"
	kubectl_folder  = ".k8v"
	kubectl_mac_arm = "https://dl.k8s.io/release/v%s/bin/darwin/arm64/kubectl"
)

func DownloadKubectl(version string) {
	out, err := os.Create(kubectl_file)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	url := fmt.Sprintf(kubectl_mac_arm, version)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(kubectl_file, 0700)
	if err != nil {
		log.Fatal(err)
	}
}

func GetSrcDistFolder() (string, string) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	return pwd, userHomeDir + "/" + kubectl_folder
}

func SetKubectlVersion(version string) {
	src, dest := GetSrcDistFolder()
	oldLocation := fmt.Sprintf("%s/%s", src, kubectl_file)
	newLocation := fmt.Sprintf("%s/%s", dest, kubectl_file)

	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}
