#================================================================
#   Copyright (C) 2020 D. All rights reserved.
#   
#   文件名称：makefile
#   创 建 者：D
#   创建日期：2020年01月02日
#   描    述：
#
#================================================================

g:=go
MAKE:=make all
SRC:=$(wildcard ./*.go)
WORK:=$(patsubst ./%.go, ./go_%,$(SRC))

all:$(WORK)

run:
	$(MAKE)&&$(g) run $(FLAG).go

go_%:%.go
	$(g) build -o $@ $<

clean:
	-rm *_

.PHONY: all run clean
