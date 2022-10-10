namespace Namespace {
    
    using System.Collections.Generic;
    
    using System;
    
    using System.Linq;
    
    public static class Module {
        
        public static object red = "\x1b[31m";
        
        public static object green = "\x1b[32m";
        
        public static object yellow = "\x1b[33m";
        
        public static object blue = "\x1b[34m";
        
        public static object magenta = "\x1b[95m";
        
        public static object white = "\x1b[97m";
        
        // WITH TEST
        public static object rubik() {
            Func<object> getOrderedRubik = ()
                //global red, green, yellow, blue, magenta, white
                var rubikValues = new Dictionary<object, object> {
                };
                var rubikColors = new Dictionary<object, object> {
                };
                var sides = new List<object> {
                    "U",
                    "B",
                    "R",
                    "D",
                    "L",
                    "F"
                };
                foreach (var side in sides) {
                    foreach (var i in Enumerable.Range(1, 10 - 1)) {
                        var index = side + i.ToString();
                        rubikValues[index] = index;

                        // TODO: Falta codigo aqui
                    }
                }
            };
            Func<object, object, object, object> rubikMove = (rubik,posA,posN) => {
                var tempRubik = copy.deepcopy(rubik);
                foreach (var i in Enumerable.Range(0, posA.Count - 0)) {
                    var posActual = posA[i];
                    var posNew = posN[i];
                    rubik[0][posActual] = tempRubik[0][posNew];
                    rubik[1][posActual] = tempRubik[1][posNew];
                }
                return rubik;
            };
            Func<object, object> moveU = rubik => {
                // Rotate A Side
                var posToChange = new List<object> {
                    "U1",
                    "U2",
                    "U3",
                    "U6",
                    "U9",
                    "U8",
                    "U7",
                    "U4",
                    "R1",
                    "R4",
                    "R7",
                    "F1",
                    "F2",
                    "F3",
                    "L3",
                    "L6",
                    "L9",
                    "B7",
                    "B8",
                    "B9"
                };
                var posNewVal = new List<object> {
                    "U7",
                    "U4",
                    "U1",
                    "U2",
                    "U3",
                    "U6",
                    "U9",
                    "U8",
                    "B7",
                    "B8",
                    "B9",
                    "R7",
                    "R4",
                    "R1",
                    "F1",
                    "F2",
                    "F3",
                    "L9",
                    "L6",
                    "L3"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveUP = rubik => {
                // Rotate A Side in reverse
                var posToChange = new List<object> {
                    "U1",
                    "U2",
                    "U3",
                    "U6",
                    "U9",
                    "U8",
                    "U7",
                    "U4",
                    "R1",
                    "R4",
                    "R7",
                    "F1",
                    "F2",
                    "F3",
                    "L3",
                    "L6",
                    "L9",
                    "B7",
                    "B8",
                    "B9"
                };
                var posNewVal = new List<object> {
                    "U3",
                    "U6",
                    "U9",
                    "U8",
                    "U7",
                    "U4",
                    "U1",
                    "U2",
                    "F3",
                    "F2",
                    "F1",
                    "L3",
                    "L6",
                    "L9",
                    "B9",
                    "B8",
                    "B7",
                    "R1",
                    "R4",
                    "R7"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveR = rubik => {
                // Rotate R Side
                var posToChange = new List<object> {
                    "R1",
                    "R2",
                    "R3",
                    "R6",
                    "R9",
                    "R7",
                    "R8",
                    "R4",
                    "B3",
                    "B6",
                    "B9",
                    "D7",
                    "D4",
                    "D1",
                    "F3",
                    "F6",
                    "F9",
                    "U3",
                    "U6",
                    "U9"
                };
                var posNewVal = new List<object> {
                    "R7",
                    "R4",
                    "R1",
                    "R2",
                    "R3",
                    "R9",
                    "R6",
                    "R8",
                    "U3",
                    "U6",
                    "U9",
                    "B3",
                    "B6",
                    "B9",
                    "D7",
                    "D4",
                    "D1",
                    "F3",
                    "F6",
                    "F9"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveRP = rubik => {
                // Rotate R Side in reverse
                var posToChange = new List<object> {
                    "R1",
                    "R2",
                    "R3",
                    "R6",
                    "R9",
                    "R8",
                    "R7",
                    "R4",
                    "B3",
                    "B6",
                    "B9",
                    "D7",
                    "D4",
                    "D1",
                    "F3",
                    "F6",
                    "F9",
                    "U3",
                    "U6",
                    "U9"
                };
                var posNewVal = new List<object> {
                    "R3",
                    "R6",
                    "R9",
                    "R8",
                    "R7",
                    "R4",
                    "R1",
                    "R2",
                    "D7",
                    "D4",
                    "D1",
                    "F3",
                    "F6",
                    "F9",
                    "U3",
                    "U6",
                    "U9",
                    "B3",
                    "B6",
                    "B9"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveL = rubik => {
                // Rotate L Side
                var posToChange = new List<object> {
                    "L1",
                    "L2",
                    "L3",
                    "L6",
                    "L9",
                    "L7",
                    "L8",
                    "L4",
                    "B1",
                    "B4",
                    "B7",
                    "U1",
                    "U4",
                    "U7",
                    "F1",
                    "F4",
                    "F7",
                    "D3",
                    "D6",
                    "D9"
                };
                var posNewVal = new List<object> {
                    "L7",
                    "L4",
                    "L1",
                    "L2",
                    "L3",
                    "L9",
                    "L6",
                    "L8",
                    "D9",
                    "D6",
                    "D3",
                    "B1",
                    "B4",
                    "B7",
                    "U1",
                    "U4",
                    "U7",
                    "F7",
                    "F4",
                    "F1"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveLP = rubik => {
                // Rotate L Side in reverse
                var posToChange = new List<object> {
                    "L1",
                    "L2",
                    "L3",
                    "L6",
                    "L9",
                    "L8",
                    "L7",
                    "L4",
                    "B1",
                    "B4",
                    "B7",
                    "U1",
                    "U4",
                    "U7",
                    "F1",
                    "F4",
                    "F7",
                    "D3",
                    "D6",
                    "D9"
                };
                var posNewVal = new List<object> {
                    "L3",
                    "L6",
                    "L9",
                    "L8",
                    "L7",
                    "L4",
                    "L1",
                    "L2",
                    "U1",
                    "U4",
                    "U7",
                    "F1",
                    "F4",
                    "F7",
                    "D9",
                    "D6",
                    "D3",
                    "B7",
                    "B4",
                    "B1"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveF = rubik => {
                // Rotate F Side
                var posToChange = new List<object> {
                    "F1",
                    "F2",
                    "F3",
                    "F6",
                    "F9",
                    "F7",
                    "F8",
                    "F4",
                    "L7",
                    "L8",
                    "L9",
                    "U7",
                    "U8",
                    "U9",
                    "R7",
                    "R8",
                    "R9",
                    "D7",
                    "D8",
                    "D9"
                };
                var posNewVal = new List<object> {
                    "F7",
                    "F4",
                    "F1",
                    "F2",
                    "F3",
                    "F9",
                    "F6",
                    "F8",
                    "D7",
                    "D8",
                    "D9",
                    "L7",
                    "L8",
                    "L9",
                    "U7",
                    "U8",
                    "U9",
                    "R7",
                    "R8",
                    "R9"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveFP = rubik => {
                // Rotate F Side in reverse
                var posToChange = new List<object> {
                    "F1",
                    "F2",
                    "F3",
                    "F6",
                    "F9",
                    "F8",
                    "F7",
                    "F4",
                    "L7",
                    "L8",
                    "L9",
                    "U7",
                    "U8",
                    "U9",
                    "R7",
                    "R8",
                    "R9",
                    "D7",
                    "D8",
                    "D9"
                };
                var posNewVal = new List<object> {
                    "F3",
                    "F6",
                    "F9",
                    "F8",
                    "F7",
                    "F4",
                    "F1",
                    "F2",
                    "U7",
                    "U8",
                    "U9",
                    "R7",
                    "R8",
                    "R9",
                    "D7",
                    "D8",
                    "D9",
                    "L7",
                    "L8",
                    "L9"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveB = rubik => {
                // Rotate B Side
                var posToChange = new List<object> {
                    "B1",
                    "B2",
                    "B3",
                    "B6",
                    "B9",
                    "B8",
                    "B7",
                    "B4",
                    "L1",
                    "L2",
                    "L3",
                    "U1",
                    "U2",
                    "U3",
                    "R1",
                    "R2",
                    "R3",
                    "D1",
                    "D2",
                    "D3"
                };
                var posNewVal = new List<object> {
                    "B7",
                    "B4",
                    "B1",
                    "B2",
                    "B3",
                    "B6",
                    "B9",
                    "B8",
                    "U1",
                    "U2",
                    "U3",
                    "R1",
                    "R2",
                    "R3",
                    "D1",
                    "D2",
                    "D3",
                    "L1",
                    "L2",
                    "L3"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveBP = rubik => {
                // Rotate B Side in reverse
                var posToChange = new List<object> {
                    "B1",
                    "B2",
                    "B3",
                    "B6",
                    "B9",
                    "B8",
                    "B7",
                    "B4",
                    "L1",
                    "L2",
                    "L3",
                    "U1",
                    "U2",
                    "U3",
                    "R1",
                    "R2",
                    "R3",
                    "D1",
                    "D2",
                    "D3"
                };
                var posNewVal = new List<object> {
                    "B3",
                    "B6",
                    "B9",
                    "B8",
                    "B7",
                    "B4",
                    "B1",
                    "B2",
                    "D1",
                    "D2",
                    "D3",
                    "L1",
                    "L2",
                    "L3",
                    "U1",
                    "U2",
                    "U3",
                    "R1",
                    "R2",
                    "R3"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveD = rubik => {
                // Rotate D Side
                var posToChange = new List<object> {
                    "D1",
                    "D2",
                    "D3",
                    "D6",
                    "D9",
                    "D8",
                    "D7",
                    "D4",
                    "B1",
                    "B2",
                    "B3",
                    "L1",
                    "L4",
                    "L7",
                    "R3",
                    "R6",
                    "R9",
                    "F7",
                    "F8",
                    "F9"
                };
                var posNewVal = new List<object> {
                    "D7",
                    "D4",
                    "D1",
                    "D2",
                    "D3",
                    "D6",
                    "D9",
                    "D8",
                    "R3",
                    "R6",
                    "R9",
                    "B3",
                    "B2",
                    "B1",
                    "F9",
                    "F8",
                    "F7",
                    "L1",
                    "L4",
                    "L7"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> moveDP = rubik => {
                // Rotate D Side in reverse
                var posToChange = new List<object> {
                    "D1",
                    "D2",
                    "D3",
                    "D6",
                    "D9",
                    "D8",
                    "D7",
                    "D4",
                    "B1",
                    "B2",
                    "B3",
                    "L1",
                    "L4",
                    "L7",
                    "R3",
                    "R6",
                    "R9",
                    "F7",
                    "F8",
                    "F9"
                };
                var posNewVal = new List<object> {
                    "D3",
                    "D6",
                    "D9",
                    "D8",
                    "D7",
                    "D4",
                    "D1",
                    "D2",
                    "L7",
                    "L4",
                    "L1",
                    "F7",
                    "F8",
                    "F9",
                    "B1",
                    "B2",
                    "B3",
                    "R9",
                    "R6",
                    "R3"
                };
                rubik = rubikMove(rubik, posToChange, posNewVal);
                return rubik;
            };
            Func<object, object> sexyMove = rubik => {
                rubik = moveR(rubik);
                rubik = moveU(rubik);
                rubik = moveRP(rubik);
                rubik = moveUP(rubik);
                return rubik;
            };
            Func<object, object, object> printPos = (rubik,pos) => {
                return rubik[1][pos] + rubik[0][pos] + white;
            };
            Func<object, object> formatCli = rubik => {
                Console.WriteLine("         |--------|" + "\n" + "         |" + printPos(rubik, "B1") + "-" + printPos(rubik, "B2") + "-" + printPos(rubik, "B3") + "|" + "\n" + "         |" + printPos(rubik, "B4") + "-" + printPos(rubik, "B5") + "-" + printPos(rubik, "B6") + "|" + "\n" + "         |" + printPos(rubik, "B7") + "-" + printPos(rubik, "B8") + "-" + printPos(rubik, "B9") + "|" + "\n" + "|--------|--------|--------|--------|" + "\n" + "|" + printPos(rubik, "L1") + "-" + printPos(rubik, "L2") + "-" + printPos(rubik, "L3") + "|" + printPos(rubik, "U1") + "-" + printPos(rubik, "U2") + "-" + printPos(rubik, "U3") + "|" + printPos(rubik, "R1") + "-" + printPos(rubik, "R2") + "-" + printPos(rubik, "R3") + "|" + printPos(rubik, "D1") + "-" + printPos(rubik, "D2") + "-" + printPos(rubik, "D3") + "|" + "\n" + "|" + printPos(rubik, "L4") + "-" + printPos(rubik, "L5") + "-" + printPos(rubik, "L6") + "|" + printPos(rubik, "U4") + "-" + printPos(rubik, "U5") + "-" + printPos(rubik, "U6") + "|" + printPos(rubik, "R4") + "-" + printPos(rubik, "R5") + "-" + printPos(rubik, "R6") + "|" + printPos(rubik, "D4") + "-" + printPos(rubik, "D5") + "-" + printPos(rubik, "D6") + "|" + "\n" + "|" + printPos(rubik, "L7") + "-" + printPos(rubik, "L8") + "-" + printPos(rubik, "L9") + "|" + printPos(rubik, "U7") + "-" + printPos(rubik, "U8") + "-" + printPos(rubik, "U9") + "|" + printPos(rubik, "R7") + "-" + printPos(rubik, "R8") + "-" + printPos(rubik, "R9") + "|" + printPos(rubik, "D7") + "-" + printPos(rubik, "D8") + "-" + printPos(rubik, "D9") + "|" + "\n" + "|--------|--------|--------|--------|" + "\n" + "         |" + printPos(rubik, "F1") + "-" + printPos(rubik, "F2") + "-" + printPos(rubik, "F3") + "|" + "\n" + "         |" + printPos(rubik, "F4") + "-" + printPos(rubik, "F5") + "-" + printPos(rubik, "F6") + "|" + "\n" + "         |" + printPos(rubik, "F7") + "-" + printPos(rubik, "F8") + "-" + printPos(rubik, "F9") + "|" + "\n" + "         |--------|" + "\n");
            };
            Func<object, object, object> printPosHtml = (rubik,pos) => {
                return "<span style=\"background-color: " + translateColorToHtml(rubik[pos]["c"]) + "\">" + rubik[pos]["v"] + "</span>";
            };
            Func<object, object> formatHtml = rubik => {
                Console.WriteLine("<table>");
                Console.WriteLine(" <tr><td></td><td>" + printPosHtml(rubik, "B1") + "-" + printPosHtml(rubik, "B2") + "-" + printPosHtml(rubik, "B3") + "</td></tr>");
                Console.WriteLine(" <tr><td></td><td>" + printPosHtml(rubik, "B4") + "-" + printPosHtml(rubik, "B5") + "-" + printPosHtml(rubik, "B6") + "</td></tr>");
                Console.WriteLine(" <tr><td></td><td>" + printPosHtml(rubik, "B7") + "-" + printPosHtml(rubik, "B8") + "-" + printPosHtml(rubik, "B9") + "</td></tr>");
                Console.WriteLine(" <tr><td>" + printPosHtml(rubik, "L1") + "-" + printPosHtml(rubik, "L2") + "-" + printPosHtml(rubik, "L3") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "U1") + "-" + printPosHtml(rubik, "U2") + "-" + printPosHtml(rubik, "U3") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "R1") + "-" + printPosHtml(rubik, "R2") + "-" + printPosHtml(rubik, "R3") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "D1") + "-" + printPosHtml(rubik, "D2") + "-" + printPosHtml(rubik, "D3") + "</td></tr>");
                Console.WriteLine(" <tr><td>" + printPosHtml(rubik, "L4") + "-" + printPosHtml(rubik, "L5") + "-" + printPosHtml(rubik, "L6") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "U4") + "-" + printPosHtml(rubik, "U5") + "-" + printPosHtml(rubik, "U6") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "R4") + "-" + printPosHtml(rubik, "R5") + "-" + printPosHtml(rubik, "R6") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "D4") + "-" + printPosHtml(rubik, "D5") + "-" + printPosHtml(rubik, "D6") + "</td></tr>");
                Console.WriteLine(" <tr><td>" + printPosHtml(rubik, "L7") + "-" + printPosHtml(rubik, "L8") + "-" + printPosHtml(rubik, "L9") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "U7") + "-" + printPosHtml(rubik, "U8") + "-" + printPosHtml(rubik, "U9") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "R7") + "-" + printPosHtml(rubik, "R8") + "-" + printPosHtml(rubik, "R9") + "</td>");
                Console.WriteLine("     <td>" + printPosHtml(rubik, "D7") + "-" + printPosHtml(rubik, "D8") + "-" + printPosHtml(rubik, "D9") + "</td></tr>");
                Console.WriteLine(" <tr><td></td><td>" + printPosHtml(rubik, "F1") + "-" + printPosHtml(rubik, "F2") + "-" + printPosHtml(rubik, "F3") + "</td></tr>");
                Console.WriteLine(" <tr><td></td><td>" + printPosHtml(rubik, "F4") + "-" + printPosHtml(rubik, "F5") + "-" + printPosHtml(rubik, "F6") + "</td></tr>");
                Console.WriteLine(" <tr><td></td><td>" + printPosHtml(rubik, "F7") + "-" + printPosHtml(rubik, "F8") + "-" + printPosHtml(rubik, "F9") + "</td></tr>");
                Console.WriteLine("</table>");
            };
            Func<object> main = () => {
                var rubik = getOrderedRubik();
                // print(rubik)
                // for j in range (0,6):
                //    rubik=sexyMove(rubik)
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
                //         
                //     )
                // 
                rubik = moveDP(rubik);
                //if typeof argv[0] is not 'undefined') {
                // if command line cli format is displayed
                formatCli(rubik);
                // else {
                // else html format is displayed
                //    formatHtml(rubik);
                // var_dump($rubik);
            };
        }
    }
}