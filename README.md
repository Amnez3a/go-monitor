# Go-monitor

## CLI Программа для мониторинга серверов

### Использование

*Go-monitor* берет данные о серверах с [servers.json](https://github.com/Amnez3a/go-monitor/blob/main/example-configuration.json). Или же можно указать путь к файлу через флаг --file
```bash
./bin/go-monitor --file $USER/nameFile.json 
```
Если Go-monitor будет запущен в директории в которой есть servers.json, он сразу подхватит файл

### Сборка / Docker

см. [Makefile](https://github.com/Amnez3a/go-monitor/blob/main/Makefile) & [Dockerfile](https://github.com/Amnez3a/go-monitor/blob/main/Dockerfile)

### Зависимости

1. go

больше ничего и не надо.
