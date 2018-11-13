#!/bin/bash
set -x #echo on

BASEDIR=`pwd`
unamestr=`uname`

case $1 in
    sub.update)
        rm -rf $BASEDIR/svc
        git config -f .gitmodules --get-regexp '^submodule\..*\.path$' |
        while read path_key path
        do
            url_key=$(echo $path_key | sed 's/\.path/.url/')
            url=$(git config -f .gitmodules --get "$url_key")
            git submodule add --force $url $path
            cd $path && git pull
        done
        ;;
    esac
exit 0
