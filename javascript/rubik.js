
var red = "\x1b[31m";
var green = "\x1b[32m";
var yellow = "\x1b[33m";
var blue = "\x1b[34m";
var magenta = "\x1b[95m";
var white = "\x1b[97m";

function getOrderedRubik()
{
    /* WITH TEST */
    //global red, green, yellow, blue, magenta, white;
    var rubikValues = Array();
    var rubikColors = Array();
    var sides = ['U','B','R','D','L','F']
    sides.forEach(side => {
        for (var i = 1; i < 10; i++) {
            rubikValues[side+i] = side+i;
            switch (side) {
                case 'U':
                    rubikColors[side+i] = yellow;
                    break;
                case 'B':
                    rubikColors[side+i] = magenta;
                    break;
                case 'R':
                    rubikColors[side+i] = green;
                    break;
                case 'D':
                    rubikColors[side+i] = white;
                    break;
                case 'L':
                    rubikColors[side+i] = blue;
                    break;
                case 'F':
                    rubikColors[side+i] = red;
                    break;
            }
        }
    })
    var rubik = Array(rubikValues,rubikColors)
    return rubik;
}
function copyRubik(rubik){
    var newRubik = Array(Array(),Array());
    for (var i= 0; i < rubik[0].length; i++) {
        newRubik[0][i] = rubik[0][i]
        newRubik[1][i] = rubik[1][i]
    }
    return newRubik
}
function rubikMove(rubik, posA, posN)
{
    var tempRubik = copyRubik(rubik);
    for (var i = 0; i < posA.length; i++) {
        var posActual = posA[i];
        var posNew = posN[i];
        rubik[0][posActual] = tempRubik[0][posNew];
        rubik[1][posActual] = tempRubik[1][posNew];
    }
    return rubik;
}
function moveU(rubik)
{
    /* WITH TEST */
    // Rotate A Side
    var movedRubik = copyRubik(rubik);
    var posToChange = ['U1','U2','U3','U6','U9','U8','U7','U4','R1','R4','R7','F1','F2','F3','L3','L6','L9','B7','B8','B9'];
    var posNewVal = ['U7','U4','U1','U2','U3','U6','U9','U8','B7','B8','B9','R7','R4','R1','F1','F2','F3','L9','L6','L3'];
    movedRubik = rubikMove(movedRubik, posToChange, posNewVal);
    return movedRubik;
}
function moveUP(rubik)
{
    /* WITH TEST */
    // Rotate A Side in reverse
    var movedRubik = copyRubik(rubik);
    var posToChange = ['U1','U2','U3','U6','U9','U8','U7','U4','R1','R4','R7','F1','F2','F3','L3','L6','L9','B7','B8','B9'];
    var posNewVal = ['U3','U6','U9','U8','U7','U4','U1','U2','F3','F2','F1','L3','L6','L9','B9','B8','B7','R1','R4','R7'];
    movedRubik = rubikMove(movedRubik, posToChange, posNewVal);
    return movedRubik;
}
function moveR(rubik)
{
    /* WITH TEST */
    // Rotate R Side
    var movedRubik = copyRubik(rubik);
    var posToChange = ['R1','R2','R3','R6','R9','R7','R8','R4','B3','B6','B9','D7','D4','D1','F3','F6','F9','U3','U6','U9'];
    var posNewVal = ['R7','R4','R1','R2','R3','R9','R6','R8','U3','U6','U9','B3','B6','B9','D7','D4','D1','F3','F6','F9'];
    movedRubik = rubikMove(movedRubik, posToChange, posNewVal);
    return movedRubik;
}
function moveRP(rubik)
{
    /* WITH TEST */
    // Rotate R Side in reverse
    var movedRubik = copyRubik(rubik);
    var posToChange = ['R1','R2','R3','R6','R9','R8','R7','R4','B3','B6','B9','D7','D4','D1','F3','F6','F9','U3','U6','U9'];
    var posNewVal = ['R3','R6','R9','R8','R7','R4','R1','R2','D7','D4','D1','F3','F6','F9','U3','U6','U9','B3','B6','B9'];
    movedRubik = rubikMove(movedRubik, posToChange, posNewVal);
    return movedRubik;
}
function moveL(rubik)
{
    /* WITH TEST */
    // Rotate L Side
    var posToChange = ['L1','L2','L3','L6','L9','L7','L8','L4','B1','B4','B7','U1','U4','U7','F1','F4','F7','D3','D6','D9'];
    var posNewVal = ['L7','L4','L1','L2','L3','L9','L6','L8','D9','D6','D3','B1','B4','B7','U1','U4','U7','F7','F4','F1'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveLP(rubik)
{
    /* WITH TEST */
    // Rotate L Side in reverse
    var posToChange = ['L1','L2','L3','L6','L9','L8','L7','L4','B1','B4','B7','U1','U4','U7','F1','F4','F7','D3','D6','D9'];
    var posNewVal = ['L3','L6','L9','L8','L7','L4','L1','L2','U1','U4','U7','F1','F4','F7','D9','D6','D3','B7','B4','B1'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveF(rubik)
{
    /* WITH TEST */
    // Rotate F Side
    var posToChange = ['F1','F2','F3','F6','F9','F7','F8','F4','L7','L8','L9','U7','U8','U9','R7','R8','R9','D7','D8','D9'];
    var posNewVal = ['F7','F4','F1','F2','F3','F9','F6','F8','D7','D8','D9','L7','L8','L9','U7','U8','U9','R7','R8','R9'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveFP(rubik)
{
    /* WITH TEST */
    // Rotate F Side in reverse
    var posToChange = ['F1','F2','F3','F6','F9','F8','F7','F4','L7','L8','L9','U7','U8','U9','R7','R8','R9','D7','D8','D9'];
    var posNewVal = ['F3','F6','F9','F8','F7','F4','F1','F2','U7','U8','U9','R7','R8','R9','D7','D8','D9','L7','L8','L9'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveB(rubik)
{
    /* WITH TEST */
    // Rotate B Side
    var posToChange = ['B1','B2','B3','B6','B9','B8','B7','B4','L1','L2','L3','U1','U2','U3','R1','R2','R3','D1','D2','D3'];
    var posNewVal = ['B7','B4','B1','B2','B3','B6','B9','B8','U1','U2','U3','R1','R2','R3','D1','D2','D3','L1','L2','L3'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveBP(rubik)
{
    /* WITH TEST */
    // Rotate B Side in reverse
    var posToChange = ['B1','B2','B3','B6','B9','B8','B7','B4','L1','L2','L3','U1','U2','U3','R1','R2','R3','D1','D2','D3'];
    var posNewVal = ['B3','B6','B9','B8','B7','B4','B1','B2','D1','D2','D3','L1','L2','L3','U1','U2','U3','R1','R2','R3'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveD(rubik)
{
    /* WITH TEST */
    // Rotate D Side
    var posToChange = ['D1','D2','D3','D6','D9','D8','D7','D4','B1','B2','B3','L1','L4','L7','R3','R6','R9','F7','F8','F9'];
    var posNewVal = ['D7','D4','D1','D2','D3','D6','D9','D8','R3','R6','R9','B3','B2','B1','F9','F8','F7','L1','L4','L7'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function moveDP(rubik)
{
    /* WITH TEST */
    // Rotate D Side in reverse
    var posToChange = ['D1','D2','D3','D6','D9','D8','D7','D4','B1','B2','B3','L1','L4','L7','R3','R6','R9','F7','F8','F9'];
    var posNewVal = ['D3','D6','D9','D8','D7','D4','D1','D2','L7','L4','L1','F7','F8','F9','B1','B2','B3','R9','R6','R3'];
    rubik = rubikMove(rubik, posToChange, posNewVal);
    return rubik;
}
function sexyMove(sexyRubik)
{
    /* WITH TEST */
    sexyRubik = moveR(sexyRubik);
    sexyRubik = moveU(sexyRubik);
    sexyRubik = moveRP(sexyRubik);
    sexyRubik = moveUP(sexyRubik);
    return sexyRubik;
}
function printPos(rubik, pos)
{
    //global red, green, yellow, blue, magenta, white;
    return rubik[1][pos]+rubik[0][pos]+white;
}
function formatCli(rubik)
{
    console.log("         |--------|"+"\n" +
                "         |"+printPos(rubik, 'B1')+"-"+printPos(rubik, 'B2')+"-"+printPos(rubik, 'B3')+"|"+"\n" +
                "         |"+printPos(rubik, 'B4')+"-"+printPos(rubik, 'B5')+"-"+printPos(rubik, 'B6')+"|"+"\n" +
                "         |"+printPos(rubik, 'B7')+"-"+printPos(rubik, 'B8')+"-"+printPos(rubik, 'B9')+"|"+"\n" +
                "|--------|--------|--------|--------|"+"\n" +
                "|"+printPos(rubik, 'L1')+"-"+printPos(rubik, 'L2')+"-"+printPos(rubik, 'L3')+"|" +      
                printPos(rubik, 'U1')+"-"+printPos(rubik, 'U2')+"-"+printPos(rubik, 'U3') +      
                "|"+printPos(rubik, 'R1')+"-"+printPos(rubik, 'R2')+"-"+printPos(rubik, 'R3')+"|" +      
                printPos(rubik, 'D1')+"-"+printPos(rubik, 'D2')+"-"+printPos(rubik, 'D3')+"|"+"\n" +      
                "|"+printPos(rubik, 'L4')+"-"+printPos(rubik, 'L5')+"-"+printPos(rubik, 'L6')+"|" +      
                printPos(rubik, 'U4')+"-"+printPos(rubik, 'U5')+"-"+printPos(rubik, 'U6') +      
                "|"+printPos(rubik, 'R4')+"-"+printPos(rubik, 'R5')+"-"+printPos(rubik, 'R6')+"|" +      
                printPos(rubik, 'D4')+"-"+printPos(rubik, 'D5')+"-"+printPos(rubik, 'D6')+"|"+"\n" +      
                "|"+printPos(rubik, 'L7')+"-"+printPos(rubik, 'L8')+"-"+printPos(rubik, 'L9')+"|" +      
                printPos(rubik, 'U7')+"-"+printPos(rubik, 'U8')+"-"+printPos(rubik, 'U9') +      
                "|"+printPos(rubik, 'R7')+"-"+printPos(rubik, 'R8')+"-"+printPos(rubik, 'R9')+"|" +      
                printPos(rubik, 'D7')+"-"+printPos(rubik, 'D8')+"-"+printPos(rubik, 'D9')+"|"+"\n" +      
                "|--------|--------|--------|--------|"+"\n" +      
                "         |"+printPos(rubik, 'F1')+"-"+printPos(rubik, 'F2')+"-"+printPos(rubik, 'F3')+"|"+"\n" +      
                "         |"+printPos(rubik, 'F4')+"-"+printPos(rubik, 'F5')+"-"+printPos(rubik, 'F6')+"|"+"\n" +      
                "         |"+printPos(rubik, 'F7')+"-"+printPos(rubik, 'F8')+"-"+printPos(rubik, 'F9')+"|"+"\n" +      
                "         |--------|"+"\n");
}
function printPosHtml(rubik, pos)
{
    //global red, green, yellow, blue, magenta, white;
    return '<span style="background-color: '+translateColorToHtml(rubik[pos]['c'])+'">'+rubik[pos]['v']+'</span>';
}
function formatHtml(rubik)
{
    console.log("<table>");
    console.log(" <tr><td></td><td>"+printPosHtml(rubik, 'B1')+"-"+printPosHtml(rubik, 'B2')+"-"+printPosHtml(rubik, 'B3')+"</td></tr>");
    console.log(" <tr><td></td><td>"+printPosHtml(rubik, 'B4')+"-"+printPosHtml(rubik, 'B5')+"-"+printPosHtml(rubik, 'B6')+"</td></tr>");
    console.log(" <tr><td></td><td>"+printPosHtml(rubik, 'B7')+"-"+printPosHtml(rubik, 'B8')+"-"+printPosHtml(rubik, 'B9')+"</td></tr>");
    console.log(" <tr><td>"+printPosHtml(rubik, 'L1')+"-"+printPosHtml(rubik, 'L2')+"-"+printPosHtml(rubik, 'L3')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'U1')+"-"+printPosHtml(rubik, 'U2')+"-"+printPosHtml(rubik, 'U3')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'R1')+"-"+printPosHtml(rubik, 'R2')+"-"+printPosHtml(rubik, 'R3')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'D1')+"-"+printPosHtml(rubik, 'D2')+"-"+printPosHtml(rubik, 'D3')+"</td></tr>");
    console.log(" <tr><td>"+printPosHtml(rubik, 'L4')+"-"+printPosHtml(rubik, 'L5')+"-"+printPosHtml(rubik, 'L6')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'U4')+"-"+printPosHtml(rubik, 'U5')+"-"+printPosHtml(rubik, 'U6')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'R4')+"-"+printPosHtml(rubik, 'R5')+"-"+printPosHtml(rubik, 'R6')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'D4')+"-"+printPosHtml(rubik, 'D5')+"-"+printPosHtml(rubik, 'D6')+"</td></tr>");
    console.log(" <tr><td>"+printPosHtml(rubik, 'L7')+"-"+printPosHtml(rubik, 'L8')+"-"+printPosHtml(rubik, 'L9')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'U7')+"-"+printPosHtml(rubik, 'U8')+"-"+printPosHtml(rubik, 'U9')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'R7')+"-"+printPosHtml(rubik, 'R8')+"-"+printPosHtml(rubik, 'R9')+"</td>");
    console.log("     <td>"+printPosHtml(rubik, 'D7')+"-"+printPosHtml(rubik, 'D8')+"-"+printPosHtml(rubik, 'D9')+"</td></tr>");
    console.log(" <tr><td></td><td>"+printPosHtml(rubik, 'F1')+"-"+printPosHtml(rubik, 'F2')+"-"+printPosHtml(rubik, 'F3')+"</td></tr>");
    console.log(" <tr><td></td><td>"+printPosHtml(rubik, 'F4')+"-"+printPosHtml(rubik, 'F5')+"-"+printPosHtml(rubik, 'F6')+"</td></tr>");
    console.log(" <tr><td></td><td>"+printPosHtml(rubik, 'F7')+"-"+printPosHtml(rubik, 'F8')+"-"+printPosHtml(rubik, 'F9')+"</td></tr>");
    console.log("</table>");
}
function translateColorToHtml(color)
{
    //global red, green, yellow, blue, magenta, white;
    switch (color) {
        case red:
            return "red";
            break;
        case green:
            return "lightgreen";
            break;
        case yellow:
            return "yellow";
            break;
        case blue:
            return "lightblue";
            break;
        case magenta:
            return "orange";
            break;
        case white:
            return "white";
            break;
    }
}
var rubik = getOrderedRubik();

//for (var k=0; k<6; k++) {
//  rubik=sexyMove(rubik);
//}
//formatCli(rubik);

// Filter get var
// movesParam = filter_input(1, 'm', 513);
// // If moves are setted via string query
// if (typeof movesParam !== 'undefined') {
//     moves = str_split(movesParam);
//     moves.foreach(move => {
//         switch (move) {
//             case 'U':
//                 rubik = moveU(rubik);
//                 break;
//             case 'u':
//                 rubik = moveUP(rubik);
//                 break;
//             case 'R':
//                 rubik = moveR(rubik);
//                 break;
//             case 'r':
//                 rubik = moveRP(rubik);
//                 break;
//             case 'L':
//                 rubik = moveL(rubik);
//                 break;
//             case 'l':
//                 rubik = moveLP(rubik);
//                 break;
//             case 'F':
//                 rubik = moveF(rubik);
//                 break;
//             case 'f':
//                 rubik = moveFP(rubik);
//                 break;
//             case 'B':
//                 rubik = moveB(rubik);
//                 break;
//             case 'b':
//                 rubik = moveBP(rubik);
//                 break;
//             case 'D':
//                 rubik = moveD(rubik);
//                 break;
//             case 'd':
//                 rubik = moveDP(rubik);
//                 break;
//         }
//     })
// }
//rubik=moveDP(rubik);
var rubik2 = copyRubik(rubik);
rubik2 = moveR(rubik2);
var rubik3 = copyRubik(rubik2)
rubik3 = moveU(rubik3);
//rubik = moveRP(rubik);
//rubik = moveUP(rubik);

//if (typeof argv[0] !== 'undefined') {
    // if command line cli format is displayed
    formatCli(rubik3);
//} else {
    // else html format is displayed
//    formatHtml(rubik);
//}
// var_dump($rubik);