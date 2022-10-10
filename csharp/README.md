# CSharp Rubik

# Run in Terminal 

It just shows ordered 3x3 colored rubik cube (magenta instead orange)

php rubik.php

# Run in web

docker run --name rubikcsharp -it --rm -v ${PWD}:/workdir -p 4800:80 mcr.microsoft.com/dotnet/samples:aspnetapp

http://127.0.0.1:4800/

# Run in docker

docker run --name rubikcsharp -it --rm -v ${PWD}:/workdir mcr.microsoft.com/dotnet/samples
cd /working/csharp
php rubik.php

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

This test tries every single move starting from ordered cube
Finally the test tries sexyMove

```bash
php test.php
```

Note: Actual test were made manually
TODO: create phpunit standard test

