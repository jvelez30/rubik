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
function assertmoveUPRubik($rubik){
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
    echo "assertmoveUPRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertmoveUPRubik\n".$red."Assert False\n".$white;
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
function assertMoveRPrimeRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveRPrimeValues, $moveRPrimeColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveRPrimeValues[$i] || 
      $rubik[$pos]['c'] !== $moveRPrimeColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveRPrimeRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveRPrimeRubik\n".$red."Assert False\n".$white;
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
function assertMoveLPrimeRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveLPrimeValues, $moveLPrimeColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveLPrimeValues[$i] || 
      $rubik[$pos]['c'] !== $moveLPrimeColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveLPrimeRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveLPrimeRubik\n".$red."Assert False\n".$white;
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
function assertMoveFPrimeRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveFPrimeValues, $moveFPrimeColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveFPrimeValues[$i] || 
      $rubik[$pos]['c'] !== $moveFPrimeColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveFPrimeRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveFPrimeRubik\n".$red."Assert False\n".$white;
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
function assertMoveBPrimeRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveBPrimeValues, $moveBPrimeColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveBPrimeValues[$i] || 
      $rubik[$pos]['c'] !== $moveBPrimeColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveBPrimeRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveBPrimeRubik\n".$red."Assert False\n".$white;
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
function assertMoveDPrimeRubik($rubik){
  global $red, $green, $yellow, $blue, $magenta, $white;
  global $positions, $moveDPrimeValues, $moveDPrimeColors;
  $assert = true;
  for ($i = 0;$i < count($positions);$i++){
    $pos = $positions[$i];
    if ($rubik[$pos]['v'] !== $moveDPrimeValues[$i] || 
      $rubik[$pos]['c'] !== $moveDPrimeColors[$i]){
      $assert = false;
    }
  }
  if ($assert){
    echo "assertMoveDPrimeRubik\n".$green."Assert True\n".$white;
  } else {
    echo "assertMoveDPrimeRubik\n".$red."Assert False\n".$white;
  }
}