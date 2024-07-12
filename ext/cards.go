package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"

	"github.com/schollz/progressbar/v3"
)

func main() {
	filepath := "/home/amar/amar.vsix"
	fmt.Println("Downloading")
	// downloadFile("amar.vsix", "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/esbenp/vsextensions/prettier-vscode/10.4.0/vspackage")
	download(filepath, "https://marketplace.visualstudio.com/_apis/public/gallery/publishers/esbenp/vsextensions/prettier-vscode/10.4.0/vspackage")
	installvsix(filepath)
}

func downloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func download(destinationPath, downloadUrl string) error {
	tempDestinationPath := destinationPath + ".tmp"
	req, _ := http.NewRequest("GET", downloadUrl, nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	f, _ := os.OpenFile(tempDestinationPath, os.O_CREATE|os.O_WRONLY, 0644)

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)
	io.Copy(io.MultiWriter(f, bar), resp.Body)
	os.Rename(tempDestinationPath, destinationPath)
	return nil
}

func installvsix(filepath string) (err error) {
	vsix_install := "code"
	installation_arg := "--install-extension"
	extension_name := filepath
	cmd := exec.Command(vsix_install, installation_arg, extension_name)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(stdout))
	return nil
}
