# Scene Mover - Backend Implementation

This repository holds the backend implementation for a locally running web application I wrote to help me write a screenplay. [I published a blog article about the application](https://rickschubert.net/blog/posts/how-i-wrote-an-application-to-help-me-write-a-screenplay/), if you are keen to find out more.

Using go routines, the program spawns at launch a file watcher which looks for changes to a local folder, and a web server which is queried by the [Frontend Implementation](https://github.com/rickschubert/scenemover-visual-studio-code-extension) to rearrange scenes.

# Development
- Run `CompileDaemon -command="./scenemover" -exclude-dir ".git" -exclude-dir "scenes" -verbose` to automatically always restart the script should something change (this uses https://github.com/githubnemo/CompileDaemon)

# Build
- Compile for windows:

```sh
GOOS=windows GOARCH=amd64 go build -o scenemover.exe main.go
```

- Compile for Mac:
```sh
GOOS=darwin GOARCH=amd64 go build -o scenemover main.go
```
