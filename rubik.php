<?php
 $rubik = array();
 $sides = array('A','B','C','D','E','F');
 function resetRubik(){
  global $sides, $rubik;
  foreach ($sides as $side){
   for ($i = 1;$i<10;$i++){ 
    $rubik[$side.$i] = $side.$i;
   }
  }
 }
 resetRubik();
 var_dump($rubik);
