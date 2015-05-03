package ServerForReact

import "io/ioutil"
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

	server.Static("/javascript", "./"+appConfiguration["root"]+"/"+appConfiguration["javascriptDir"])
	server.Static("/stylesheet", "./"+appConfiguration["root"]+"/"+appConfiguration["stylesheetDir"])
	server.Static("/image", "./"+appConfiguration["root"]+"/"+appConfiguration["imageDir"])

	server.GET("/version", ServeVersion)
	server.GET("/raw/:fileName", ServeFile)

	server.GET("/", ServeRoot)

	server.Run(appConfiguration["listen"])
}

// Server index.html as root
func ServeRoot(c *gin.Context) {
	if fileCache["index.html"] == "" {
		c, _ := fileLoader("index.html")
		fileCache["index.html"] = c
	}
	content, _ := fileCache["index.html"]
	if content != "" {
		c.HTMLString(200, content)
	} else {
		c.HTMLString(404, "404 not found")
	}
}

// Server files
func ServeFile(c *gin.Context) {
	fileName := c.Params.ByName("fileName")
	if fileCache[fileName] == "" {
		c, _ := fileLoader(fileName)
		fileCache[fileName] = c
	}
	content, _ := fileCache[c.Params.ByName("fileName")]
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
	//log.Println("Load new file: " + filePath)
	content, error := ioutil.ReadFile(appConfiguration["root"] + "/" + filePath)
	if error != nil {
		return "", error
	}
	return string(content), nil
}
