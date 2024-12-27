# Python Rubik

## Run in Terminal

It just shows ordered 3x3 colored rubik cube (magenta instead orange)

python3 rubik.py

## Run in docker

docker run --name rubikpy -it --rm -v ${PWD}:/workdir -w /workdir python /bin/bash
python3 rubik.py

## Run in web (this is not working well yet)

run in docker

pip3 install flask
flask --app rubik run

<http://127.0.0.1:4700>

TODO: use some web framework like flask in order this run on the web

## Moves

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

## Param for moves (this is not working yet)

Use query string parameter m as follow

<http://127.0.0.1:4700/rubik.py?m=U>

This show rubik ordered plus U move

<http://127.0.0.1:4700/rubik.py?m=RUru>

This shows the rubik ordered after sexy move

## Testing

This test tries every single move starting from ordered cube
Finally the test tries sexyMove

python3 test.py

Note: no standard python test library is used
