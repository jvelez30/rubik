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
echo $blue."MoveU\n".$white;
assertMoveURubik($rubik);
// Test moveUPrime function
$rubik=moveUPrime($rubik);
echo $blue."MoveUPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=moveUPrime($rubik);
echo $blue."MoveUPrime\n".$white;
assertMoveUPrimeRubik($rubik);
$rubik=moveU($rubik);
// Test moveR function
$rubik=moveR($rubik);
echo $blue."MoveR\n".$white;
assertMoveRRubik($rubik);
// Test moveRPrime function
$rubik=moveRPrime($rubik);
echo $blue."MoveRPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=moveRPrime($rubik);
echo $blue."MoveRPrime\n".$white;
assertMoveRPrimeRubik($rubik);
$rubik=moveR($rubik);
// Test moveL function
$rubik=moveL($rubik);
echo $blue."MoveL\n".$white;
assertMoveLRubik($rubik);
// Test moveLPrime function
$rubik=moveLPrime($rubik);
echo $blue."MoveLPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=moveLPrime($rubik);
echo $blue."MoveLPrime\n".$white;
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
echo $blue."MoveB\n".$white;
assertMoveBRubik($rubik);
// Test MoveBPrime function
$rubik=MoveBPrime($rubik);
echo $blue."MoveBPrime\n".$white;
assertOrderedRubik($rubik);
$rubik=MoveBPrime($rubik);
echo $blue."MoveBPrime\n".$white;
assertMoveBPrimeRubik($rubik);
$rubik=MoveB($rubik);
format($rubik);
