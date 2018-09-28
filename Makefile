PKG := github.com/adelowo/mapped/cmd/mapped
PKG_PACKAGES := github.com/adelowo/mapped
# VERSION := $(shell git describe --abbrev=0 --tags)
PKG_LIST := $(shell go list ${PKG_PACKAGES}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)
LAST_SHA_ONE := $(shell git rev-parse --short HEAD)
CURRENT_GIT_BRANCH := $(shell git symbolic-ref --short HEAD)
BUILD_DATE=`date -u '+%Y-%m-%d_%H:%M:%S%p'`
FMT_BIN = goimports

verify_goimports:

	hash goimports > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u golang.org/x/tools/cmd/goimports; \
	fi

all: run

fmt:
	@for file in ${GO_FILES} ; do \
		$(FMT_BIN) -w $$file; \
	done

fmt-check:
	@diff=$$($(FMT_BIN) -d $(GO_FILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

dev:
	@go build -o mapped -ldflags="-X main.BuildDate=${BUILD_DATE} -X main.Version=${CURRENT_GIT_BRANCH}@${LAST_SHA_ONE}" ${PKG}

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	go build -i -v -o ${OUT}-v${VERSION} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X main.version=${VERSION}" ${PKG}

run: server
	./${OUT}

dependencies:
	@go get -tags integration -d -t ${PKG_LIST}

unit_tests:
	@go test ${PKG_LIST}

integration_tests:
	@go test ${PKG_LIST} -tags integration

.PHONY: run server static vet lint

