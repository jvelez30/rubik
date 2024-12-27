# Php Rubik

## Run in Terminal

It just shows ordered 3x3 colored rubik cube (magenta instead orange)

php rubik.php

## Run in web

php -S 127.0.0.1:4500

<http://127.0.0.1:4500/app/rubik.php>

## Run in docker

docker run --name rubikphp -it --rm -v ${PWD}:/workdir -w /workdir php /bin/bash
php rubik.php

## Run web server in docker

docker run --name rubikphp -it --rm -p 8095:4500 -v ${PWD}:/workdir -w /workdir --entrypoint php php -S 0.0.0.0:4500

## Run nginx server in docker

docker build -t rubik-nginx ./docker
docker run --name rubiknginx -it --rm -p 8096:8095 -v $(pwd)/app:/app -v $(pwd)/lib:/app/lib rubik-nginx

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

## Param for moves

Use query string parameter m as follow

<http://127.0.0.1:8096/rubik.php?m=U>

This show rubik ordered plus U move

<http://127.0.0.1:8096/rubik.php?m=RUru>

This shows the rubik ordered after sexy move

<http://127.0.0.1:8096/rubik.php?s=1&m=RUruRUruRUruRUruRUruRUru>

It display "Ordered" after make 6 sexy moves

## Testing

This test tries every single move starting from ordered cube
Finally the test tries sexyMove

```bash
cd test
php test.php
```

Note: Actual test was made manually
TODO: create phpunit standard test
