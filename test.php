<?php
/* Functional Test */
require_once("rubikAsserts.php");
require_once("rubik.php");
// Test getOrderedRubik function
$rubik=getOrderedRubik();
echo $blue."GetRubik\n".$white;
assertOrderedRubik($rubik);
// Test moveU function
$rubik=moveU($rubik);
echo $blue."moveU\n".$white;
assertMoveURubik($rubik);
// Test moveUP function
$rubik=moveUP($rubik);
echo $blue."moveUP\n".$white;
assertOrderedRubik($rubik);
$rubik=moveUP($rubik);
echo $blue."moveUP\n".$white;
assertmoveUPRubik($rubik);
$rubik=moveU($rubik);
// Test moveR function
$rubik=moveR($rubik);
echo $blue."moveR\n".$white;
assertMoveRRubik($rubik);
// Test moveRPrime function
$rubik=moveRPrime($rubik);
echo $blue."moveRPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=moveRPrime($rubik);
echo $blue."moveRPrime\n".$white;
assertMoveRPrimeRubik($rubik);
$rubik=moveR($rubik);
// Test moveL function
$rubik=moveL($rubik);
echo $blue."moveL\n".$white;
assertMoveLRubik($rubik);
// Test moveLPrime function
$rubik=moveLPrime($rubik);
echo $blue."moveLPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=moveLPrime($rubik);
echo $blue."moveLPrime\n".$white;
assertMoveLPrimeRubik($rubik);
$rubik=moveL($rubik);
// Test moveF function
$rubik=moveF($rubik);
echo $blue."moveF\n".$white;
assertmoveFRubik($rubik);
// Test moveFPrime function
$rubik=moveFPrime($rubik);
echo $blue."moveFPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=moveFPrime($rubik);
echo $blue."moveFPrime\n".$white;
assertmoveFPrimeRubik($rubik);
$rubik=MoveF($rubik);
// Test MoveB function
$rubik=MoveB($rubik);
echo $blue."moveB\n".$white;
assertMoveBRubik($rubik);
// Test MoveBPrime function
$rubik=MoveBPrime($rubik);
echo $blue."moveBPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=MoveBPrime($rubik);
echo $blue."moveBPrime\n".$white;
assertMoveBPrimeRubik($rubik);
$rubik=MoveB($rubik);
// Test MoveD function
$rubik=MoveD($rubik);
echo $blue."moveD\n".$white;
assertMoveDRubik($rubik);
// Test MoveDPrime function
$rubik=MoveDPrime($rubik);
echo $blue."moveDPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=MoveDPrime($rubik);
echo $blue."moveDPrime\n".$white;
assertMoveDPrimeRubik($rubik);
$rubik=MoveD($rubik);
format($rubik);
