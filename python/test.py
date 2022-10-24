# Functional Test 

from rubikAsserts import *
from rubik import *
# Test getOrderedRubik function
rubik = getOrderedRubik()
print(str(str(blue) + "GetRubik\n") + str(white),end="")
assertOrderedRubik(rubik)
# Test moveU function
rubik = moveU(rubik)
print(str(str(blue) + "moveU\n") + str(white),end="")
assertMoveURubik(rubik)
# Test moveUP function
rubik = moveUP(rubik)
print(str(str(blue) + "moveUP\n") + str(white),end="")
assertOrderedRubik(rubik)
rubik = moveUP(rubik)
print(str(str(blue) + "moveUP\n") + str(white),end="")
assertMoveUPRubik(rubik)
rubik = moveU(rubik)
# Test moveR function
rubik = moveR(rubik)
print(str(str(blue) + "moveR\n") + str(white),end="")
assertMoveRRubik(rubik)
# Test moveRP function
rubik = moveRP(rubik)
print(str(str(blue) + "moveRP\n") + str(white),end="")
assertOrderedRubik(rubik)
rubik = moveRP(rubik)
print(str(str(blue) + "moveRP\n") + str(white),end="")
assertMoveRPRubik(rubik)
rubik = moveR(rubik)
# Test moveL function
rubik = moveL(rubik)
print(str(str(blue) + "moveL\n") + str(white),end="")
assertMoveLRubik(rubik)
# Test moveLP function
rubik = moveLP(rubik)
print(str(str(blue) + "moveLP\n") + str(white),end="")
assertOrderedRubik(rubik)
rubik = moveLP(rubik)
print(str(str(blue) + "moveLP\n") + str(white),end="")
assertMoveLPRubik(rubik)
rubik = moveL(rubik)
# Test moveF function
rubik = moveF(rubik)
print(str(str(blue) + "moveF\n") + str(white),end="")
assertMoveFRubik(rubik)
# Test moveFP function
rubik = moveFP(rubik)
print(str(str(blue) + "moveFP\n") + str(white),end="")
assertOrderedRubik(rubik)
rubik = moveFP(rubik)
print(str(str(blue) + "moveFP\n") + str(white),end="")
assertMoveFPRubik(rubik)
rubik = moveF(rubik)
# Test MoveB function
rubik = moveB(rubik)
print(str(str(blue) + "moveB\n") + str(white),end="")
assertMoveBRubik(rubik)
# Test MoveBP function
rubik = moveBP(rubik)
print(str(str(blue) + "moveBP\n") + str(white),end="")
assertOrderedRubik(rubik)
rubik = moveBP(rubik)
print(str(str(blue) + "moveBP\n") + str(white),end="")
assertMoveBPRubik(rubik)
rubik = moveB(rubik)
# Test MoveD function
rubik = moveD(rubik)
print(str(str(blue) + "moveD\n") + str(white),end="")
assertMoveDRubik(rubik)
# Test MoveDP function
rubik = moveDP(rubik)
print(str(str(blue) + "moveDP\n") + str(white),end="")
assertOrderedRubik(rubik)
rubik = moveDP(rubik)
print(str(str(blue) + "moveDP\n") + str(white),end="")
assertMoveDPRubik(rubik)
rubik = moveD(rubik)
# Test sexy move function
rubik = sexyMove(rubik)
print(str(str(blue) + "sexyMove\n") + str(white),end="")
assertSexyMoveRubik(rubik)
# 5 more sexy moves then cube is ordered
j = 0
while ( j < 5 ) :
    rubik = sexyMove(rubik)
    j+=1

print(str(str(blue) + "5 more sexyMoves\n") + str(white),end="")
assertOrderedRubik(rubik)
formatCli(rubik)