package configuration

import "io/ioutil"
import "encoding/json"
import "os"

//import "log"

type AppData map[string]string

var configuration = make(map[string]AppData)

func GetFromJson(fileName string) AppData {
	if _, ok := configuration[fileName]; ok {
		return configuration[fileName]
	}

	var config = make(AppData)

	//log.Println("First load from " + fileName)

	// Default values
	config["listen"] = ":" + GetEnvOr("PORT", "8080")
	config["version"] = GetEnvOr("VERSION", "v1.0")
	config["root"] = GetEnvOr("ROOT", "public")
	config["javascriptDir"] = GetEnvOr("JAVASCRIPTDIR", "javascript")
	config["stylesheetDir"] = GetEnvOr("STYLESHEETDIR", "stylesheet")
	config["imageDir"] = GetEnvOr("IMAGEDIR", "image")

	// Load from the given json file
	jsonContent, error := ioutil.ReadFile(fileName)
	if error == nil {
		error = json.Unmarshal(jsonContent, &config)
		if error != nil {
			// log.Fatal(error)
		}
	}

	// Save to memory cache
	configuration[fileName] = config

	return config
}

// Get value by key from ENV, if not defined return defaultValue
func GetEnvOr(key string, defaultValue string) string {
	value := os.Getenv(key)

	if value == "" {
		value = defaultValue
	}

	return value
}
