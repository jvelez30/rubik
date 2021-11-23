<?php
//Asserts Auxiliar Structures and values
$red="\e[31m";
$green="\e[32m"; 
$yellow="\e[33m"; 
$blue="\e[34m";
$magenta="\e[95m"; 
$white="\e[97m";
$positions = array (
  'B1','B2','B3','B4','B5','B6','B7','B8','B9',
  'L1','L2','L3','L4','L5','L6','L7','L8','L9',
  'U1','U2','U3','U4','U5','U6','U7','U8','U9',
  'R1','R2','R3','R4','R5','R6','R7','R8','R9',
  'D1','D2','D3','D4','D5','D6','D7','D8','D9',
  'F1','F2','F3','F4','F5','F6','F7','F8','F9'
);
$orderedValues = array (
  'B1','B2','B3','B4','B5','B6','B7','B8','B9',
  'L1','L2','L3','L4','L5','L6','L7','L8','L9',
  'U1','U2','U3','U4','U5','U6','U7','U8','U9',
  'R1','R2','R3','R4','R5','R6','R7','R8','R9',
  'D1','D2','D3','D4','D5','D6','D7','D8','D9',
  'F1','F2','F3','F4','F5','F6','F7','F8','F9'
);
$orderedColors = array (
  $magenta,$magenta,$magenta,$magenta,$magenta,$magenta,$magenta,$magenta,$magenta,
  $blue,$blue,$blue,$blue,$blue,$blue,$blue,$blue,$blue,
  $yellow,$yellow,$yellow,$yellow,$yellow,$yellow,$yellow,$yellow,$yellow,
  $green,$green,$green,$green,$green,$green,$green,$green,$green,
  $white,$white,$white,$white,$white,$white,$white,$white,$white,
  $red,$red,$red,$red,$red,$red,$red,$red,$red
);
$moveUValues = array (
  'B1','B2','B3','B4','B5','B6','L9','L6','L3',
  'L1','L2','F1','L4','L5','F2','L7','L8','F3',
  'U7','U4','U1','U8','U5','U2','U9','U6','U3',
  'B7','R2','R3','B8','R5','R6','B9','R8','R9',
  'D1','D2','D3','D4','D5','D6','D7','D8','D9',
  'R7','R4','R1','F4','F5','F6','F7','F8','F9'
);
$moveUColors = array (
  $magenta,$magenta,$magenta,$magenta,$magenta,$magenta,$blue,$blue,$blue,
  $blue,$blue,$red,$blue,$blue,$red,$blue,$blue,$red,
  $yellow,$yellow,$yellow,$yellow,$yellow,$yellow,$yellow,$yellow,$yellow,
  $magenta,$green,$green,$magenta,$green,$green,$magenta,$green,$green,
  $white,$white,$white,$white,$white,$white,$white,$white,$white,
  $green,$green,$green,$red,$red,$red,$red,$red,$red
);