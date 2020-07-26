#!/usr/bin/env bash
set -e
source .commands/colors.sh

echo -e "${YELLOW} Testing ${NC}"
cd tests/

if go test -count=1 -v ; then
    echo -e "${LBLUE}PASS${NC}"
    exit 0
fi


echo -e "${RED}FAIL${NC}"
exit 1
