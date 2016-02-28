#!/bin/bash
for D in *; do
    if [ -d "${D}" ]; then
        echo "testing ${D} implementation"
        cp *_test.go "${D}"
        cp *_interface.go "${D}"
        if [ "template" != "${D}" ]; then
          cd "${D}"
          go test
          cd ..
        fi
    fi
done
