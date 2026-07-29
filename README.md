# Go-monitor

## CLI Server Monitoring Program

### Usage

*Go-monitor* gets server data from [servers.json](https://gitlab.com/amnes0a/go-monitor/-/blob/main/example-configuration.json?ref_type=heads). Alternatively, you can specify the path to the file using the --file flag

```bash
./bin/go-monitor --file $USER/nameFile.json
```

If *Go-monitor* is run in a directory that contains servers.json, it will immediately pick up the file

### Build / Docker

See [Makefile](https://gitlab.com/amnes0a/go-monitor/-/blob/main/Makefile?ref_type=heads) & [Dockerfile](https://gitlab.com/amnes0a/go-monitor/-/blob/main/Dockerfile?ref_type=heads)

#### RU

[README-RU](https://gitlab.com/amnes0a/go-monitor/-/blob/main/README-RU.md?ref_type=heads)
