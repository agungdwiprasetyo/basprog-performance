#!/bin/bash
echo "Test C with opt:"
time ./bignum-C
echo "Test C++ with opt:"
time ./bignum-C++
echo "Test Go:"
time ./bignum-golang
echo "Test :"
