#!/usr/bin/env bash

rm $(pwd)/.git/hooks/pre-commit
ln -s $(pwd)/.hooks/govet-check $(pwd)/.git/hooks/pre-commit
chmod u+x $(pwd)/.hooks/govet-check
