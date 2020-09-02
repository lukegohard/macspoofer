build:
	go get -u golang.org/x/sys/unix
	go build macspoofer.go

move:
	mv macspoofer /usr/bin

uninstall:
	rm /usr/bin/macspoofer
	rm -r golang.org/x/sys/unix
