build:
	@go build -o bin/rssagg

run: build
	@./bin/rssagg