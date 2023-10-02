#!/bin/sh

go build -a --ldflags "-w -s" -gcflags=all="-l -B" -o hero .
upx --best --lzma hero

#tinygo build -ldflags '-flto -Os -fdata-sections -ffunction-sections -Wl,--gc-sections -static -specs /usr/lib/x86_64-linux-musl/musl-gcc.specs' -o hero .