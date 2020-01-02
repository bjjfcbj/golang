#================================================================
#   Copyright (C) 2020 D. All rights reserved.
#   
#   文件名称：makefile
#   创 建 者：D
#   创建日期：2020年01月02日
#   描    述：
#
#================================================================

MSE:=..
.PHONY: all push pull

all:

push:
	git add .&&git commit -m '$(MSE)'&&git push
pull:
	git pull
