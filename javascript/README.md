# Javascript Rubik

# Run in Terminal: It just shows ordered 3x3 colored rubik cube (magenta instead orange)

nodejs rubik.js

# Run in web

php -S 127.0.0.1:4500

http://127.0.0.1:4500/rubik

# Run in docker

docker run --name rubikjs -it --rm -v ${PWD}:/workdir node /bin/bash

# Moves

R = Right
r = Right prime
U = Up
u = Up prime
L = Left
l = Left prime
F = Front
f = Front prime
B = Back
b = Back prime
D = Down
d = Down prime

# Param for moves

Use query string parameter m as follow

http://127.0.0.1:4500/rubik.php?m=U

This show rubik ordered plus U move

http://127.0.0.1:4500/rubik.php?m=RUru

This shows the rubik ordered after sexy move

# Testing

php test.php

This tries every single move starting from ordered cube
Finally the test tries sexyMove