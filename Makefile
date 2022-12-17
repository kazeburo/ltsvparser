ifdef update
  u=-u
endif

VERSION=0.1.1
LDFLAGS=-ldflags "-X main.version=${VERSION}"
GO111MODULE=on


.PHONY: deps

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin main

check:
	go test .