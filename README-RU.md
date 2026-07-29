# Go-monitor

## CLI Программа для мониторинга серверов

### Использование

*Go-monitor* берет данные о серверах с [servers.json](https://gitlab.com/amnes0a/go-monitor/-/blob/main/example-configuration.json?ref_type=heads). Или же можно указать путь к файлу через флаг --file

```bash
./bin/go-monitor --file $USER/nameFile.json 
```

Если Go-monitor будет запущен в директории в которой есть servers.json, он сразу подхватит файл

### Сборка / Docker

см. [Makefile](https://gitlab.com/amnes0a/go-monitor/-/blob/main/Makefile?ref_type=heads) & [Dockerfile](https://gitlab.com/amnes0a/go-monitor/-/blob/main/Makefile?ref_type=heads)

#### EN

[README-EN](https://gitlab.com/amnes0a/go-monitor/-/blob/main/README.md?ref_type=heads)
