package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

// JDKInstalledInfo
//
// JDK installation information structure.
type JDKInstalledInfo struct {
	Version  string `json:"version"`
	JavaHome string `json:"javaHome"`
}

// String
//
// The toString Implementation of JDKInstallingDnfo Structure
func (j *JDKInstalledInfo) String() string {
	return fmt.Sprintf("{Version: %s, JavaHome: %s}", j.Version, j.JavaHome)
}

// getJavaHomePaths
//
// Retrieve the Path array of JavaHome installed in the system.
func getJavaHomePaths() []string {
	javaDevelopmentKitPaths := getRegistryAllSubKeysValues(JavaDevelopmentKitPath, "JavaHome")
	jdkPaths := getRegistryAllSubKeysValues(JDKPath, "JavaHome")
	paths := append(javaDevelopmentKitPaths, jdkPaths...)
	return slices.Compact(paths)
}

// readJdkRelease
//
// Read release file information through JavaHome directory.
func readJdkRelease(path string) map[string]string {
	releaseFilePath := path + `\release`
	file, err := os.Open(releaseFilePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	scanner := bufio.NewScanner(file)
	releaseData := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			cleanValue := strings.Trim(value, "\"")
			releaseData[key] = cleanValue
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	return releaseData
}

// InstalledJDKInfo
//
// installed JDK info
func InstalledJDKInfo() []*JDKInstalledInfo {
	var infoList []*JDKInstalledInfo
	paths := getJavaHomePaths()
	for _, path := range paths {
		releaseData := readJdkRelease(path)
		jdkVersion := releaseData["JAVA_VERSION"]
		infoList = append(infoList, &JDKInstalledInfo{jdkVersion, path})
	}
	return infoList
}

// CurrentUseJdkVersion
//
// The current version of JDK being used.
func CurrentUseJdkVersion() string {
	javaHome := os.Getenv("JAVA_HOME")
	releaseData := readJdkRelease(javaHome)
	return releaseData["JAVA_VERSION"]
}
