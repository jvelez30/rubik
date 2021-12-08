# Php Rubik
# Run in Terminal: It just show ordered 3x3 colored rubik cube (magenta instead orange)

php rubik.php

# Run in web

php -S 127.0.0.1:4500

http://127.0.0.1:4500/rubik.php

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

# Moves get param

Use query string parameter m as follow

http://127.0.0.1:4500/rubik.php?m=U

This show rubik ordered plus U move

http://127.0.0.1:4500/rubik.php?m=RUrp

This shows the rubik ordered after sexy move

# Testing

php test.php

This tries every single move starting from ordered cube
Finally the test tries sexyMove