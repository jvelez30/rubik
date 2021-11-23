<?php
/* Functional Test */
require_once("rubikAsserts.php");
require_once("rubik.php");
// Test getOrderedRubik function
$rubik=getOrderedRubik();
$rubik=moveU($rubik);
//assertOrderedRubik($rubik);
assertMoveURubik($rubik);