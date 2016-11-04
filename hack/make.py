#!/usr/bin/env python

import sys
import subprocess
from ntpath import *

GOC = 'go'  # if ENV in ['dev'] else 'godep go'

def die(status):
    if status:
        sys.exit(status)

def call(cmd, stdin=None):
    print(cmd)
    return subprocess.call([expandvars(cmd)], shell=True, stdin=stdin)

def gen_extpoints():
    die(call('go generate main.go'))


def fmt():
    die(call('goimports -w pkg'))
    call('go fmt ./pkg/...')

def default():
    gen_extpoints()
    fmt()
    die(call('GO15VENDOREXPERIMENT=1 ' + GOC + ' install .'))

if __name__ == "__main__":
    default()
