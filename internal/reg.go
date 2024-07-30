package internal

import (
	"golang.org/x/sys/windows/registry"
	"log"
)

const (
	JavaSoftPath           = `SOFTWARE\JavaSoft\`
	JavaDevelopmentKitPath = JavaSoftPath + `\Java Development Kit`
	JDKPath                = JavaSoftPath + `\JDK`
)

// getRegistryAllSubKeysValues
//
// Retrieve the values of all sub keys in the registry.
func getRegistryAllSubKeysValues(path string, name string) []string {
	var paths []string
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, path, registry.QUERY_VALUE|registry.ENUMERATE_SUB_KEYS)
	if err != nil {
		log.Fatal(err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(key)

	subKeys, err := key.ReadSubKeyNames(-1)
	if err != nil {
		log.Fatal(err)
	}
	for _, subKeyName := range subKeys {
		openKey, err := registry.OpenKey(registry.LOCAL_MACHINE, path+`\`+subKeyName, registry.QUERY_VALUE)
		if err != nil {
			log.Fatal(err)
		}
		value, _, err := openKey.GetStringValue(name)
		paths = append(paths, value)
	}
	return paths
}
