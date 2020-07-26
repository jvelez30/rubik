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
 function format($rubik){
  echo "         |--------|" ."\n";
  echo "         |" . $rubik['B1'] . "-" . $rubik['B2'] . "-" . $rubik['B3'] . "|" ."\n";
  echo "         |" . $rubik['B4'] . "-" . $rubik['B5'] . "-" . $rubik['B6'] . "|" ."\n";
  echo "         |" . $rubik['B7'] . "-" . $rubik['B8'] . "-" . $rubik['B9'] . "|" ."\n";
 }
 $rubik=getOrderedRubik();
 format($rubik);
// var_dump($rubik);
