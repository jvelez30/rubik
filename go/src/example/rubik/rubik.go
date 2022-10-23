// Code generated. MUST EDIT!
package main

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


func main() {
    rubik := getOrderedRubik()
    rubik = moveDP(rubik)
    //fmt.Println(rubik)
    //for j :=0; j<6; j++ {
    //    rubik = sexyMove(rubik);
    //}

    //if isset(argv[0]) /* cannot convert */ {
        // if command line cli format is displayed
       formatCli(rubik)
    //} else {
        // else html format is displayed
    //    formatHtml(rubik)
    //}
}


// Filter get var
// movesParam = filter_input(INPUT_GET, "m", FILTER_SANITIZE_STRING)
// If moves are setted via string query
// if isset(movesParam) /* cannot convert */ {
//     moves = str_split(movesParam)
//     for _, move := range moves {
//         switch (move) {
//             case "U":
//                 rubik = moveU(rubik)
//                 break;
//             case "u":
//                 rubik = moveUP(rubik)
//                 break;
//             case "R":
//                 rubik = moveR(rubik)
//                 break;
//             case "r":
//                 rubik = moveRP(rubik)
//                 break;
//             case "L":
//                 rubik = moveL(rubik)
//                 break;
//             case "l":
//                 rubik = moveLP(rubik)
//                 break;
//             case "F":
//                 rubik = moveF(rubik)
//                 break;
//             case "f":
//                 rubik = moveFP(rubik)
//                 break;
//             case "B":
//                 rubik = moveB(rubik)
//                 break;
//             case "b":
//                 rubik = moveBP(rubik)
//                 break;
//             case "D":
//                 rubik = moveD(rubik)
//                 break;
//             case "d":
//                 rubik = moveDP(rubik)
//                 break;
//         }
//     }
// }

