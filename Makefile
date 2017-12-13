#!/bin/make

SOURCES := $(wildcard *.go)
SOURCES += $(wildcard junit/*.go)

build: tesht

install: build
	install tesht $(DESTDIR)/usr/bin

uninstall:
	rm -f $(DESTDIR)/usr/bin/tesht

.PHONY: build install uninstall

tesht: $(SOURCES)
	go build

