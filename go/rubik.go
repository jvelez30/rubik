// Code generated. MUST EDIT!



red = "\33[31m"
green = "\33[32m"
yellow = "\33[33m"
blue = "\33[34m"
magenta = "\33[95m"
white = "\33[97m"

func getOrderedRubik() {
    /* WITH TEST */
    global red, green, yellow, blue, magenta, white; /* cannot convert */
    rubik = []interface{}{
,
    }
    sides = []interface{}{
        "U",
        "B",
        "R",
        "D",
        "L",
        "F",
    }
    for _, side := range sides {
        for i = 1; i < 10; i++ {
            rubik[side . i]["v"] = side . i
            switch (side) {
                case "U":
                    rubik[side . i]["c"] = yellow
                    break;
                case "B":
                    rubik[side . i]["c"] = magenta
                    break;
                case "R":
                    rubik[side . i]["c"] = green
                    break;
                case "D":
                    rubik[side . i]["c"] = white
                    break;
                case "L":
                    rubik[side . i]["c"] = blue
                    break;
                case "F":
                    rubik[side . i]["c"] = red
                    break;
            }
        }
    }
    return rubik
}
func rubikMove(rubik, posA, posN) {
    tempRubik = rubik
    for i = 0; i < len(posA); i++ {
        posActual = posA[i]
        posNew = posN[i]
        rubik[posActual]["v"] = tempRubik[posNew]["v"]
        rubik[posActual]["c"] = tempRubik[posNew]["c"]
    }
    return rubik
}
func moveU(rubik) {
    /* WITH TEST */
    // Rotate A Side
    posToChange = []interface{}{
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
        "B9",
    }
    posNewVal = []interface{}{
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
        "L3",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveUP(rubik) {
    /* WITH TEST */
    // Rotate A Side in reverse
    posToChange = []interface{}{
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
        "B9",
    }
    posNewVal = []interface{}{
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
        "R7",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveR(rubik) {
    /* WITH TEST */
    // Rotate R Side
    posToChange = []interface{}{
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
        "U9",
    }
    posNewVal = []interface{}{
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
        "F9",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveRP(rubik) {
    /* WITH TEST */
    // Rotate R Side in reverse
    posToChange = []interface{}{
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
        "U9",
    }
    posNewVal = []interface{}{
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
        "B9",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveL(rubik) {
    /* WITH TEST */
    // Rotate L Side
    posToChange = []interface{}{
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
        "D9",
    }
    posNewVal = []interface{}{
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
        "F1",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveLP(rubik) {
    /* WITH TEST */
    // Rotate L Side in reverse
    posToChange = []interface{}{
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
        "D9",
    }
    posNewVal = []interface{}{
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
        "B1",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveF(rubik) {
    /* WITH TEST */
    // Rotate F Side
    posToChange = []interface{}{
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
        "D9",
    }
    posNewVal = []interface{}{
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
        "R9",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveFP(rubik) {
    /* WITH TEST */
    // Rotate F Side in reverse
    posToChange = []interface{}{
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
        "D9",
    }
    posNewVal = []interface{}{
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
        "L9",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveB(rubik) {
    /* WITH TEST */
    // Rotate B Side
    posToChange = []interface{}{
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
        "D3",
    }
    posNewVal = []interface{}{
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
        "L3",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveBP(rubik) {
    /* WITH TEST */
    // Rotate B Side in reverse
    posToChange = []interface{}{
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
        "D3",
    }
    posNewVal = []interface{}{
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
        "R3",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveD(rubik) {
    /* WITH TEST */
    // Rotate D Side
    posToChange = []interface{}{
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
        "F9",
    }
    posNewVal = []interface{}{
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
        "L7",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func moveDP(rubik) {
    /* WITH TEST */
    // Rotate D Side in reverse
    posToChange = []interface{}{
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
        "F9",
    }
    posNewVal = []interface{}{
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
        "R3",
    }
    rubik = rubikMove(rubik, posToChange, posNewVal)
    return rubik
}
func sexyMove(rubik) {
    /* WITH TEST */
    rubik = moveR(rubik)
    rubik = moveU(rubik)
    rubik = moveRP(rubik)
    rubik = moveUP(rubik)
    return rubik
}
func printPos(rubik, pos) {
    global red, green, yellow, blue, magenta, white; /* cannot convert */
    return rubik[pos]["c"] . rubik[pos]["v"] . white
}
func formatCli(rubik) {
    fmt.Print("         |--------|" . "\n")
    fmt.Print("         |" . printPos(rubik, "B1") . "-" . printPos(rubik, "B2") . "-" . printPos(rubik, "B3") . "|" . "\n")
    fmt.Print("         |" . printPos(rubik, "B4") . "-" . printPos(rubik, "B5") . "-" . printPos(rubik, "B6") . "|" . "\n")
    fmt.Print("         |" . printPos(rubik, "B7") . "-" . printPos(rubik, "B8") . "-" . printPos(rubik, "B9") . "|" . "\n")
    fmt.Print("|--------|--------|--------|--------|" . "\n")
    fmt.Print("|" . printPos(rubik, "L1") . "-" . printPos(rubik, "L2") . "-" . printPos(rubik, "L3") . "|")
    fmt.Print(printPos(rubik, "U1") . "-" . printPos(rubik, "U2") . "-" . printPos(rubik, "U3"))
    fmt.Print("|" . printPos(rubik, "R1") . "-" . printPos(rubik, "R2") . "-" . printPos(rubik, "R3") . "|")
    fmt.Print(printPos(rubik, "D1") . "-" . printPos(rubik, "D2") . "-" . printPos(rubik, "D3") . "|" . "\n")
    fmt.Print("|" . printPos(rubik, "L4") . "-" . printPos(rubik, "L5") . "-" . printPos(rubik, "L6") . "|")
    fmt.Print(printPos(rubik, "U4") . "-" . printPos(rubik, "U5") . "-" . printPos(rubik, "U6"))
    fmt.Print("|" . printPos(rubik, "R4") . "-" . printPos(rubik, "R5") . "-" . printPos(rubik, "R6") . "|")
    fmt.Print(printPos(rubik, "D4") . "-" . printPos(rubik, "D5") . "-" . printPos(rubik, "D6") . "|" . "\n")
    fmt.Print("|" . printPos(rubik, "L7") . "-" . printPos(rubik, "L8") . "-" . printPos(rubik, "L9") . "|")
    fmt.Print(printPos(rubik, "U7") . "-" . printPos(rubik, "U8") . "-" . printPos(rubik, "U9"))
    fmt.Print("|" . printPos(rubik, "R7") . "-" . printPos(rubik, "R8") . "-" . printPos(rubik, "R9") . "|")
    fmt.Print(printPos(rubik, "D7") . "-" . printPos(rubik, "D8") . "-" . printPos(rubik, "D9") . "|" . "\n")
    fmt.Print("|--------|--------|--------|--------|" . "\n")
    fmt.Print("         |" . printPos(rubik, "F1") . "-" . printPos(rubik, "F2") . "-" . printPos(rubik, "F3") . "|" . "\n")
    fmt.Print("         |" . printPos(rubik, "F4") . "-" . printPos(rubik, "F5") . "-" . printPos(rubik, "F6") . "|" . "\n")
    fmt.Print("         |" . printPos(rubik, "F7") . "-" . printPos(rubik, "F8") . "-" . printPos(rubik, "F9") . "|" . "\n")
    fmt.Print("         |--------|" . "\n")
}
func printPosHtml(rubik, pos) {
    global red, green, yellow, blue, magenta, white; /* cannot convert */
    return "<span style=\"background-color: " . translateColorToHtml(rubik[pos]["c"]) . "\">" . rubik[pos]["v"] . "</span>"
}
func formatHtml(rubik) {
    fmt.Print("<table>")
    fmt.Print(" <tr><td></td><td>" . printPosHtml(rubik, "B1") . "-" . printPosHtml(rubik, "B2") . "-" . printPosHtml(rubik, "B3") . "</td></tr>")
    fmt.Print(" <tr><td></td><td>" . printPosHtml(rubik, "B4") . "-" . printPosHtml(rubik, "B5") . "-" . printPosHtml(rubik, "B6") . "</td></tr>")
    fmt.Print(" <tr><td></td><td>" . printPosHtml(rubik, "B7") . "-" . printPosHtml(rubik, "B8") . "-" . printPosHtml(rubik, "B9") . "</td></tr>")
    fmt.Print(" <tr><td>" . printPosHtml(rubik, "L1") . "-" . printPosHtml(rubik, "L2") . "-" . printPosHtml(rubik, "L3") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "U1") . "-" . printPosHtml(rubik, "U2") . "-" . printPosHtml(rubik, "U3") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "R1") . "-" . printPosHtml(rubik, "R2") . "-" . printPosHtml(rubik, "R3") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "D1") . "-" . printPosHtml(rubik, "D2") . "-" . printPosHtml(rubik, "D3") . "</td></tr>")
    fmt.Print(" <tr><td>" . printPosHtml(rubik, "L4") . "-" . printPosHtml(rubik, "L5") . "-" . printPosHtml(rubik, "L6") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "U4") . "-" . printPosHtml(rubik, "U5") . "-" . printPosHtml(rubik, "U6") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "R4") . "-" . printPosHtml(rubik, "R5") . "-" . printPosHtml(rubik, "R6") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "D4") . "-" . printPosHtml(rubik, "D5") . "-" . printPosHtml(rubik, "D6") . "</td></tr>")
    fmt.Print(" <tr><td>" . printPosHtml(rubik, "L7") . "-" . printPosHtml(rubik, "L8") . "-" . printPosHtml(rubik, "L9") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "U7") . "-" . printPosHtml(rubik, "U8") . "-" . printPosHtml(rubik, "U9") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "R7") . "-" . printPosHtml(rubik, "R8") . "-" . printPosHtml(rubik, "R9") . "</td>")
    fmt.Print("     <td>" . printPosHtml(rubik, "D7") . "-" . printPosHtml(rubik, "D8") . "-" . printPosHtml(rubik, "D9") . "</td></tr>")
    fmt.Print(" <tr><td></td><td>" . printPosHtml(rubik, "F1") . "-" . printPosHtml(rubik, "F2") . "-" . printPosHtml(rubik, "F3") . "</td></tr>")
    fmt.Print(" <tr><td></td><td>" . printPosHtml(rubik, "F4") . "-" . printPosHtml(rubik, "F5") . "-" . printPosHtml(rubik, "F6") . "</td></tr>")
    fmt.Print(" <tr><td></td><td>" . printPosHtml(rubik, "F7") . "-" . printPosHtml(rubik, "F8") . "-" . printPosHtml(rubik, "F9") . "</td></tr>")
    fmt.Print("</table>")
}
func translateColorToHtml(color) {
    global red, green, yellow, blue, magenta, white; /* cannot convert */
    switch (color) {
        case red:
            return "red"
            break;
        case green:
            return "lightgreen"
            break;
        case yellow:
            return "yellow"
            break;
        case blue:
            return "lightblue"
            break;
        case magenta:
            return "orange"
            break;
        case white:
            return "white"
            break;
    }
}
rubik = getOrderedRubik()

/*
for ($j=0;$j<6;$j++){
  $rubik=sexyMove($rubik);
}
*/
// Filter get var
movesParam = filter_input(INPUT_GET, "m", FILTER_SANITIZE_STRING)
// If moves are setted via string query
if isset(movesParam) /* cannot convert */ {
    moves = str_split(movesParam)
    for _, move := range moves {
        switch (move) {
            case "U":
                rubik = moveU(rubik)
                break;
            case "u":
                rubik = moveUP(rubik)
                break;
            case "R":
                rubik = moveR(rubik)
                break;
            case "r":
                rubik = moveRP(rubik)
                break;
            case "L":
                rubik = moveL(rubik)
                break;
            case "l":
                rubik = moveLP(rubik)
                break;
            case "F":
                rubik = moveF(rubik)
                break;
            case "f":
                rubik = moveFP(rubik)
                break;
            case "B":
                rubik = moveB(rubik)
                break;
            case "b":
                rubik = moveBP(rubik)
                break;
            case "D":
                rubik = moveD(rubik)
                break;
            case "d":
                rubik = moveDP(rubik)
                break;
        }
    }
}

rubik = moveDP(rubik)

if isset(argv[0]) /* cannot convert */ {
    // if command line cli format is displayed
    formatCli(rubik)
} else {
    // else html format is displayed
    formatHtml(rubik)
}

//var_dump($rubik);
