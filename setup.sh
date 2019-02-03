#!/usr/bin/env bash

ln -s $(pwd)/.hooks/pre-commit $(pwd)/.git/hooks/pre-commit
chmod u+x $(pwd)/.hooks/pre-commit
