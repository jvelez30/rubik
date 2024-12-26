<?php
 /* Rubik Asserts */
 require_once __DIR__ . '/lib/assertAux.php';
 $red="\e[31m";
 $green="\e[32m"; 
 $yellow="\e[33m"; 
 $blue="\e[34m";
 $magenta="\e[95m"; 
 $white="\e[97m";
 
 function isOrderedRubik($rubik){
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
   return $assert;
 }

 function getOrderedRubik(){
   /* WITH TEST */
   global $red, $green, $yellow, $blue, $magenta, $white;
   $rubik = array();
   $sides = array('B','L','U','R','D','F');
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

 function setRubikByString($statePositions, $stateColors){
   global $red, $green, $yellow, $blue, $magenta, $white;
   $rubik = array();
   $statePositionsArray = str_split($statePositions,2);
   // Validate statePositions
   if (count ($statePositionsArray) != 54){
      return false;
   }
   $stateColorsArray = str_split($stateColors,1);
   // Validate stateColors
   if (count ($stateColorsArray) != 54){
      return false;
   }
   $sides = array('B','L','U','R','D','F');
   $count = 0;
   foreach ($sides as $side){
      for ($i = 1;$i<10;$i++){ 
         $rubik[$side.$i]['v'] = $statePositionsArray[$count];
         switch ($stateColorsArray[$count]){
         case 'Y':
            $rubik[$side.$i]['c'] =  $yellow;
            break;
         case 'M':
            $rubik[$side.$i]['c'] =  $magenta;
            break;
         case 'G':
            $rubik[$side.$i]['c'] =  $green;
            break;
         case 'W':
            $rubik[$side.$i]['c'] =  $white;
            break;
         case 'B':
            $rubik[$side.$i]['c'] =  $blue;
            break;
         case 'R':
            $rubik[$side.$i]['c'] =  $red;
            break;
         }
         $count = $count + 1;
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
    /* WITH TEST */
    // Rotate A Side
   $posToChange = array (
      'U1','U2','U3','U6','U9','U8','U7','U4',
      'R1','R4','R7','F1','F2','F3','L3','L6','L9',
      'B7','B8','B9'
   );
   $posNewVal = array (
      'U7','U4','U1','U2','U3','U6','U9','U8',
      'B7','B8','B9','R7','R4','R1','F1','F2','F3',
      'L9','L6','L3'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveUP($rubik){
   /* WITH TEST */
   // Rotate A Side in reverse
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
   /* WITH TEST */
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
 function moveRP($rubik){
   /* WITH TEST */
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
 function moveL($rubik){
   /* WITH TEST */
   // Rotate L Side
   $posToChange = array (
      'L1','L2','L3','L6','L9','L7','L8','L4',
      'B1','B4','B7','U1','U4','U7','F1','F4','F7',
      'D3','D6','D9'
   );
   $posNewVal = array (
      'L7','L4','L1','L2','L3','L9','L6','L8',
      'D9','D6','D3','B1','B4','B7','U1','U4','U7',
      'F7','F4','F1'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveLP($rubik){
   /* WITH TEST */
   // Rotate L Side in reverse
   $posToChange = array (
      'L1','L2','L3','L6','L9','L8','L7','L4',
      'B1','B4','B7','U1','U4','U7','F1','F4','F7',
      'D3','D6','D9'
   );
   $posNewVal = array (
      'L3','L6','L9','L8','L7','L4','L1','L2',
      'U1','U4','U7','F1','F4','F7','D9','D6','D3',
      'B7','B4','B1',
      
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveF($rubik){
   /* WITH TEST */
   // Rotate F Side
   $posToChange = array (
      'F1','F2','F3','F6','F9','F7','F8','F4',
      'L7','L8','L9','U7','U8','U9','R7','R8','R9',
      'D7','D8','D9'
   );
   $posNewVal = array (
      'F7','F4','F1','F2','F3','F9','F6','F8',
      'D7','D8','D9','L7','L8','L9','U7','U8','U9',
      'R7','R8','R9'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveFP($rubik){
   /* WITH TEST */
   // Rotate F Side in reverse
   $posToChange = array (
      'F1','F2','F3','F6','F9','F8','F7','F4',
      'L7','L8','L9','U7','U8','U9','R7','R8','R9',
      'D7','D8','D9'
   );
   $posNewVal = array (
      'F3','F6','F9','F8','F7','F4','F1','F2',
      'U7','U8','U9','R7','R8','R9','D7','D8','D9',
      'L7','L8','L9'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveB($rubik){
   /* WITH TEST */
   // Rotate B Side
   $posToChange = array (
      'B1','B2','B3','B6','B9','B8','B7','B4',
      'L1','L2','L3','U1','U2','U3','R1','R2','R3',
      'D1','D2','D3'
   );
   $posNewVal = array (
      'B7','B4','B1','B2','B3','B6','B9','B8',
      'U1','U2','U3','R1','R2','R3','D1','D2','D3',
      'L1','L2','L3'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveBP($rubik){
   /* WITH TEST */
   // Rotate B Side in reverse
   $posToChange = array (
      'B1','B2','B3','B6','B9','B8','B7','B4',
      'L1','L2','L3','U1','U2','U3','R1','R2','R3',
      'D1','D2','D3'
   );
   $posNewVal = array (
      'B3','B6','B9','B8','B7','B4','B1','B2',
      'D1','D2','D3','L1','L2','L3','U1','U2','U3',
      'R1','R2','R3'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveD($rubik){
   /* WITH TEST */
   // Rotate D Side
   $posToChange = array (
      'D1','D2','D3','D6','D9','D8','D7','D4',
      'B1','B2','B3','L1','L4','L7','R3','R6','R9',
      'F7','F8','F9'
   );
   $posNewVal = array (
      'D7','D4','D1','D2','D3','D6','D9','D8',
      'R3','R6','R9','B3','B2','B1','F9','F8','F7',
      'L1','L4','L7'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function moveDP($rubik){
   /* WITH TEST */
   // Rotate D Side in reverse
   $posToChange = array (
      'D1','D2','D3','D6','D9','D8','D7','D4',
      'B1','B2','B3','L1','L4','L7','R3','R6','R9',
      'F7','F8','F9'
   );
   $posNewVal = array (
      'D3','D6','D9','D8','D7','D4','D1','D2',
      'L7','L4','L1','F7','F8','F9','B1','B2','B3',
      'R9','R6','R3'
   );
   $rubik = rubikMove($rubik, $posToChange, $posNewVal);
   return $rubik;
 }
 function sexyMove($rubik){
   /* WITH TEST */
   $rubik=moveR($rubik);
   $rubik=moveU($rubik);
   $rubik=moveRP($rubik);
   $rubik=moveUP($rubik);
   return $rubik;
 }
 function printPos($rubik,$pos){
   global $red, $green, $yellow, $blue, $magenta, $white;
    return $rubik[$pos]['c'].$rubik[$pos]['v'].$white;
 }
 function formatCli($rubik){
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
 function printPosHtml($rubik,$pos){
   global $red, $green, $yellow, $blue, $magenta, $white;
   return '<span style="background-color: '.translateColorToHtml($rubik[$pos]['c']).'">'.$rubik[$pos]['v'].'</span>';
 }
 function printHeadHtml(){
   echo '<head>' . "\n";
   echo '  <meta charset="UTF-8">' . "\n";
   echo '  <meta name="viewport" content="width=device-width, initial-scale=1.0">' . "\n";
   echo '  <meta http-equiv="X-UA-Compatible" content="IE=edge">' . "\n";
   echo '  <title>' . "Rubik" . '</title>' . "\n";
   echo '</head>' . "\n";
 }
 function formatHtml($rubik){
   echo '<!DOCTYPE html>' . "\n";
   echo '<html lang="es">' . "\n";
   printHeadHtml();
   echo '<body>' . "\n";
   printRubik2DHtml($rubik);
   echo '</body>' . "\n";
   echo '</html>' . "\n";
 }
 function printRubik2DHtml($rubik){
   echo "<table>";
   echo " <tr><td></td><td>" . printPosHtml($rubik,'B1') . "-" . printPosHtml($rubik,'B2') . "-" . printPosHtml($rubik,'B3') . "</td></tr>";
   echo " <tr><td></td><td>" . printPosHtml($rubik,'B4') . "-" . printPosHtml($rubik,'B5') . "-" . printPosHtml($rubik,'B6') . "</td></tr>";
   echo " <tr><td></td><td>" . printPosHtml($rubik,'B7') . "-" . printPosHtml($rubik,'B8') . "-" . printPosHtml($rubik,'B9') . "</td></tr>";
   echo " <tr><td>" . printPosHtml($rubik,'L1') . "-" . printPosHtml($rubik,'L2') . "-" . printPosHtml($rubik,'L3') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'U1') . "-" . printPosHtml($rubik,'U2') . "-" . printPosHtml($rubik,'U3') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'R1') . "-" . printPosHtml($rubik,'R2') . "-" . printPosHtml($rubik,'R3') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'D1') . "-" . printPosHtml($rubik,'D2') . "-" . printPosHtml($rubik,'D3') . "</td></tr>";
   echo " <tr><td>" . printPosHtml($rubik,'L4') . "-" . printPosHtml($rubik,'L5') . "-" . printPosHtml($rubik,'L6') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'U4') . "-" . printPosHtml($rubik,'U5') . "-" . printPosHtml($rubik,'U6') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'R4') . "-" . printPosHtml($rubik,'R5') . "-" . printPosHtml($rubik,'R6') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'D4') . "-" . printPosHtml($rubik,'D5') . "-" . printPosHtml($rubik,'D6') . "</td></tr>";
   echo " <tr><td>" . printPosHtml($rubik,'L7') . "-" . printPosHtml($rubik,'L8') . "-" . printPosHtml($rubik,'L9') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'U7') . "-" . printPosHtml($rubik,'U8') . "-" . printPosHtml($rubik,'U9') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'R7') . "-" . printPosHtml($rubik,'R8') . "-" . printPosHtml($rubik,'R9') . "</td>";
   echo "     <td>" . printPosHtml($rubik,'D7') . "-" . printPosHtml($rubik,'D8') . "-" . printPosHtml($rubik,'D9') . "</td></tr>";
   echo " <tr><td></td><td>" . printPosHtml($rubik,'F1') . "-" . printPosHtml($rubik,'F2') . "-" . printPosHtml($rubik,'F3') . "</td></tr>";
   echo " <tr><td></td><td>" . printPosHtml($rubik,'F4') . "-" . printPosHtml($rubik,'F5') . "-" . printPosHtml($rubik,'F6') . "</td></tr>";
   echo " <tr><td></td><td>" . printPosHtml($rubik,'F7') . "-" . printPosHtml($rubik,'F8') . "-" . printPosHtml($rubik,'F9') . "</td></tr>";
   echo "</table>";
 }
 function translateColorToHtml($color){
   global $red, $green, $yellow, $blue, $magenta, $white;
   switch ($color){
      case $red:
         return "red";
         break;
      case $green:
         return "lightgreen";
         break;
      case $yellow:
         return "yellow";
         break;
      case $blue:
         return "lightblue";
         break;
      case $magenta:
         return "orange";
         break;
      case $white:
         return "white";
         break;
   }
 }
 //$rubik=getOrderedRubik();
 $statePositions  = "B1B2B3B4B5B6B7B8B9L1L2L3L4L5L6L7L8L9";
 $statePositions  = $statePositions . "U1U2U3U4U5U6U7U8U9R1R2R3R4R5R6R7R8R9";
 $statePositions  = $statePositions . "D1D2D3D4D5D6D7D8D9F1F2F3F4F5F6F7F8F9";
 $stateColors     = "MMMMMMMMMBBBBBBBBBYYYYYYYYY";
 $stateColors     = $stateColors . "GGGGGGGGGWWWWWWWWWRRRRRRRRR";
 $rubik = setRubikByString($statePositions, $stateColors);
 /*
 for ($j=0;$j<6;$j++){
   $rubik=sexyMove($rubik);
 }
 */
 // Filter get var
 $movesParam = filter_input(INPUT_GET, 'm', FILTER_SANITIZE_SPECIAL_CHARS);
 $stateParam = filter_input(INPUT_GET, 's', FILTER_SANITIZE_SPECIAL_CHARS);
 // If moves are setted via string query
 if(isset($movesParam)){
   $moves = str_split($movesParam);
   foreach ($moves as $move){
      switch ($move){
         case 'U':
            $rubik=moveU($rubik);
            break;
         case 'u':
            $rubik=moveUP($rubik);
            break;
         case 'R':
            $rubik=moveR($rubik);
            break;
         case 'r':
            $rubik=moveRP($rubik);
            break;
         case 'L':
            $rubik=moveL($rubik);
            break;
         case 'l':
            $rubik=moveLP($rubik);
            break;
         case 'F':
            $rubik=moveF($rubik);
            break;
         case 'f':
            $rubik=moveFP($rubik);
            break;
         case 'B':
            $rubik=moveB($rubik);
            break;
         case 'b':
            $rubik=moveBP($rubik);
            break; 
         case 'D':
            $rubik=moveD($rubik);
            break;
         case 'd':
            $rubik=moveDP($rubik);
            break;     
      }
   }
 }

 //$rubik=moveDP($rubik);

 if (isset($argv[0])){
   // if command line cli format is displayed
   formatCli($rubik);
 } else {
   if (isset($stateParam) && $stateParam != ""){
      if (isOrderedRubik($rubik)){
         echo "Ordered";
      } else {
         echo "Unordered";
      }
   } else {
      // else html format is displayed
      formatHtml($rubik);
   }
 }

 //var_dump($rubik);
 
