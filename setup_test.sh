#!/bin/sh

echo -n "Creating testdata...\t\t\t"
mkdir testdata
echo "done"

echo -n "Download rules...\t\t\t"
git clone https://github.com/Yara-Rules/rules.git testdata/rules
echo "done"

echo -n "Creating uploads...\t\t\t"
mkdir uploads
echo "done"
