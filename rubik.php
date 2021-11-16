<?php
 function getOrderedRubik(){
  $rubik = array();
  $sides = array('A','B','C','D','E','F');
  foreach ($sides as $side){
   for ($i = 1;$i<10;$i++){ 
    $rubik[$side.$i] = $side.$i;
   }
  }
  return $rubik;
 }
 function moveA($rubik){
    // Rotate A Side
   $tempRubik = $rubik;
   $rubik['A1'] = $tempRubik['A7'];
   $rubik['A2'] = $tempRubik['A4'];
   $rubik['A3'] = $tempRubik['A1'];
   $rubik['A3'] = $tempRubik['A1'];
   $rubik['A6'] = $tempRubik['A2'];
   $rubik['A9'] = $tempRubik['A3'];
   $rubik['A7'] = $tempRubik['A9'];
   $rubik['A8'] = $tempRubik['A6'];
   $rubik['A4'] = $tempRubik['A8'];
   // Shift other sides
   $rubik['C1'] = $tempRubik['B7'];
   $rubik['C4'] = $tempRubik['B8'];
   $rubik['C7'] = $tempRubik['B9'];
   $rubik['F1'] = $tempRubik['C7'];
   $rubik['F2'] = $tempRubik['C4'];
   $rubik['F3'] = $tempRubik['C1'];
   $rubik['E3'] = $tempRubik['F1'];
   $rubik['E6'] = $tempRubik['F2'];
   $rubik['E9'] = $tempRubik['F3'];
   $rubik['B7'] = $tempRubik['E9'];
   $rubik['B8'] = $tempRubik['E6'];
   $rubik['B9'] = $tempRubik['E3'];
   return $rubik;
 }
 function moveAPrime($rubik){
   // Rotate A Side
   $tempRubik = $rubik;
   $rubik['A1'] = $tempRubik['A3'];
   $rubik['A2'] = $tempRubik['A6'];
   $rubik['A3'] = $tempRubik['A9'];
   $rubik['A6'] = $tempRubik['A8'];
   $rubik['A9'] = $tempRubik['A7'];
   $rubik['A8'] = $tempRubik['A4'];
   $rubik['A7'] = $tempRubik['A1'];
   $rubik['A4'] = $tempRubik['A2'];
   // Shift other sides
   $rubik['C1'] = $tempRubik['F3'];
   $rubik['C4'] = $tempRubik['F2'];
   $rubik['C7'] = $tempRubik['F1'];
   $rubik['F1'] = $tempRubik['E3'];
   $rubik['F2'] = $tempRubik['E6'];
   $rubik['F3'] = $tempRubik['E9'];
   $rubik['E3'] = $tempRubik['B9'];
   $rubik['E6'] = $tempRubik['B8'];
   $rubik['E9'] = $tempRubik['B7'];
   $rubik['B7'] = $tempRubik['C1'];
   $rubik['B8'] = $tempRubik['C4'];
   $rubik['B9'] = $tempRubik['C7'];
   return $rubik;
 }
 function moveC($rubik){
   // Rotate C Side
   $tempRubik = $rubik;
   $rubik['C1'] = $tempRubik['C7'];
   $rubik['C2'] = $tempRubik['C4'];
   $rubik['C3'] = $tempRubik['C1'];
   $rubik['C3'] = $tempRubik['C1'];
   $rubik['C6'] = $tempRubik['C2'];
   $rubik['C9'] = $tempRubik['C3'];
   $rubik['C7'] = $tempRubik['C9'];
   $rubik['C8'] = $tempRubik['C6'];
   $rubik['C4'] = $tempRubik['C8'];
   // Shift other sides
   $rubik['B3'] = $tempRubik['A3'];
   $rubik['B6'] = $tempRubik['A6'];
   $rubik['B9'] = $tempRubik['A9'];
   $rubik['D7'] = $tempRubik['B3'];
   $rubik['D4'] = $tempRubik['B6'];
   $rubik['D1'] = $tempRubik['B9'];
   $rubik['F3'] = $tempRubik['D7'];
   $rubik['F6'] = $tempRubik['D4'];
   $rubik['F9'] = $tempRubik['D1'];
   $rubik['A3'] = $tempRubik['F3'];
   $rubik['A6'] = $tempRubik['F6'];
   $rubik['A9'] = $tempRubik['F9'];
   return $rubik;
 }
 function moveCPrime($rubik){
   // Rotate C Side in reverse
   $tempRubik = $rubik;
   $rubik['C1'] = $tempRubik['C3'];
   $rubik['C2'] = $tempRubik['C6'];
   $rubik['C3'] = $tempRubik['C9'];
   $rubik['C6'] = $tempRubik['C8'];
   $rubik['C9'] = $tempRubik['C7'];
   $rubik['C8'] = $tempRubik['C4'];
   $rubik['C7'] = $tempRubik['C1'];
   $rubik['C4'] = $tempRubik['C2'];
   // Shift other sides
   $rubik['B3'] = $tempRubik['D7'];
   $rubik['B6'] = $tempRubik['D4'];
   $rubik['B9'] = $tempRubik['D1'];
   $rubik['D7'] = $tempRubik['F3'];
   $rubik['D4'] = $tempRubik['F6'];
   $rubik['D1'] = $tempRubik['F9'];
   $rubik['F3'] = $tempRubik['A3'];
   $rubik['F6'] = $tempRubik['A6'];
   $rubik['F9'] = $tempRubik['A9'];
   $rubik['A3'] = $tempRubik['B3'];
   $rubik['A6'] = $tempRubik['B6'];
   $rubik['A9'] = $tempRubik['B9'];
   return $rubik;
 }
 function sexyMove($rubik){
   $rubik=moveC($rubik);
   $rubik=moveA($rubik);
   $rubik=moveCPrime($rubik);
   $rubik=moveAPrime($rubik);
   return $rubik;
 }
 function format($rubik){
   echo "         |--------|" ."\n";
   echo "         |" . $rubik['B1'] . "-" . $rubik['B2'] . "-" . $rubik['B3'] . "|" ."\n";
   echo "         |" . $rubik['B4'] . "-" . $rubik['B5'] . "-" . $rubik['B6'] . "|" ."\n";
   echo "         |" . $rubik['B7'] . "-" . $rubik['B8'] . "-" . $rubik['B9'] . "|" ."\n";
   echo "|--------|--------|--------|--------|" ."\n";
   echo "|" . $rubik['E1'] . "-" . $rubik['E2'] . "-" . $rubik['E3'] . "|";
   echo $rubik['A1'] . "-" . $rubik['A2'] . "-" . $rubik['A3'];
   echo "|" . $rubik['C1'] . "-" . $rubik['C2'] . "-" . $rubik['C3'] . "|";
   echo $rubik['D1'] . "-" . $rubik['D2'] . "-" . $rubik['D3']. "|" . "\n";
   echo "|" . $rubik['E4'] . "-" . $rubik['E5'] . "-" . $rubik['E6'] . "|";
   echo $rubik['A4'] . "-" . $rubik['A5'] . "-" . $rubik['A6'];
   echo "|" . $rubik['C4'] . "-" . $rubik['C5'] . "-" . $rubik['C6'] . "|";
   echo $rubik['D4'] . "-" . $rubik['D5'] . "-" . $rubik['D6']. "|" . "\n";
   echo "|" . $rubik['E7'] . "-" . $rubik['E8'] . "-" . $rubik['E9'] . "|";
   echo $rubik['A7'] . "-" . $rubik['A8'] . "-" . $rubik['A9'];
   echo "|" . $rubik['C7'] . "-" . $rubik['C8'] . "-" . $rubik['C9'] . "|";
   echo $rubik['D7'] . "-" . $rubik['D8'] . "-" . $rubik['D9']. "|" . "\n";
   echo "|--------|--------|--------|--------|" ."\n";
   echo "         |" . $rubik['F1'] . "-" . $rubik['F2'] . "-" . $rubik['F3'] . "|" ."\n";
   echo "         |" . $rubik['F4'] . "-" . $rubik['F5'] . "-" . $rubik['F6'] . "|" ."\n";
   echo "         |" . $rubik['F7'] . "-" . $rubik['F8'] . "-" . $rubik['F9'] . "|" ."\n";
   echo "         |--------|" ."\n";
 }
 $rubik=getOrderedRubik();
 
 for ($j=0;$j<6;$j++){
   $rubik=sexyMove($rubik);
 }
 
 //$rubik=moveCPrime($rubik);
 format($rubik);
 // var_dump($rubik);
