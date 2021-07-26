#!/usr/bin/env bashTIME=$(date)
# set exit status
EXIT_STATUS=$?

# Save the pwd before we run anything
PRE_PWD=`pwd`
# Determine the build script's actual directory
SOURCE="${BASH_SOURCE[0]}"
mkdir -p build/bin
BUILD_DIR="$(cd -P "$(dirname "$SOURCE")" && pwd)"

cd build/
printf "Building empty-tt..."
if go build -o bin/empty-tt ../app/empty-tt/main.go; then 
    printf "\rempty-tt: Build Succeeded\n"
else
    printf "\rempty-tt: Build Failed\n"
fi

if [ $EXIT_STATUS == 0 ]; then
  printf "Build succeeded\n"
else
  printf "Issues encountered. Build failed\n"
fi

mkdir -p resources/font
printf "Copying resources..."
# Copy resources
if cp -rf ../resources/font/* ./resources/font/; then
printf "\rCopying resources: done\n"
fi

exit $EXIT_STATUS