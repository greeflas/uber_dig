.PHONY: run
run:
	go run .

.PHONY: echo
echo:
	curl -X POST -d 'hello' http://localhost:8080/echo

.PHONY: hello
hello:
	curl -X POST -d 'Gopher' http://localhost:8080/hello
