#!/bin/bash

echo "Build C..."
gcc -O3 -o bignum-C bignum-C.c
echo "Build C++..."
g++ -O3 -o bignum-C++ bignum-C++.cpp
echo "Build Java..."
javac Bignumjava.java
echo "Build Go..."
go build bignum-golang.go
echo "Build Rust..."
rustc bignum-rust.rs

echo "C:"
time ./bignum-C

echo "\nC++:"
time ./bignum-C++

echo "\nGolang:"
time ./bignum-golang

echo "\nRust:"
time ./bignum-rust

echo "\nJava:"
time java Bignumjava

echo "\nNodejs:"
time node bignum-nodejs.js

echo "\nPHP:"
time php bignum-php.php

echo "\nPython:"
time python bignum-python.py
