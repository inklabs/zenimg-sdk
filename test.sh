#!/bin/bash

echo "PHP:"
php php/test.php

echo
echo "Ruby:"
ruby ruby/test.rb

echo
echo "Python:"
cd python && python test.py
cd ..

echo
echo "Javascript:"
cd js && node test.js
