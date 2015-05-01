package main

import "io/ioutil"
import "encoding/json"
import "github.com/gin-gonic/gin"
import "log"
import "os"

type appData map[string]string

func main() {
  server := gin.Default()

  server.Static("/javascript", "./public/javascript")
  server.Static("/stylesheet", "./public/stylesheet")
  server.Static("/image", "./public/image")

  server.GET("/version", version)
  server.GET("/html/:fileName", servePublicFile)

  server.GET("/", serveRoot)

  appDetails := getConfig()

  server.Run(appDetails["listen"])
}

func serveRoot(c *gin.Context) {
  var fileCache = make(map[string]string)

  if fileCache["index.html"] == "" {
    c, _ := fileLoader("index.html")
    fileCache["index.html"] = c
  }
  content, _ := fileCache["index.html"]
  c.HTMLString(200, content)
}

func servePublicFile(c *gin.Context) {
  var fileCache = make(map[string]string)

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

func version(c *gin.Context) {
  appDetails := getConfig()
  c.JSON(200, gin.H{"version": appDetails["version"]})
}

func fileLoader(filePath string) (string, error) {
  content, error := ioutil.ReadFile("public/" + filePath)
  if error != nil {
    return "", error
  }
  return string(content), nil
}

func getConfig() appData {
  var appDetails appData

  appDetailsJson, error := ioutil.ReadFile("app.json")
  if error != nil {
    log.Fatal(error)
  }

  error = json.Unmarshal(appDetailsJson, &appDetails)
  if error != nil {
    log.Fatal(error)
  }

  if _, ok := appDetails["listen"]; !ok {
    appDetails["listen"] = os.Getenv("PORT")
  }

  if _, ok := appDetails["version"]; !ok {
    appDetails["version"] = os.Getenv("VERSION")
  }

  return appDetails
}
