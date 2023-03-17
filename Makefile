.PHONY: run
run:
	go run .

.PHONY: echo
echo:
	curl -X POST -d 'hello' http://localhost:8080/echo
