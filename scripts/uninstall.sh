#!/bin/bash

# This script is assumed that invoked from makefile

BINARY_NAME=todo

: "go clean" && {
    go clean -i -x
}

: "remove built binary" && {
    set -x
    rm -f ${BINARY_NAME}
    set +x
}

: "remove datafile and command binary" && {
    set -x
    rm -f ${GOPATH}/bin/${BINARY_NAME}
    set +x

    if [ -z "${XDG_DATA_HOME}" ]; then
        set -x
        rm -rf ${HOME}/.local/home/share/todo
        set +x
    else
        set -x
        rm -rf ${XDG_DATA_HOME}/todo
        set +x
    fi
}
