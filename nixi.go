package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func readConfigFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	checkError(err)
	return string(data)
}

func writeConfigFile(filePath, data string) {
	err := os.WriteFile(filePath, []byte(data), 0664)
	checkError(err)
}

func updatePackagesList(dataAsStr, userParam string, install bool) string {
	regex := `environment\.systemPackages\s*=\s*with\s*pkgs;\s*\[([\s\S]*?)\];`
	compiled := regexp.MustCompile(regex)
	match := compiled.FindStringSubmatch(dataAsStr)
	if len(match) < 2 {
		log.Fatal("Your NixOS config file is not valid")
	}

	existingPkgs := match[1]
	var updatedPkgs string

	if install {
		newPkgs := strings.ReplaceAll(userParam, ",", " ")
		updatedPkgs = existingPkgs + "\n" + newPkgs
	} else {
		slice := strings.Split(userParam, ",")
		updatedPkgs = existingPkgs
		for _, pkg := range slice {
			updatedPkgs = strings.ReplaceAll(updatedPkgs, pkg, "")
		}
	}

	return strings.Replace(dataAsStr, existingPkgs, updatedPkgs, 1)
}

func rebuildNixOS() {
	fmt.Println("Building NixOS...")
	cmd := exec.Command("nixos-rebuild", "switch")
	out, err := cmd.CombinedOutput()
	checkError(err)
	fmt.Println(string(out))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s install/remove package1,package2,...", os.Args[0])
	}

	action := strings.ToLower(strings.Trim(os.Args[1], " "))
	userParam := strings.ToLower(strings.Trim(os.Args[2], " "))

	nixosConfigFile := "/etc/nixos/configuration.nix"

	var install bool
	switch action {
	case "install":
		install = true
	case "remove":
		install = false
	default:
		log.Fatalf("Unknown action: %s. Use install or remove.", action)
	}

	dataAsStr := readConfigFile(nixosConfigFile)
	dataAsStr = updatePackagesList(dataAsStr, userParam, install)
	writeConfigFile(nixosConfigFile, dataAsStr)

	actionStr := "installed"
	if !install {
		actionStr = "removed"
	}

	fmt.Printf("Successfully %s: %s\n", actionStr, userParam)

	rebuildNixOS()
}
