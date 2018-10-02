VERSION = $(shell cat VERSION)

ROOTDIR = $(shell pwd)
APPNAME = gokit
APPPATH = github.com/mrlyc/${APPNAME}

GOPATH ?= /tmp/gopath
GODIR = $(shell echo "$${GOPATH%%:*}")
SRCDIR = ${GODIR}/src/${APPPATH}

GOENV = GOPATH=${GOPATH} GO15VENDOREXPERIMENT=1
GOROOTDIR = $(shell cd "${SRCDIR}" && go list -f '{{.Root}}')
PACKAGES = $(shell cd "${SRCDIR}" && go list ./...)
SUBDIRS = $(shell cd "${SRCDIR}" && go list -f '{{.Dir}}' ./...)
GOFILES = $(shell cd "${SRCDIR}" && go list -f '{{$$dir := .Dir}}{{range .GoFiles}}{{$$dir}}/{{.}}{{end}}' ./...)
GOTESTCASES = $(shell cd "${SRCDIR}" && go list -f '{{$$dir := .Dir}}{{range .TestGoFiles}}{{$$dir}}/{{.}}{{end}}{{range .XTestGoFiles}}{{$$dir}}/{{.}}{{end}}' ./...)
TOPFILES = $(wildcard *.*)

GO = ${GOENV} go
DEP = ${GOENV} dep

LDFLAGS = -X ${APPPATH}/Version=${VERSION}
DEBUGLDFLAGS = ${LDFLAGS} -X ${APPPATH}/Mode=debug
RELEASELDFLAGS = -w ${LDFLAGS} -X ${APPPATH}/Mode=release

.PHONY: .EXPORT_ALL_VARIABLES

build: init .EXPORT_ALL_VARIABLES
	${GO} install ${APPPATH}

${SRCDIR}:
	@mkdir -p `dirname "${SRCDIR}"`
	@ln -s ${ROOTDIR} ${SRCDIR}

.PHONY: init
init: ${SRCDIR}

.PHONY: update
update: ${SRCDIR} .EXPORT_ALL_VARIABLES
	cd "${SRCDIR}" && ${DEP} ensure -v

.PHONY: gofmt
gofmt: init .EXPORT_ALL_VARIABLES
	@gofmt -w ${GOFILES}

.PHONY: lint
lint: init .EXPORT_ALL_VARIABLES
	@golint ${PACKAGES}

.PHONY: test
test: init .EXPORT_ALL_VARIABLES
	@go test ${PACKAGES}

.PHONY: go-env
go-env:
	@${GOENV} go env
