package ServerForReact

import (
	"io/ioutil"
	"path/filepath"
)
import "github.com/gin-gonic/gin"
import "github.com/yitsushi/Server-For-React/configuration"

//import "log"

// Memory cache for server files
var fileCache = make(map[string]string)

// Configuration holder
var appConfiguration configuration.AppData

// Run server
func Run(configurationFilename string) {
	server := gin.Default()

	appConfiguration = configuration.GetFromJson(configurationFilename)

	server.Static("/javascript", filepath.Join(appConfiguration["root"], appConfiguration["javascriptDir"]))
	server.Static("/stylesheet", filepath.Join(appConfiguration["root"], appConfiguration["stylesheetDir"]))
	server.Static("/image", filepath.Join(appConfiguration["root"], appConfiguration["imageDir"]))

	server.GET("/version", ServeVersion)
	server.GET("/raw/:fileName", ServeFile)

	server.GET("/", ServeRoot)

	server.Run(appConfiguration["listen"])
}

// Server index.html as root
func ServeRoot(c *gin.Context) {
	content, _ := fileLoader("index.html")
	if content != "" {
		c.HTMLString(200, content)
	} else {
		c.HTMLString(404, "404 not found")
	}
}

// Server files
func ServeFile(c *gin.Context) {
	content, _ := fileLoader(c.Params.ByName("fileName"))
	if content != "" {
		c.HTMLString(200, content)
	} else {
		c.HTMLString(404, "404 not found")
	}
}

// Serve current version number as a simple endpoint
// God to test if server is running or not
func ServeVersion(c *gin.Context) {
	c.JSON(200, gin.H{"version": appConfiguration["version"]})
}

// Load file from ROOT
func fileLoader(filePath string) (string, error) {
	var content string
	if appConfiguration["cacheFiles"] == "false" || fileCache[filePath] == "" {
		c, error := ioutil.ReadFile(filepath.Join(appConfiguration["root"], filePath))
		if error != nil {
			return "", error
		}
		content = string(c)
		fileCache[filePath] = content
	} else {
		content, _ = fileCache[filePath]
	}
	return content, nil
}
