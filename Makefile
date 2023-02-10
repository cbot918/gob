clean: gob.exe
	rm gob.exe
	rm -rf myproj

install-win: gob.exe config.json
	mkdir -p $(USERPROFILE)\myprogram\gob\bin
	cp gob.exe config.json $(USERPROFILE)\myprogram\gob\bin

test: gob.exe
	.\gob.exe myproj 

gob.exe: main.go
	go build .
