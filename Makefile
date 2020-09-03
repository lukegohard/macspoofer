build:
	go get -u golang.org/x/sys/unix
	go build macspoofer.go

move:
	cp macspoofer /usr/bin

uninstall:
	rm /usr/bin/macspoofer
