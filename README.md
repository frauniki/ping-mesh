# ping-mesh

## Build
### Agent
```shell
# Linux Agent
$ make linux-windows

# Windows Agent
$ make build-windows
```

### Docker Image
```shell
# Agent
$ make docker-build-agent

# Server
$ make docker-build-server
```

## How to use
### Agent
- Linux
```shell
$ ./bin/agent --config <ConfigFile path>
```

- Windows
```powershell
PS > ./bin/agent.exe --config <ConfigFile path>
```

### Server
```shell
$ pip install .
$ uvicorn ping_mesh:app --host 0.0.0.0 --port 8000
```
- Docs: http://127.0.0.1:8000/docs

## Configuration
### Agent
```yaml
#server_url: http://127.0.0.1:8000/ping-mesh # Ping-Mesh Server URL
#interval: 5s # Ticker interval time
#log_level: info #Log level
#ping:
  #count: 1 # Ping count
  #interval: 1s # Ping interval time
  #timeout: 2s # Ping timeout time

hosts:
  - name: router01 # Ping target hostname
    host: 10.0.0.1 # Ping target host (FQDN or IP Addres)
```
