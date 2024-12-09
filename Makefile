build:
	@go build -o yung-gpt-cli main.go

run: build
	@./yung-gpt-cli

docker.build:
	@docker build -t yung-gpt-cli:latest .

docker.run:
	@docker run -it -e RAPIDAPI_KEY={RAPIDAPI_KEY} yung-gpt-cli:latest