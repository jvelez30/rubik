//Asserts Auxiliar Structures and values
package assertAux
// red := "\033[31m"
// green := "\033[32m" 
// yellow := "\033[33m" 
// blue := "\033[34m"
// magenta := "\033[95m" 
// white := "\033[97m"
red := "red"
green := "green"
yellow := "yellow"
blue := "blue"
magenta := "magenta"
white := "white"
positions := [] string {
  "B1","B2","B3","B4","B5","B6","B7","B8","B9",
  "L1","L2","L3","L4","L5","L6","L7","L8","L9",
  "U1","U2","U3","U4","U5","U6","U7","U8","U9",
  "R1","R2","R3","R4","R5","R6","R7","R8","R9",
  "D1","D2","D3","D4","D5","D6","D7","D8","D9",
  "F1","F2","F3","F4","F5","F6","F7","F8","F9"
}
orderedValues := [] string {
  "B1","B2","B3","B4","B5","B6","B7","B8","B9",
  "L1","L2","L3","L4","L5","L6","L7","L8","L9",
  "U1","U2","U3","U4","U5","U6","U7","U8","U9",
  "R1","R2","R3","R4","R5","R6","R7","R8","R9",
  "D1","D2","D3","D4","D5","D6","D7","D8","D9",
  "F1","F2","F3","F4","F5","F6","F7","F8","F9"
}
orderedColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
  blue,blue,blue,blue,blue,blue,blue,blue,blue,
  yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
  green,green,green,green,green,green,green,green,green,
  white,white,white,white,white,white,white,white,white,
  red,red,red,red,red,red,red,red,red
}
moveUValues := [] string {
  "B1","B2","B3","B4","B5","B6","L9","L6","L3",
  "L1","L2","F1","L4","L5","F2","L7","L8","F3",
  "U7","U4","U1","U8","U5","U2","U9","U6","U3",
  "B7","R2","R3","B8","R5","R6","B9","R8","R9",
  "D1","D2","D3","D4","D5","D6","D7","D8","D9",
  "R7","R4","R1","F4","F5","F6","F7","F8","F9"
}
moveUColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,blue,blue,blue,
  blue,blue,red,blue,blue,red,blue,blue,red,
  yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
  magenta,green,green,magenta,green,green,magenta,green,green,
  white,white,white,white,white,white,white,white,white,
  green,green,green,red,red,red,red,red,red
}
moveUPValues := [] string {
  "B1","B2","B3","B4","B5","B6","R1","R4","R7",
  "L1","L2","B9","L4","L5","B8","L7","L8","B7",
  "U3","U6","U9","U2","U5","U8","U1","U4","U7",
  "F3","R2","R3","F2","R5","R6","F1","R8","R9",
  "D1","D2","D3","D4","D5","D6","D7","D8","D9",
  "L3","L6","L9","F4","F5","F6","F7","F8","F9"
}
moveUPColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,green,green,green,
  blue,blue,magenta,blue,blue,magenta,blue,blue,magenta,
  yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
  red,green,green,red,green,green,red,green,green,
  white,white,white,white,white,white,white,white,white,
  blue,blue,blue,red,red,red,red,red,red
}
moveRValues := [] string {
  "B1","B2","U3","B4","B5","U6","B7","B8","U9",
  "L1","L2","L3","L4","L5","L6","L7","L8","L9",
  "U1","U2","F3","U4","U5","F6","U7","U8","F9",
  "R7","R4","R1","R8","R5","R2","R9","R6","R3",
  "B9","D2","D3","B6","D5","D6","B3","D8","D9",
  "F1","F2","D7","F4","F5","D4","F7","F8","D1"
}
moveRColors := [] string {
  magenta,magenta,yellow,magenta,magenta,yellow,magenta,magenta,yellow,
  blue,blue,blue,blue,blue,blue,blue,blue,blue,
  yellow,yellow,red,yellow,yellow,red,yellow,yellow,red,
  green,green,green,green,green,green,green,green,green,
  magenta,white,white,magenta,white,white,magenta,white,white,
  red,red,white,red,red,white,red,red,white
}
moveRPValues := [] string {
  "B1","B2","D7","B4","B5","D4","B7","B8","D1",
  "L1","L2","L3","L4","L5","L6","L7","L8","L9",
  "U1","U2","B3","U4","U5","B6","U7","U8","B9",
  "R3","R6","R9","R2","R5","R8","R1","R4","R7",
  "F9","D2","D3","F6","D5","D6","F3","D8","D9",
  "F1","F2","U3","F4","F5","U6","F7","F8","U9"
}
moveRPColors := [] string {
  magenta,magenta,white,magenta,magenta,white,magenta,magenta,white,
  blue,blue,blue,blue,blue,blue,blue,blue,blue,
  yellow,yellow,magenta,yellow,yellow,magenta,yellow,yellow,magenta,
  green,green,green,green,green,green,green,green,green,
  red,white,white,red,white,white,red,white,white,
  red,red,yellow,red,red,yellow,red,red,yellow
}
moveLValues := [] string {
  "D9","B2","B3","D6","B5","B6","D3","B8","B9",
  "L7","L4","L1","L8","L5","L2","L9","L6","L3",
  "B1","U2","U3","B4","U5","U6","B7","U8","U9",
  "R1","R2","R3","R4","R5","R6","R7","R8","R9",
  "D1","D2","F7","D4","D5","F4","D7","D8","F1",
  "U1","F2","F3","U4","F5","F6","U7","F8","F9"
}
moveLColors := [] string {
  white,magenta,magenta,white,magenta,magenta,white,magenta,magenta,
  blue,blue,blue,blue,blue,blue,blue,blue,blue,
  magenta,yellow,yellow,magenta,yellow,yellow,magenta,yellow,yellow,
  green,green,green,green,green,green,green,green,green,
  white,white,red,white,white,red,white,white,red,
  yellow,red,red,yellow,red,red,yellow,red,red
}
moveLPValues := [] string {
  "U1","B2","B3","U4","B5","B6","U7","B8","B9",
  "L3","L6","L9","L2","L5","L8","L1","L4","L7",
  "F1","U2","U3","F4","U5","U6","F7","U8","U9",
  "R1","R2","R3","R4","R5","R6","R7","R8","R9",
  "D1","D2","B7","D4","D5","B4","D7","D8","B1",
  "D9","F2","F3","D6","F5","F6","D3","F8","F9"
}
moveLPColors := [] string {
  yellow,magenta,magenta,yellow,magenta,magenta,yellow,magenta,magenta,
  blue,blue,blue,blue,blue,blue,blue,blue,blue,
  red,yellow,yellow,red,yellow,yellow,red,yellow,yellow,
  green,green,green,green,green,green,green,green,green,
  white,white,magenta,white,white,magenta,white,white,magenta,
  white,red,red,white,red,red,white,red,red
}
moveFValues := [] string {
  "B1","B2","B3","B4","B5","B6","B7","B8","B9",
  "L1","L2","L3","L4","L5","L6","D7","D8","D9",
  "U1","U2","U3","U4","U5","U6","L7","L8","L9",
  "R1","R2","R3","R4","R5","R6","U7","U8","U9",
  "D1","D2","D3","D4","D5","D6","R7","R8","R9",
  "F7","F4","F1","F8","F5","F2","F9","F6","F3"
}
moveFColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
  blue,blue,blue,blue,blue,blue,white,white,white,
  yellow,yellow,yellow,yellow,yellow,yellow,blue,blue,blue,
  green,green,green,green,green,green,yellow,yellow,yellow,
  white,white,white,white,white,white,green,green,green,
  red,red,red,red,red,red,red,red,red
}
moveFPValues := [] string {
  "B1","B2","B3","B4","B5","B6","B7","B8","B9",
  "L1","L2","L3","L4","L5","L6","U7","U8","U9",
  "U1","U2","U3","U4","U5","U6","R7","R8","R9",
  "R1","R2","R3","R4","R5","R6","D7","D8","D9",
  "D1","D2","D3","D4","D5","D6","L7","L8","L9",
  "F3","F6","F9","F2","F5","F8","F1","F4","F7"
}
moveFPColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
  blue,blue,blue,blue,blue,blue,yellow,yellow,yellow,
  yellow,yellow,yellow,yellow,yellow,yellow,green,green,green,
  green,green,green,green,green,green,white,white,white,
  white,white,white,white,white,white,blue,blue,blue,
  red,red,red,red,red,red,red,red,red
}
moveBValues := [] string {
  "B7","B4","B1","B8","B5","B2","B9","B6","B3",
  "U1","U2","U3","L4","L5","L6","L7","L8","L9",
  "R1","R2","R3","U4","U5","U6","U7","U8","U9",
  "D1","D2","D3","R4","R5","R6","R7","R8","R9",
  "L1","L2","L3","D4","D5","D6","D7","D8","D9",
  "F1","F2","F3","F4","F5","F6","F7","F8","F9"
}
moveBColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
  yellow,yellow,yellow,blue,blue,blue,blue,blue,blue,
  green,green,green,yellow,yellow,yellow,yellow,yellow,yellow,
  white,white,white,green,green,green,green,green,green,
  blue,blue,blue,white,white,white,white,white,white,
  red,red,red,red,red,red,red,red,red
}
moveBPValues := [] string {
  "B3","B6","B9","B2","B5","B8","B1","B4","B7",
  "D1","D2","D3","L4","L5","L6","L7","L8","L9",
  "L1","L2","L3","U4","U5","U6","U7","U8","U9",
  "U1","U2","U3","R4","R5","R6","R7","R8","R9",
  "R1","R2","R3","D4","D5","D6","D7","D8","D9",
  "F1","F2","F3","F4","F5","F6","F7","F8","F9"
}
moveBPColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
  white,white,white,blue,blue,blue,blue,blue,blue,
  blue,blue,blue,yellow,yellow,yellow,yellow,yellow,yellow,
  yellow,yellow,yellow,green,green,green,green,green,green,
  green,green,green,white,white,white,white,white,white,
  red,red,red,red,red,red,red,red,red
}
moveDValues := [] string {
  "R3","R6","R9","B4","B5","B6","B7","B8","B9",
  "B3","L2","L3","B2","L5","L6","B1","L8","L9",
  "U1","U2","U3","U4","U5","U6","U7","U8","U9",
  "R1","R2","F9","R4","R5","F8","R7","R8","F7",
  "D7","D4","D1","D8","D5","D2","D9","D6","D3",
  "F1","F2","F3","F4","F5","F6","L1","L4","L7"
}
moveDColors := [] string {
  green,green,green,magenta,magenta,magenta,magenta,magenta,magenta,
  magenta,blue,blue,magenta,blue,blue,magenta,blue,blue,
  yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
  green,green,red,green,green,red,green,green,red,
  white,white,white,white,white,white,white,white,white,
  red,red,red,red,red,red,blue,blue,blue
}
moveDPValues := [] string {
  "L7","L4","L1","B4","B5","B6","B7","B8","B9",
  "F7","L2","L3","F8","L5","L6","F9","L8","L9",
  "U1","U2","U3","U4","U5","U6","U7","U8","U9",
  "R1","R2","B1","R4","R5","B2","R7","R8","B3",
  "D3","D6","D9","D2","D5","D8","D1","D4","D7",
  "F1","F2","F3","F4","F5","F6","R9","R6","R3"
}
moveDPColors := [] string {
  blue,blue,blue,magenta,magenta,magenta,magenta,magenta,magenta,
  red,blue,blue,red,blue,blue,red,blue,blue,
  yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
  green,green,magenta,green,green,magenta,green,green,magenta,
  white,white,white,white,white,white,white,white,white,
  red,red,red,red,red,red,green,green,green
}
sexyMoveValues := [] string {
  "B1","B2","B3","B4","B5","B6","R1","R4","B7",
  "L1","L2","B9","L4","L5","L6","L7","L8","L9",
  "U3","U6","L3","U4","U5","F6","U7","U8","F9",
  "U1","R2","R3","R8","R5","R6","R9","B8","U9",
  "D1","D2","D3","D4","D5","D6","R7","D8","D9",
  "F1","F2","D7","F4","F5","U2","F7","F8","F3"
}
sexyMoveColors := [] string {
  magenta,magenta,magenta,magenta,magenta,magenta,green,green,magenta,
  blue,blue,magenta,blue,blue,blue,blue,blue,blue,
  yellow,yellow,blue,yellow,yellow,red,yellow,yellow,red,
  yellow,green,green,green,green,green,green,magenta,yellow,
  white,white,white,white,white,white,green,white,white,
  red,red,white,red,red,yellow,red,red,red
}