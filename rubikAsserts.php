<?php
/* Rubik Asserts */
require_once("assertAux.php");
function assertOrderedRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $orderedValues, $orderedColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $orderedValues[$i] || 
      $rubik[$pos]['c'] !== $orderedColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertOrderedRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertOrderedRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveURubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveUValues, $moveUColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveUValues[$i] || 
      $rubik[$pos]['c'] !== $moveUColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveURubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveURubik\n".$red."Assert False\n".$white;
  }
}