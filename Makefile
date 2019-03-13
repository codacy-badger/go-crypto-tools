.PHONY: test/web_explorer
test/web_explorer:
	go test -json ./web_explorer/...

.PHONY: test
test: test/web_explorer