#!/bin/make

SOURCES := $(wildcard *.go)
SOURCES += $(wildcard junit/*.go)

build: tesht

install: build
	install tesht $(DESTDIR)/usr/bin

uninstall:
	rm -f $(DESTDIR)/usr/bin/tesht

test: build
	cp junit/test/results.xml .
	./tesht true
	./tesht false
	./tesht find . -name results.xml

.PHONY: build install uninstall

tesht: $(SOURCES)
	go build

