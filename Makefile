LOCAL_BIN_PATH=/usr/local/bin

install: bin
	sudo cp bin/gob $(LOCAL_BIN_PATH)

bin:
	mkdir -p bin
	go build -o bin/gob cmd/main.go

run:
	go run cmd/main.go init .

clean:
	rm *.test

test-install:
	go run test/test_install/main.go



.PHONY: clean run test
.SILENT: clean run test