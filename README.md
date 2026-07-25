# Go-monitor

## CLI Server Monitoring Program

### Usage

*Go-monitor* gets server data from [servers.json](https://github.com/Amnez3a/go-monitor/blob/main/example-configuration.json). Alternatively, you can specify the path to the file using the --file flag

```bash
./bin/go-monitor --file $USER/nameFile.json
```

If *Go-monitor* is run in a directory that contains servers.json, it will immediately pick up the file

### Build / Docker

See [Makefile](https://github.com/Amnez3a/go-monitor/blob/main/Makefile) & [Dockerfile](https://github.com/Amnez3a/go-monitor/blob/main/Dockerfile)

#### RU

[README-RU](https://github.com/Amnez3a/go-monitor/blob/main/README-RU.md)
