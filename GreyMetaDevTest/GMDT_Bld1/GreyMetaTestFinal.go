/*
     Author: Keith Armstrong
Description: Package Installations, development test example.
    Created: 06/16/2018 - KA
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

type userCommand struct {
	commandType   string
	commandParams []string
}

type progCommand struct {
	commandPrefix string
	commandSuffix string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fullCommand := scanner.Text()
		results := strings.Split(scanner.Text(), " ")
		uiCommand := userCommand{}
		uiCommandCount := len(results)
		uiParams := make([]string, uiCommandCount)
		for i := range results {
			if i == 0 {
				uiCommand.commandType = results[i]
			} else {
				uiParams[i] = results[i]
			}
		}
		currentCommandType := uiCommand.commandType

		if uiCommandCount > 0 {
			uiCommand.commandParams = uiParams
		}

		switch currentCommandType {

		case "LIST":
			fmt.Println(fullCommand)
			ListManualInstalledPkgsExport()
			ListAutoInstalledPkgsExport()
			break
		case "DEPEND":
			fmt.Println(fullCommand)
			for i := 1; i < len(uiParams); i++ {
				fullStr := "apt-cache show " + uiParams[i]
				ListDependanciesForPkgs(fullStr)
			}
			break
		case "INSTALL":
			fmt.Println(fullCommand)
			InstallPackageAndDependancies(uiParams[1])
			break
		case "REMOVE":
			fmt.Println(fullCommand)
			UnInstallPackageAndDependancies(uiParams[1])
			break
		case "END":
			fmt.Printf(fullCommand)
			runtime.Goexit()
			break

		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UnInstallPackageAndDependancies(pkgName string) {
	isPkgInstalled := IsPkgInstalled(pkgName)

	if isPkgInstalled == false {
		fmt.Println("Package not installed")
	} else {
		UnInstallPackage(pkgName)
	}
}

func InstallPackageAndDependancies(pkgName string) {
	isPkgInstalled := IsPkgInstalled(pkgName)
	if isPkgInstalled == true {
		fmt.Println("Package already installed")
	} else {
		installDeps := GetDependanciesForPkgs(pkgName)
		for i := range installDeps {
			installSubDeps := GetDependanciesForPkgs(installDeps[i])
			for j := range installSubDeps {
				if IsPkgInstalled(installSubDeps[j]) == false {
					InstallPackage(installSubDeps[j])
				}
			}
			if IsPkgInstalled(installDeps[i]) == false {
				InstallPackage(installDeps[i])
			}
		}
		InstallPackage(pkgName)

	}
}

func IsPkgInstalled(pkgName string) bool {
	fullStr := "apt-cache policy " + pkgName
	out, err := exec.Command("/bin/bash", "-c", fullStr).Output()
	check(err)
	return ExtractInstallationStatus(out)
}

func IsInstallPresent(pkgInfo []byte) bool {
	scanner := bufio.NewScanner(bytes.NewReader(pkgInfo))
	foundInstall := false
	for scanner.Scan() {
		lineString := scanner.Text()
		if len(scanner.Text()) >= 15 {
			subString := lineString[8:15]
			if subString == "install" {
				foundInstall = true
			}
		}
	}
	return foundInstall
}

func ExtractInstallationStatus(pkgInfo []byte) bool {
	scanner := bufio.NewScanner(bytes.NewReader(pkgInfo))
	for scanner.Scan() {
		lineString := scanner.Text()
		if len(scanner.Text()) >= 20 {
			fmt.Println(scanner.Text())
			subString := lineString[10:15]
			if subString == "(none)" {
				return false
			}
		}
	}
	return true
}

func InstallPackage(pkgName string) {
	fmt.Println("Installing..." + pkgName)
	d1 := []byte("sudo apt-get install " + pkgName)
	err := ioutil.WriteFile("./dat1.sh", d1, 0644)
	check(err)
	cmd := exec.Command("/bin/bash", "-c", "chmod +x dat1.sh")
	err1 := cmd.Run()
	check(err1)
	cmd1 := exec.Command("/bin/bash", "-c", "./dat1.sh")
	err2 := cmd1.Run()
	check(err2)
	fmt.Println("install done")
}

func UnInstallPackage(pkgName string) {
	fmt.Println("Un-installing..." + pkgName)
	d1 := []byte("sudo dpkg --purge --force-all " + pkgName)
	err := ioutil.WriteFile("./dat1.sh", d1, 0644)
	check(err)
	cmd := exec.Command("/bin/bash", "-c", "chmod +x dat1.sh")
	err1 := cmd.Run()
	check(err1)
	cmd1 := exec.Command("/bin/bash", "-c", "./dat1.sh")
	err2 := cmd1.Run()
	check(err2)
	fmt.Println("Un-install done, cleaning up...")
	d2 := []byte("apt-get autoremove")
	err3 := ioutil.WriteFile("./dat1.sh", d2, 0644)
	check(err3)
	fmt.Println("Clean-up done")
}

func GetPackageInfo(pkgName string) []byte {
	fullStr := "apt-cache show " + pkgName
	out, err := exec.Command("/bin/bash", "-c", fullStr).Output()
	if err != nil {
		log.Fatal(err)
	}
	return out
}

func ListDependanciesForPkgs(pkgName string) {

	pkgInfo := GetPackageInfo(pkgName)

	mainPkgDeps := ExtractDependancies(pkgInfo)

	fmt.Printf("DEPEND ")
	for i := range mainPkgDeps {
		fmt.Printf(mainPkgDeps[i] + " ")
	}

	fmt.Printf("\nDEPEND ")

	for i := range mainPkgDeps {
		fullStr := "apt-cache show " + mainPkgDeps[i]
		subPkgInfo := GetPackageInfo(fullStr)
		subPkgDeps := ExtractDependancies(subPkgInfo)
		for i := range subPkgDeps {
			fmt.Printf(subPkgDeps[i] + " ")
		}

		if i < len(subPkgDeps) {
			fmt.Printf("\nDEPEND ")
		}

	}
	fmt.Println("\n")
}

func GetDependanciesForPkgs(pkgName string) []string {
	pkgInfo := GetPackageInfo(pkgName)
	mainPkgDeps := ExtractDependancies(pkgInfo)
	allDeps := make([]string, 0)

	for i := range mainPkgDeps {
		allDeps = append(allDeps, mainPkgDeps[i])
	}

	for i := range mainPkgDeps {
		fullStr := "apt-cache show " + mainPkgDeps[i]
		subPkgInfo := GetPackageInfo(fullStr)
		subPkgDeps := ExtractDependancies(subPkgInfo)
		for i := range subPkgDeps {
			allDeps = append(allDeps, subPkgDeps[i])
		}
	}
	return allDeps
}

func ExtractDependancies(pkgBytes []byte) []string {
	scanner := bufio.NewScanner(bytes.NewReader(pkgBytes))
	subDeps := make([]string, 0)
	for scanner.Scan() {
		lineString := scanner.Text()
		if len(scanner.Text()) >= 9 {
			subString := lineString[0:8]
			if subString == "Depends:" {
				subDeps := GetSubDependancyNames(lineString)
				return subDeps
			}
		}
	}
	return subDeps
}

var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func GetSubDependancyNames(depStr string) []string {
	dependancyNames := make([]string, 0)
	wordSplit := strings.Split(depStr, " ")
	for i := range wordSplit {
		firstChar := wordSplit[i][0:1]
		if IsLetter(firstChar) && wordSplit[i] != "Depends:" {
			dependancyNames = append(dependancyNames, wordSplit[i])
		}
	}
	return dependancyNames
}

func ListManualInstalledPkgsExport() {

	out, err := exec.Command("/bin/bash", "-c", "apt-mark showmanual").Output()
	check(err)
	fmt.Printf("%s\n", out)
}

func ListAutoInstalledPkgsExport() {

	out, err := exec.Command("/bin/bash", "-c", "apt-mark showauto").Output()
	check(err)
	fmt.Printf("%s\n", out)
}
