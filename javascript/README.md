# Javascript Rubik

# Run in Terminal (this is not working yet)

It just shows ordered 3x3 colored rubik cube (magenta instead orange)

nodejs rubik.js
TODO: Review full copy of array because references are pointing to the same

# Run in web (this is not working yet)
TODO: use express framework to running on web

# Run in docker (this is not working yet)

docker run --name rubikjs -it --rm -v ${PWD}:/workdir -w /workdir node /bin/bash
nodejs rubik.js

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

# Param for moves (this is not working well yet)

Use query string parameter m as follow

http://127.0.0.1:4500/rubik.php?m=U

This show rubik ordered plus U move

http://127.0.0.1:4500/rubik.php?m=RUru

This shows the rubik ordered after sexy move

# Testing (this is not working well yet)

This test tries every single move starting from ordered cube
Finally the test tries sexyMove

php test.php

