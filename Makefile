.PHONY: all setup schema build install run clean
.DEFAULT_GOAL=all

GRAPHQL_SCHEMA_BINDATA=graphql/schema/bindata.go

UNIT_DIR=unit/
SYSTEMD_DIR=$(DESTDIR)/usr/lib/systemd/system/
USR_BIN_DIR=$(DESTDIR)/usr/bin/

CURATOR=curator
CURATOR_SERVICE=curator.service

all: build

setup:
	go get -u github.com/jteeuwen/go-bindata/...

schema: $(GRAPHQL_SCHEMA_BINDATA)

build: $(CURATOR)

run: schema
	go run main.go

install: build
	[ -d $(SYSTEMD_DIR) ] || install -d $(SYSTEMD_DIR)
		install --mode 644 $(UNIT_DIR)$(CURATOR_SERVICE) $(SYSTEMD_DIR)$(CURATOR_SERVICE)
	[ -d $(USR_BIN_DIR) ] || install -d $(USR_BIN_DIR)
	    install --mode 500 $(CURATOR) $(USR_BIN_DIR)$(CURATOR)

clean:
	rm -f graphql/schema/bindata.go
	go clean

$(GRAPHQL_SCHEMA_BINDATA):
	go-bindata -ignore=\.go -pkg=schema -o=graphql/schema/bindata.go graphql/schema/...

$(CURATOR): schema
	go build
