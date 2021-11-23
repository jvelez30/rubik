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
assertMoveUPRubik($rubik);
$rubik=moveU($rubik);
// Test moveR function
$rubik=moveR($rubik);
echo $blue."moveR\n".$white;
assertMoveRRubik($rubik);
// Test moveRP function
$rubik=moveRP($rubik);
echo $blue."moveRP\n".$white;
assertOrderedRubik($rubik);
$rubik=moveRP($rubik);
echo $blue."moveRP\n".$white;
assertMoveRPRubik($rubik);
$rubik=moveR($rubik);
// Test moveL function
$rubik=moveL($rubik);
echo $blue."moveL\n".$white;
assertMoveLRubik($rubik);
// Test moveLP function
$rubik=moveLP($rubik);
echo $blue."moveLP\n".$white;
assertOrderedRubik($rubik);
$rubik=moveLP($rubik);
echo $blue."moveLP\n".$white;
assertMoveLPRubik($rubik);
$rubik=moveL($rubik);
// Test moveF function
$rubik=moveF($rubik);
echo $blue."moveF\n".$white;
assertMoveFRubik($rubik);
// Test moveFP function
$rubik=moveFP($rubik);
echo $blue."moveFP\n".$white;
assertOrderedRubik($rubik);
$rubik=moveFP($rubik);
echo $blue."moveFP\n".$white;
assertMoveFPRubik($rubik);
$rubik=moveF($rubik);
// Test MoveB function
$rubik=moveB($rubik);
echo $blue."moveB\n".$white;
assertMoveBRubik($rubik);
// Test MoveBP function
$rubik=moveBP($rubik);
echo $blue."moveBP\n".$white;
assertOrderedRubik($rubik);
$rubik=moveBP($rubik);
echo $blue."moveBP\n".$white;
assertMoveBPRubik($rubik);
$rubik=moveB($rubik);
// Test MoveD function
$rubik=moveD($rubik);
echo $blue."moveD\n".$white;
assertMoveDRubik($rubik);
// Test MoveDP function
$rubik=moveDP($rubik);
echo $blue."moveDP\n".$white;
assertOrderedRubik($rubik);
$rubik=moveDP($rubik);
echo $blue."moveDP\n".$white;
assertMoveDPRubik($rubik);
$rubik=moveD($rubik);
// Test sexy move function
$rubik=sexyMove($rubik);
echo $blue."sexyMove\n".$white;
assertSexyMoveRubik($rubik);
// 5 more sexy moves then cube is ordered
for ($j=0;$j<5;$j++){
  $rubik=sexyMove($rubik);
}
echo $blue."5 more sexyMoves\n".$white;
assertOrderedRubik($rubik);
formatCli($rubik);
