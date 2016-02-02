#!/bin/bash

echo "Creating ./build and ./dist directories"
mkdir ./build 2> /dev/null
mkdir ./dist 2> /dev/null

TAG=`git describe --tags`

function make_release {
    EXEC=build/wundercli
    if [ $1 == "windows" ]
        then
            EXEC=$EXEC.exe
    fi

    echo ""
    echo "Making release for $1-$2"
    env GOOS=$1 GOARCH=$2 go build -o $EXEC
    echo "Creating tar archive"
    tar czf ./dist/wundercli-$TAG-$1-$2.tar.gz --directory=build . 2> /dev/null
    echo "Removing executable file"
    rm $EXEC
}

echo "Copying LICENSE and README.md to ./build"
cp LICENSE README.md build
make_release linux 386
make_release darwin 386
make_release windows 386

echo "Removing ./build temporary directory"
rm -rf build