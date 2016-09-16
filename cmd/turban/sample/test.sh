#!/bin/sh

if [[ "$1" == "fail" ]]; then
    echo "Let's fail"
    exit 1
elif [[ "$1" == "foo" ]] && [[ "$2" == "bar" ]]; then
    echo "got foo bar!"
    exit 0
elif [[ "$1" == "foo" ]]; then
    echo "bar"
    exit 0
elif [[ "$1" == "read" ]]; then
    echo -n "please enter anything: "
    read input
    echo "got: $input"
    exit 0
else
    echo "invalid arguments specified"
    exit 1
fi
