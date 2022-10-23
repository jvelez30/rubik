package main
/* Functional Test */

import "fmt"
import "strconv"

type rubikobj map[string] interface{}

func getOrderedRubik() rubikobj {
    /* WITH TEST */
    rubik := rubikobj{} 
    sides := [6] string {
        "U",
        "B",
        "R",
        "D",
        "L",
        "F",
    }
    for _, side := range sides {
        for i := 1; i < 10; i++ {
            posstring := side + strconv.Itoa(i)
            rubik[posstring + "v"] = posstring
            switch (side) {
                case "U":
                    rubik[posstring + "c"] = "yellow"
                    break;
                case "B":
                    rubik[posstring + "c"] = "magenta"
                    break;
                case "R":
                    rubik[posstring + "c"] = "green"
                    break;
                case "D":
                    rubik[posstring + "c"] = "white"
                    break;
                case "L":
                    rubik[posstring + "c"] = "blue"
                    break;
                case "F":
                    rubik[posstring + "c"] = "red"
                    break;
            }
        }
    }
    return rubik
}
func rubikMove(rubik rubikobj, posA[] string, posN[] string) rubikobj {
    tempRubik := rubikobj{}
    for k, v := range rubik {
        tempRubik[k] = v
    }
    for i := 0; i < len(posA); i++ {
        posActual := posA[i]
        posNew := posN[i]
        rubik[posActual + "v"] = tempRubik[posNew + "v"]
        rubik[posActual + "c"] = tempRubik[posNew + "c"]
    }
    return rubik
}
func moveU(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate U Side
    posToChange := []string{
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
    posNewVal := []string{
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
func moveUP(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate U Side in reverse
    posToChange := []string{
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
    posNewVal := []string{
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
func moveR(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate R Side
    posToChange := []string{
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
    posNewVal := []string{
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
func moveRP(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate R Side in reverse
    posToChange := []string{
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
    posNewVal := []string{
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
func moveL(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate L Side
    posToChange := []string{
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
    posNewVal := []string{
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
func moveLP(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate L Side in reverse
    posToChange := []string{
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
    posNewVal := []string{
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
func moveF(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate F Side
    posToChange := []string{
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
    posNewVal := []string{
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
func moveFP(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate F Side in reverse
    posToChange := []string{
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
    posNewVal := []string{
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
func moveB(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate B Side
    posToChange := []string{
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
    posNewVal := []string{
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
func moveBP(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate B Side in reverse
    posToChange := []string{
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
    posNewVal := []string{
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
func moveD(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate D Side
    posToChange := []string{
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
    posNewVal := []string{
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
func moveDP(rubik rubikobj) rubikobj {
    /* WITH TEST */
    // Rotate D Side in reverse
    posToChange := []string{
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
    posNewVal := []string{
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
func sexyMove(rubik rubikobj) rubikobj {
    /* WITH TEST */
    rubik = moveR(rubik)
    rubik = moveU(rubik)
    rubik = moveRP(rubik)
    rubik = moveUP(rubik)
    return rubik
}
func printPos(rubik rubikobj, pos string) string {
    primaryColor := rubik[pos + "c"].(string)
    posValue := rubik[pos + "v"].(string)
    return translateColorForCli(primaryColor) + posValue + translateColorForCli("white")
}
func formatCli(rubik rubikobj){
    fmt.Print("         |--------|" + "\n")
    fmt.Print("         |" + printPos(rubik, "B1") + "-" + printPos(rubik, "B2") + "-" + printPos(rubik, "B3") + "|" + "\n")
    fmt.Print("         |" + printPos(rubik, "B4") + "-" + printPos(rubik, "B5") + "-" + printPos(rubik, "B6") + "|" + "\n")
    fmt.Print("         |" + printPos(rubik, "B7") + "-" + printPos(rubik, "B8") + "-" + printPos(rubik, "B9") + "|" + "\n")
    fmt.Print("|--------|--------|--------|--------|" + "\n")
    fmt.Print("|" + printPos(rubik, "L1") + "-" + printPos(rubik, "L2") + "-" + printPos(rubik, "L3") + "|")
    fmt.Print(printPos(rubik, "U1") + "-" + printPos(rubik, "U2") + "-" + printPos(rubik, "U3"))
    fmt.Print("|" + printPos(rubik, "R1") + "-" + printPos(rubik, "R2") + "-" + printPos(rubik, "R3") + "|")
    fmt.Print(printPos(rubik, "D1") + "-" + printPos(rubik, "D2") + "-" + printPos(rubik, "D3") + "|" + "\n")
    fmt.Print("|" + printPos(rubik, "L4") + "-" + printPos(rubik, "L5") + "-" + printPos(rubik, "L6") + "|")
    fmt.Print(printPos(rubik, "U4") + "-" + printPos(rubik, "U5") + "-" + printPos(rubik, "U6"))
    fmt.Print("|" + printPos(rubik, "R4") + "-" + printPos(rubik, "R5") + "-" + printPos(rubik, "R6") + "|")
    fmt.Print(printPos(rubik, "D4") + "-" + printPos(rubik, "D5") + "-" + printPos(rubik, "D6") + "|" + "\n")
    fmt.Print("|" + printPos(rubik, "L7") + "-" + printPos(rubik, "L8") + "-" + printPos(rubik, "L9") + "|")
    fmt.Print(printPos(rubik, "U7") + "-" + printPos(rubik, "U8") + "-" + printPos(rubik, "U9"))
    fmt.Print("|" + printPos(rubik, "R7") + "-" + printPos(rubik, "R8") + "-" + printPos(rubik, "R9") + "|")
    fmt.Print(printPos(rubik, "D7") + "-" + printPos(rubik, "D8") + "-" + printPos(rubik, "D9") + "|" + "\n")
    fmt.Print("|--------|--------|--------|--------|" + "\n")
    fmt.Print("         |" + printPos(rubik, "F1") + "-" + printPos(rubik, "F2") + "-" + printPos(rubik, "F3") + "|" + "\n")
    fmt.Print("         |" + printPos(rubik, "F4") + "-" + printPos(rubik, "F5") + "-" + printPos(rubik, "F6") + "|" + "\n")
    fmt.Print("         |" + printPos(rubik, "F7") + "-" + printPos(rubik, "F8") + "-" + printPos(rubik, "F9") + "|" + "\n")
    fmt.Print("         |--------|" + "\n")
}
func printPosHtml(rubik rubikobj, pos string) string {
    primaryColor := rubik[pos + "c"].(string)
    posValue := rubik[pos + "v"].(string)
    return "<span style=\"background-color: " + primaryColor + "\">" + posValue + "</span>"
}
func formatHtml(rubik rubikobj){
    fmt.Print("<table>")
    fmt.Print(" <tr><td></td><td>" + printPosHtml(rubik, "B1") + "-" + printPosHtml(rubik, "B2") + "-" + printPosHtml(rubik, "B3") + "</td></tr>")
    fmt.Print(" <tr><td></td><td>" + printPosHtml(rubik, "B4") + "-" + printPosHtml(rubik, "B5") + "-" + printPosHtml(rubik, "B6") + "</td></tr>")
    fmt.Print(" <tr><td></td><td>" + printPosHtml(rubik, "B7") + "-" + printPosHtml(rubik, "B8") + "-" + printPosHtml(rubik, "B9") + "</td></tr>")
    fmt.Print(" <tr><td>" + printPosHtml(rubik, "L1") + "-" + printPosHtml(rubik, "L2") + "-" + printPosHtml(rubik, "L3") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "U1") + "-" + printPosHtml(rubik, "U2") + "-" + printPosHtml(rubik, "U3") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "R1") + "-" + printPosHtml(rubik, "R2") + "-" + printPosHtml(rubik, "R3") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "D1") + "-" + printPosHtml(rubik, "D2") + "-" + printPosHtml(rubik, "D3") + "</td></tr>")
    fmt.Print(" <tr><td>" + printPosHtml(rubik, "L4") + "-" + printPosHtml(rubik, "L5") + "-" + printPosHtml(rubik, "L6") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "U4") + "-" + printPosHtml(rubik, "U5") + "-" + printPosHtml(rubik, "U6") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "R4") + "-" + printPosHtml(rubik, "R5") + "-" + printPosHtml(rubik, "R6") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "D4") + "-" + printPosHtml(rubik, "D5") + "-" + printPosHtml(rubik, "D6") + "</td></tr>")
    fmt.Print(" <tr><td>" + printPosHtml(rubik, "L7") + "-" + printPosHtml(rubik, "L8") + "-" + printPosHtml(rubik, "L9") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "U7") + "-" + printPosHtml(rubik, "U8") + "-" + printPosHtml(rubik, "U9") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "R7") + "-" + printPosHtml(rubik, "R8") + "-" + printPosHtml(rubik, "R9") + "</td>")
    fmt.Print("     <td>" + printPosHtml(rubik, "D7") + "-" + printPosHtml(rubik, "D8") + "-" + printPosHtml(rubik, "D9") + "</td></tr>")
    fmt.Print(" <tr><td></td><td>" + printPosHtml(rubik, "F1") + "-" + printPosHtml(rubik, "F2") + "-" + printPosHtml(rubik, "F3") + "</td></tr>")
    fmt.Print(" <tr><td></td><td>" + printPosHtml(rubik, "F4") + "-" + printPosHtml(rubik, "F5") + "-" + printPosHtml(rubik, "F6") + "</td></tr>")
    fmt.Print(" <tr><td></td><td>" + printPosHtml(rubik, "F7") + "-" + printPosHtml(rubik, "F8") + "-" + printPosHtml(rubik, "F9") + "</td></tr>")
    fmt.Print("</table>")
}
func translateColorForCli(color string) string {
    red := "\033[31m"
    green := "\033[32m"
    yellow := "\033[33m"
    blue := "\033[34m"
    magenta := "\033[95m"
    white := "\033[97m"
    colorResult := white
    switch (color) {
        case "red":
            colorResult = red
            break;
        case "green":
            colorResult = green
            break;
        case "yellow":
            colorResult = yellow
            break;
        case "blue":
            colorResult = blue
            break;
        case "magenta":
            colorResult = magenta
            break;
        case "white":
            colorResult = white
            break;
    }
    return colorResult
}

//Rubik Asserts

func assertOrderedRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertOrderedRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertOrderedRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveURubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveURubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveURubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveUPRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveUPRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveUPRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveRRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveRRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveRRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveRPRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveRPRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveRPRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveLRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveLRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveLRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveLPRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveLPRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveLPRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveFRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveFRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveFRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveFPRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveFPRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveFPRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveBRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveBRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveBRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveBPRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveBPRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveBPRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveDRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveDRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveDRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertMoveDPRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveDPRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertMoveDPRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}
func assertSexyMoveRubik(rubik rubikobj, positions []string, values []string, colors []string) {
    assert := true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] != values[i] || rubik[pos + "c"] != colors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertSexyMoveRubik\n" + translateColorForCli("green") + "Assert True\n" + translateColorForCli("white"))
    } else {
        fmt.Print("assertSexyMoveRubik\n" + translateColorForCli("red") + "Assert False\n" + translateColorForCli("white"))
    }
}

func main() {
    // Set assretAux values
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
      "F1","F2","F3","F4","F5","F6","F7","F8","F9",
    }
    orderedValues := [] string {
      "B1","B2","B3","B4","B5","B6","B7","B8","B9",
      "L1","L2","L3","L4","L5","L6","L7","L8","L9",
      "U1","U2","U3","U4","U5","U6","U7","U8","U9",
      "R1","R2","R3","R4","R5","R6","R7","R8","R9",
      "D1","D2","D3","D4","D5","D6","D7","D8","D9",
      "F1","F2","F3","F4","F5","F6","F7","F8","F9",
    }
    orderedColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
      blue,blue,blue,blue,blue,blue,blue,blue,blue,
      yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
      green,green,green,green,green,green,green,green,green,
      white,white,white,white,white,white,white,white,white,
      red,red,red,red,red,red,red,red,red,
    }
    moveUValues := [] string {
      "B1","B2","B3","B4","B5","B6","L9","L6","L3",
      "L1","L2","F1","L4","L5","F2","L7","L8","F3",
      "U7","U4","U1","U8","U5","U2","U9","U6","U3",
      "B7","R2","R3","B8","R5","R6","B9","R8","R9",
      "D1","D2","D3","D4","D5","D6","D7","D8","D9",
      "R7","R4","R1","F4","F5","F6","F7","F8","F9",
    }
    moveUColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,blue,blue,blue,
      blue,blue,red,blue,blue,red,blue,blue,red,
      yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
      magenta,green,green,magenta,green,green,magenta,green,green,
      white,white,white,white,white,white,white,white,white,
      green,green,green,red,red,red,red,red,red,
    }
    moveUPValues := [] string {
      "B1","B2","B3","B4","B5","B6","R1","R4","R7",
      "L1","L2","B9","L4","L5","B8","L7","L8","B7",
      "U3","U6","U9","U2","U5","U8","U1","U4","U7",
      "F3","R2","R3","F2","R5","R6","F1","R8","R9",
      "D1","D2","D3","D4","D5","D6","D7","D8","D9",
      "L3","L6","L9","F4","F5","F6","F7","F8","F9",
    }
    moveUPColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,green,green,green,
      blue,blue,magenta,blue,blue,magenta,blue,blue,magenta,
      yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
      red,green,green,red,green,green,red,green,green,
      white,white,white,white,white,white,white,white,white,
      blue,blue,blue,red,red,red,red,red,red,
    }
    moveRValues := [] string {
      "B1","B2","U3","B4","B5","U6","B7","B8","U9",
      "L1","L2","L3","L4","L5","L6","L7","L8","L9",
      "U1","U2","F3","U4","U5","F6","U7","U8","F9",
      "R7","R4","R1","R8","R5","R2","R9","R6","R3",
      "B9","D2","D3","B6","D5","D6","B3","D8","D9",
      "F1","F2","D7","F4","F5","D4","F7","F8","D1",
    }
    moveRColors := [] string {
      magenta,magenta,yellow,magenta,magenta,yellow,magenta,magenta,yellow,
      blue,blue,blue,blue,blue,blue,blue,blue,blue,
      yellow,yellow,red,yellow,yellow,red,yellow,yellow,red,
      green,green,green,green,green,green,green,green,green,
      magenta,white,white,magenta,white,white,magenta,white,white,
      red,red,white,red,red,white,red,red,white,
    }
    moveRPValues := [] string {
      "B1","B2","D7","B4","B5","D4","B7","B8","D1",
      "L1","L2","L3","L4","L5","L6","L7","L8","L9",
      "U1","U2","B3","U4","U5","B6","U7","U8","B9",
      "R3","R6","R9","R2","R5","R8","R1","R4","R7",
      "F9","D2","D3","F6","D5","D6","F3","D8","D9",
      "F1","F2","U3","F4","F5","U6","F7","F8","U9",
    }
    moveRPColors := [] string {
      magenta,magenta,white,magenta,magenta,white,magenta,magenta,white,
      blue,blue,blue,blue,blue,blue,blue,blue,blue,
      yellow,yellow,magenta,yellow,yellow,magenta,yellow,yellow,magenta,
      green,green,green,green,green,green,green,green,green,
      red,white,white,red,white,white,red,white,white,
      red,red,yellow,red,red,yellow,red,red,yellow,
    }
    moveLValues := [] string {
      "D9","B2","B3","D6","B5","B6","D3","B8","B9",
      "L7","L4","L1","L8","L5","L2","L9","L6","L3",
      "B1","U2","U3","B4","U5","U6","B7","U8","U9",
      "R1","R2","R3","R4","R5","R6","R7","R8","R9",
      "D1","D2","F7","D4","D5","F4","D7","D8","F1",
      "U1","F2","F3","U4","F5","F6","U7","F8","F9",
    }
    moveLColors := [] string {
      white,magenta,magenta,white,magenta,magenta,white,magenta,magenta,
      blue,blue,blue,blue,blue,blue,blue,blue,blue,
      magenta,yellow,yellow,magenta,yellow,yellow,magenta,yellow,yellow,
      green,green,green,green,green,green,green,green,green,
      white,white,red,white,white,red,white,white,red,
      yellow,red,red,yellow,red,red,yellow,red,red,
    }
    moveLPValues := [] string {
      "U1","B2","B3","U4","B5","B6","U7","B8","B9",
      "L3","L6","L9","L2","L5","L8","L1","L4","L7",
      "F1","U2","U3","F4","U5","U6","F7","U8","U9",
      "R1","R2","R3","R4","R5","R6","R7","R8","R9",
      "D1","D2","B7","D4","D5","B4","D7","D8","B1",
      "D9","F2","F3","D6","F5","F6","D3","F8","F9",
    }
    moveLPColors := [] string {
      yellow,magenta,magenta,yellow,magenta,magenta,yellow,magenta,magenta,
      blue,blue,blue,blue,blue,blue,blue,blue,blue,
      red,yellow,yellow,red,yellow,yellow,red,yellow,yellow,
      green,green,green,green,green,green,green,green,green,
      white,white,magenta,white,white,magenta,white,white,magenta,
      white,red,red,white,red,red,white,red,red,
    }
    moveFValues := [] string {
      "B1","B2","B3","B4","B5","B6","B7","B8","B9",
      "L1","L2","L3","L4","L5","L6","D7","D8","D9",
      "U1","U2","U3","U4","U5","U6","L7","L8","L9",
      "R1","R2","R3","R4","R5","R6","U7","U8","U9",
      "D1","D2","D3","D4","D5","D6","R7","R8","R9",
      "F7","F4","F1","F8","F5","F2","F9","F6","F3",
    }
    moveFColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
      blue,blue,blue,blue,blue,blue,white,white,white,
      yellow,yellow,yellow,yellow,yellow,yellow,blue,blue,blue,
      green,green,green,green,green,green,yellow,yellow,yellow,
      white,white,white,white,white,white,green,green,green,
      red,red,red,red,red,red,red,red,red,
    }
    moveFPValues := [] string {
      "B1","B2","B3","B4","B5","B6","B7","B8","B9",
      "L1","L2","L3","L4","L5","L6","U7","U8","U9",
      "U1","U2","U3","U4","U5","U6","R7","R8","R9",
      "R1","R2","R3","R4","R5","R6","D7","D8","D9",
      "D1","D2","D3","D4","D5","D6","L7","L8","L9",
      "F3","F6","F9","F2","F5","F8","F1","F4","F7",
    }
    moveFPColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
      blue,blue,blue,blue,blue,blue,yellow,yellow,yellow,
      yellow,yellow,yellow,yellow,yellow,yellow,green,green,green,
      green,green,green,green,green,green,white,white,white,
      white,white,white,white,white,white,blue,blue,blue,
      red,red,red,red,red,red,red,red,red,
    }
    moveBValues := [] string {
      "B7","B4","B1","B8","B5","B2","B9","B6","B3",
      "U1","U2","U3","L4","L5","L6","L7","L8","L9",
      "R1","R2","R3","U4","U5","U6","U7","U8","U9",
      "D1","D2","D3","R4","R5","R6","R7","R8","R9",
      "L1","L2","L3","D4","D5","D6","D7","D8","D9",
      "F1","F2","F3","F4","F5","F6","F7","F8","F9",
    }
    moveBColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
      yellow,yellow,yellow,blue,blue,blue,blue,blue,blue,
      green,green,green,yellow,yellow,yellow,yellow,yellow,yellow,
      white,white,white,green,green,green,green,green,green,
      blue,blue,blue,white,white,white,white,white,white,
      red,red,red,red,red,red,red,red,red,
    }
    moveBPValues := [] string {
      "B3","B6","B9","B2","B5","B8","B1","B4","B7",
      "D1","D2","D3","L4","L5","L6","L7","L8","L9",
      "L1","L2","L3","U4","U5","U6","U7","U8","U9",
      "U1","U2","U3","R4","R5","R6","R7","R8","R9",
      "R1","R2","R3","D4","D5","D6","D7","D8","D9",
      "F1","F2","F3","F4","F5","F6","F7","F8","F9",
    }
    moveBPColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,magenta,
      white,white,white,blue,blue,blue,blue,blue,blue,
      blue,blue,blue,yellow,yellow,yellow,yellow,yellow,yellow,
      yellow,yellow,yellow,green,green,green,green,green,green,
      green,green,green,white,white,white,white,white,white,
      red,red,red,red,red,red,red,red,red,
    }
    moveDValues := [] string {
      "R3","R6","R9","B4","B5","B6","B7","B8","B9",
      "B3","L2","L3","B2","L5","L6","B1","L8","L9",
      "U1","U2","U3","U4","U5","U6","U7","U8","U9",
      "R1","R2","F9","R4","R5","F8","R7","R8","F7",
      "D7","D4","D1","D8","D5","D2","D9","D6","D3",
      "F1","F2","F3","F4","F5","F6","L1","L4","L7",
    }
    moveDColors := [] string {
      green,green,green,magenta,magenta,magenta,magenta,magenta,magenta,
      magenta,blue,blue,magenta,blue,blue,magenta,blue,blue,
      yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
      green,green,red,green,green,red,green,green,red,
      white,white,white,white,white,white,white,white,white,
      red,red,red,red,red,red,blue,blue,blue,
    }
    moveDPValues := [] string {
      "L7","L4","L1","B4","B5","B6","B7","B8","B9",
      "F7","L2","L3","F8","L5","L6","F9","L8","L9",
      "U1","U2","U3","U4","U5","U6","U7","U8","U9",
      "R1","R2","B1","R4","R5","B2","R7","R8","B3",
      "D3","D6","D9","D2","D5","D8","D1","D4","D7",
      "F1","F2","F3","F4","F5","F6","R9","R6","R3",
    }
    moveDPColors := [] string {
      blue,blue,blue,magenta,magenta,magenta,magenta,magenta,magenta,
      red,blue,blue,red,blue,blue,red,blue,blue,
      yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,yellow,
      green,green,magenta,green,green,magenta,green,green,magenta,
      white,white,white,white,white,white,white,white,white,
      red,red,red,red,red,red,green,green,green,
    }
    sexyMoveValues := [] string {
      "B1","B2","B3","B4","B5","B6","R1","R4","B7",
      "L1","L2","B9","L4","L5","L6","L7","L8","L9",
      "U3","U6","L3","U4","U5","F6","U7","U8","F9",
      "U1","R2","R3","R8","R5","R6","R9","B8","U9",
      "D1","D2","D3","D4","D5","D6","R7","D8","D9",
      "F1","F2","D7","F4","F5","U2","F7","F8","F3",
    }
    sexyMoveColors := [] string {
      magenta,magenta,magenta,magenta,magenta,magenta,green,green,magenta,
      blue,blue,magenta,blue,blue,blue,blue,blue,blue,
      yellow,yellow,blue,yellow,yellow,red,yellow,yellow,red,
      yellow,green,green,green,green,green,green,magenta,yellow,
      white,white,white,white,white,white,green,white,white,
      red,red,white,red,red,yellow,red,red,red,
    }
    // Test getOrderedRubik function
    rubik := getOrderedRubik()
    fmt.Print(translateColorForCli("blue") + "GetRubik\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    // Test moveU function
    rubik = moveU(rubik)
    fmt.Print(translateColorForCli("blue") + "moveU\n" + translateColorForCli("white"))
    assertMoveURubik(rubik, positions, moveUValues, moveUColors)
    // Test moveUP function
    rubik = moveUP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveUP\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    rubik = moveUP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveUP\n" + translateColorForCli("white"))
    assertMoveUPRubik(rubik, positions, moveUPValues, moveUPColors)
    rubik = moveU(rubik)
    // Test moveR function
    rubik = moveR(rubik)
    fmt.Print(translateColorForCli("blue") + "moveR\n" + translateColorForCli("white"))
    assertMoveRRubik(rubik, positions, moveRValues, moveRColors)
    // Test moveRP function
    rubik = moveRP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveRP\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    rubik = moveRP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveRP\n" + translateColorForCli("white"))
    assertMoveRPRubik(rubik, positions, moveRPValues, moveRPColors)
    rubik = moveR(rubik)
    // Test moveL function
    rubik = moveL(rubik)
    fmt.Print(translateColorForCli("blue") + "moveL\n" + translateColorForCli("white"))
    assertMoveLRubik(rubik, positions, moveLValues, moveLColors)
    // Test moveLP function
    rubik = moveLP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveLP\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    rubik = moveLP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveLP\n" + translateColorForCli("white"))
    assertMoveLPRubik(rubik, positions, moveLPValues, moveLPColors)
    rubik = moveL(rubik)
    // Test moveF function
    rubik = moveF(rubik)
    fmt.Print(translateColorForCli("blue") + "moveF\n" + translateColorForCli("white"))
    assertMoveFRubik(rubik, positions, moveFValues, moveFColors)
    // Test moveFP function
    rubik = moveFP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveFP\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    rubik = moveFP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveFP\n" + translateColorForCli("white"))
    assertMoveFPRubik(rubik, positions, moveFPValues, moveFPColors)
    rubik = moveF(rubik)
    // Test MoveB function
    rubik = moveB(rubik)
    fmt.Print(translateColorForCli("blue") + "moveB\n" + translateColorForCli("white"))
    assertMoveBRubik(rubik, positions, moveBValues, moveBColors)
    // Test MoveBP function
    rubik = moveBP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveBP\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    rubik = moveBP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveBP\n" + translateColorForCli("white"))
    assertMoveBPRubik(rubik, positions, moveBPValues, moveBPColors)
    rubik = moveB(rubik)
    // Test MoveD function
    rubik = moveD(rubik)
    fmt.Print(translateColorForCli("blue") + "moveD\n" + translateColorForCli("white"))
    assertMoveDRubik(rubik, positions, moveDValues, moveDColors)
    // Test MoveDP function
    rubik = moveDP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveDP\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    rubik = moveDP(rubik)
    fmt.Print(translateColorForCli("blue") + "moveDP\n" + translateColorForCli("white"))
    assertMoveDPRubik(rubik, positions, moveDPValues, moveDPColors)
    rubik = moveD(rubik)
    // Test sexy move function
    rubik = sexyMove(rubik)
    fmt.Print(translateColorForCli("blue") + "sexyMove\n" + translateColorForCli("white"))
    assertSexyMoveRubik(rubik, positions, sexyMoveValues, sexyMoveColors)
    // 5 more sexy moves then cube is ordered
    for j := 0; j < 5; j++ {
        rubik = sexyMove(rubik)
    }
    fmt.Print(translateColorForCli("blue") + "5 more sexyMoves\n" + translateColorForCli("white"))
    assertOrderedRubik(rubik, positions, orderedValues, orderedColors)
    formatCli(rubik)
}
