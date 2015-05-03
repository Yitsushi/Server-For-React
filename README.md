# Server-For-React

### Usage

#### Method 1: own server (if you want to host on heroku for instance)

Create a file with name `server.go`

```go
package main

import SFR "github.com/yitsushi/server-for-react"

func main() {
  SFR.Run("app.json")
}
```

#### Method 2: standalone for development

```
$ go get github.com/yitsushi/server-for-react/standaloneStaticServer
```

### Structure

Create directory `public` at least with an `index.html`.

My directory structure:

```
public
  javascript
    any.js
  stylesheet
    any.css
  image
    any.png
    any.jpeg
    any.webp
  inde.html
src
  javascript
    reactCodeHere
  stylesheet
    lessCodeHere
app.json
```

app.json:

```json
{
  "version": "v1.0",
  "listen": ":8080",
  "root": "public",
  "javascriptDir": "javascript",
  "stylesheetDir": "stylesheet",
  "imageDir": "image"
}
```

without app.json use environment variables:
 - `ENV['VERSION']`
 - `ENV['PORT']`
 - `ENV['ROOT']`
 - `ENV['JAVASCRIPTDIR']`
 - `ENV['STYLESHEETDIR']`
 - `ENV['IMAGEDIR']`
