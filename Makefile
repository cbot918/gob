clean:
	rm *.test

test-install:
	go run test/test_install/main.go

run:
	go run .

.PHONY: clean run
.SILENT: clean run