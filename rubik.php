<?php
 $red="\e[31m";
 $green="\e[32m"; 
 $yellow="\e[33m"; 
 $blue="\e[34m";
 $magenta="\e[95m"; 
 $white="\e[97m";
 function getOrderedRubik(){
  global $red, $green, $yellow, $blue, $magenta, $white;
  $rubik = array();
  $sides = array('U','B','R','D','L','F');
  foreach ($sides as $side){
   for ($i = 1;$i<10;$i++){ 
    $rubik[$side.$i]['v'] = $side.$i;
    switch ($side){
      case 'U':
         $rubik[$side.$i]['c'] =  $yellow;
         break;
      case 'B':
         $rubik[$side.$i]['c'] =  $magenta;
         break;
      case 'R':
         $rubik[$side.$i]['c'] =  $green;
         break;
      case 'D':
         $rubik[$side.$i]['c'] =  $white;
         break;
      case 'L':
         $rubik[$side.$i]['c'] =  $blue;
         break;
      case 'F':
         $rubik[$side.$i]['c'] =  $red;
         break;
    }
   }
  }
  return $rubik;
 }
 function rubikMove($rubik, $posA,$posN){
   $tempRubik = $rubik;
    for ($i = 0;$i < count($posA);$i++){
      $posActual = $posA[$i];
      $posNew = $posN[$i];
      $rubik[$posActual]['v'] = $tempRubik[$posNew]['v'];
      $rubik[$posActual]['c'] = $tempRubik[$posNew]['c'];
    }
    return $rubik;
 }
 function moveU($rubik){
    // Rotate A Side
   $posToChange = array (
      'U1','U2','U3','U6','U9','U7','U8','U4',
      'R1','R4','R7','F1','F2','F3','L3','L6','L9',
      'B7','B8','B9'
   );
   $posNewVal = array (
      'U7','U4','U1','U2','U3','U9','U6','U8',
      'B7','B8','B9','R7','R4','R1','F1','F2','F3',
      'L9','L6','L3'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveUPrime($rubik){
   // Rotate A Side
   $posToChange = array (
      'U1','U2','U3','U6','U9','U8','U7','U4',
      'R1','R4','R7','F1','F2','F3','L3','L6','L9',
      'B7','B8','B9'
   );
   $posNewVal = array (
      'U3','U6','U9','U8','U7','U4','U1','U2',
      'F3','F2','F1','L3','L6','L9','B9','B8','B7',
      'R1','R4','R7'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveR($rubik){
   // Rotate R Side
   $posToChange = array (
      'R1','R2','R3','R6','R9','R7','R8','R4',
      'B3','B6','B9','D7','D4','D1','F3','F6','F9',
      'U3','U6','U9'
   );
   $posNewVal = array (
      'R7','R4','R1','R2','R3','R9','R6','R8',
      'U3','U6','U9','B3','B6','B9','D7','D4','D1',
      'F3','F6','F9'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveRPrime($rubik){
   // Rotate R Side in reverse
   $posToChange = array (
      'R1','R2','R3','R6','R9','R8','R7','R4',
      'B3','B6','B9','D7','D4','D1','F3','F6','F9',
      'U3','U6','U9'
   );
   $posNewVal = array (
      'R3','R6','R9','R8','R7','R4','R1','R2',
      'D7','D4','D1','F3','F6','F9','U3','U6','U9',
      'B3','B6','B9'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function sexyMove($rubik){
   $rubik=moveR($rubik);
   $rubik=moveU($rubik);
   $rubik=moveRPrime($rubik);
   $rubik=moveUPrime($rubik);
   return $rubik;
 }
 function printPos($rubik,$pos){
   global $red, $green, $yellow, $blue, $magenta, $white;
    return $rubik[$pos]['c'].$rubik[$pos]['v'].$white;
 }
 function format($rubik){
   echo "         |--------|" ."\n";
   echo "         |" . printPos($rubik,'B1') . "-" . printPos($rubik,'B2') . "-" . printPos($rubik,'B3') . "|" ."\n";
   echo "         |" . printPos($rubik,'B4') . "-" . printPos($rubik,'B5') . "-" . printPos($rubik,'B6') . "|" ."\n";
   echo "         |" . printPos($rubik,'B7') . "-" . printPos($rubik,'B8') . "-" . printPos($rubik,'B9') . "|" ."\n";
   echo "|--------|--------|--------|--------|" ."\n";
   echo "|" . printPos($rubik,'L1') . "-" . printPos($rubik,'L2') . "-" . printPos($rubik,'L3') . "|";
   echo printPos($rubik,'U1') . "-" . printPos($rubik,'U2') . "-" . printPos($rubik,'U3');
   echo "|" . printPos($rubik,'R1') . "-" . printPos($rubik,'R2') . "-" . printPos($rubik,'R3') . "|";
   echo printPos($rubik,'D1') . "-" . printPos($rubik,'D2') . "-" . printPos($rubik,'D3'). "|" . "\n";
   echo "|" . printPos($rubik,'L4') . "-" . printPos($rubik,'L5') . "-" . printPos($rubik,'L6') . "|";
   echo printPos($rubik,'U4') . "-" . printPos($rubik,'U5') . "-" . printPos($rubik,'U6');
   echo "|" . printPos($rubik,'R4') . "-" . printPos($rubik,'R5') . "-" . printPos($rubik,'R6') . "|";
   echo printPos($rubik,'D4') . "-" . printPos($rubik,'D5') . "-" . printPos($rubik,'D6'). "|" . "\n";
   echo "|" . printPos($rubik,'L7') . "-" . printPos($rubik,'L8') . "-" . printPos($rubik,'L9') . "|";
   echo printPos($rubik,'U7') . "-" . printPos($rubik,'U8') . "-" . printPos($rubik,'U9');
   echo "|" . printPos($rubik,'R7') . "-" . printPos($rubik,'R8') . "-" . printPos($rubik,'R9') . "|";
   echo printPos($rubik,'D7') . "-" . printPos($rubik,'D8') . "-" . printPos($rubik,'D9'). "|" . "\n";
   echo "|--------|--------|--------|--------|" ."\n";
   echo "         |" . printPos($rubik,'F1') . "-" . printPos($rubik,'F2') . "-" . printPos($rubik,'F3') . "|" ."\n";
   echo "         |" . printPos($rubik,'F4') . "-" . printPos($rubik,'F5') . "-" . printPos($rubik,'F6') . "|" ."\n";
   echo "         |" . printPos($rubik,'F7') . "-" . printPos($rubik,'F8') . "-" . printPos($rubik,'F9') . "|" ."\n";
   echo "         |--------|" ."\n";
 }
 $rubik=getOrderedRubik();
 
 for ($j=0;$j<6;$j++){
   $rubik=sexyMove($rubik);
 }
 
 //$rubik=moveRPrime($rubik);
 format($rubik);
 // var_dump($rubik);
