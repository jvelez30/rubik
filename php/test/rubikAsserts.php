<?php
/* Rubik Asserts */
require_once __DIR__ . '/lib/assertAux.php';
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
function assertMoveUPRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveUPValues, $moveUPColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveUPValues[$i] || 
      $rubik[$pos]['c'] !== $moveUPColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveUPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveUPRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveRRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveRValues, $moveRColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveRValues[$i] || 
      $rubik[$pos]['c'] !== $moveRColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveRRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveRRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveRPRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveRPValues, $moveRPColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveRPValues[$i] || 
      $rubik[$pos]['c'] !== $moveRPColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveRPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveRPRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveLRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveLValues, $moveLColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveLValues[$i] || 
      $rubik[$pos]['c'] !== $moveLColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveLRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveLRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveLPRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveLPValues, $moveLPColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveLPValues[$i] || 
      $rubik[$pos]['c'] !== $moveLPColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveLPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveLPRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveFRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveFValues, $moveFColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveFValues[$i] || 
      $rubik[$pos]['c'] !== $moveFColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveFRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveFRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveFPRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveFPValues, $moveFPColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveFPValues[$i] || 
      $rubik[$pos]['c'] !== $moveFPColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveFPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveFPRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveBRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveBValues, $moveBColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveBValues[$i] || 
      $rubik[$pos]['c'] !== $moveBColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveBRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveBRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveBPRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveBPValues, $moveBPColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveBPValues[$i] || 
      $rubik[$pos]['c'] !== $moveBPColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveBPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveBPRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveDRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveDValues, $moveDColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveDValues[$i] || 
      $rubik[$pos]['c'] !== $moveDColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveDRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveDRubik\n".$red."Assert False\n".$white;
  }
}
function assertMoveDPRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveDPValues, $moveDPColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveDPValues[$i] || 
      $rubik[$pos]['c'] !== $moveDPColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveDPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveDPRubik\n".$red."Assert False\n".$white;
  }
}
function assertSexyMoveRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $sexyMoveValues, $sexyMoveColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $sexyMoveValues[$i] || 
      $rubik[$pos]['c'] !== $sexyMoveColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertSexyMoveRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertSexyMoveRubik\n".$red."Assert False\n".$white;
  }
}