#!/bin/sh

if [ "$GOARCH" = "386" ]; then LD=8; fi;
if [ "$GOARCH" = "amd64" ]; then LD=6; fi;
if [ "$GOARCH" = "arm" ]; then LD=5; fi;
LD=${LD}l

$LD $1 $2 -L `dirname $0`/_obj $3
