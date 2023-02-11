TARNAME=gh_2.23.0_linux_amd64.tar.gz
FOLDERNAME=gh_2.23.0_linux_amd64

clean:
	rm $(TARNAME) && rm -rf $(FOLDERNAME)

test-install: test
	./test install-self

test-docker: 
	docker run -it --rm cbot918/ubugo

test: test.go
	go build .



	

# run:
# 	go run .

.PHONY: clean test test-install
.SILENT: clean test test-install