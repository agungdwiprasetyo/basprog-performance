#!/bin/bash

gcc -O3 -o bignum-C bignum-C.c
g++ -O3 -o bignum-C++ bignum-C++.cpp
javac Bignumjava.java
go build bignum-golang.go

echo "C:"
time ./bignum-C

echo "\nC++:"
time ./bignum-C++

echo "\nGolang:"
time ./bignum-golang

echo "\nJava:"
time java Bignumjava

echo "\nNodejs:"
time node bignum-nodejs.js

echo "\nPHP:"
time php bignum-php.php

echo "\nPython:"
time python bignum-python.py