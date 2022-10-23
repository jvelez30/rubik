package rubikAsserts
/* Rubik Asserts */
require "example/rubik/assertAux"

type rubikobj map[string] interface{}

func assertOrderedRubik(rubik rubikobj) {
    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== orderedValues[i] || rubik[pos + "c"] !== orderedColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertOrderedRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertOrderedRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveURubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveUValues[i] || rubik[pos + "c"] !== moveUColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveURubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveURubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveUPRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveUPValues[i] || rubik[pos + "c"] !== moveUPColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveUPRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveUPRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveRRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveRValues[i] || rubik[pos + "c"] !== moveRColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveRRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveRRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveRPRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveRPValues[i] || rubik[pos + "c"] !== moveRPColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveRPRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveRPRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveLRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveLValues[i] || rubik[pos + "c"] !== moveLColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveLRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveLRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveLPRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveLPValues[i] || rubik[pos + "c"] !== moveLPColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveLPRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveLPRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveFRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveFValues[i] || rubik[pos + "c"] !== moveFColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveFRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveFRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveFPRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveFPValues[i] || rubik[pos + "c"] !== moveFPColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveFPRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveFPRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveBRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveBValues[i] || rubik[pos + "c"] !== moveBColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveBRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveBRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveBPRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveBPValues[i] || rubik[pos + "c"] !== moveBPColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveBPRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveBPRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveDRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveDValues[i] || rubik[pos + "c"] !== moveDColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveDRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveDRubik\n" . red . "Assert False\n" . white)
    }
}
func assertMoveDPRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== moveDPValues[i] || rubik[pos + "c"] !== moveDPColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertMoveDPRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertMoveDPRubik\n" . red . "Assert False\n" . white)
    }
}
func assertSexyMoveRubik(rubik rubikobj) {

    assert = true
    for i := 0; i < len(positions); i++ {
        pos := positions[i]
        if rubik[pos + "v"] !== sexyMoveValues[i] || rubik[pos + "c"] !== sexyMoveColors[i] {
            assert = false
        }
    }
    if assert {
        fmt.Print("assertSexyMoveRubik\n" . green . "Assert True\n" . white)
    } else {
        fmt.Print("assertSexyMoveRubik\n" . red . "Assert False\n" . white)
    }
}