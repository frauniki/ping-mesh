build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/agent_linux cmd/agent/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o ./bin/agent_windows.exe cmd/agent/main.go

docker-build-agent:
	docker build . -f cmd/agent/Dockerfile

docker-build-server:
	docker build . -f cmd/server/Dockerfile
