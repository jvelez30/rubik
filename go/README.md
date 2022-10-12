# Go Rubik

# Compile go file in docker

docker run -it --rm -v "$PWD"/go/:/go/ -w ./go/ -w /go/ golang:1.19
unset GOPATH
go build rubik.go
./rubik

# Run in Terminal 

It just shows ordered 3x3 colored rubik cube (magenta instead orange)

go run rubik.go

# Run in web (this is not set yet)

go -S 127.0.0.1:4500

http://127.0.0.1:4500/rubik.go

# Run in docker (this is not working yet)

docker run --name rubikgo -it --rm -v ${PWD}:/workdir go /bin/bash
cd /working/go
go rubik.go

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

# Param for moves (this is not set yet)

Use query string parameter m as follow

http://127.0.0.1:4500/rubik.go?m=U

This show rubik ordered plus U move

http://127.0.0.1:4500/rubik.go?m=RUru

This shows the rubik ordered after sexy move

# Testing (this is not set yet)

This test tries every single move starting from ordered cube
Finally the test tries sexyMove

```bash
go test.go
```

Note: Actual test were made manually
TODO: create gounit standard test

