#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
#include "rubik.h"

void initializeRubik(RubikCube *rubik) {
    char *sides[] = {"B", "L", "U", "R", "D", "F"};
    for (int s = 0; s < 6; s++) {
        for (int i = 1; i < 10; i++) {
            sprintf(rubik->pieces[s * 9 + i - 1].v, "%s%d", sides[s], i);
            switch (sides[s][0]) {
                case 'B': strcpy(rubik->pieces[s * 9 + i - 1].c, MAGENTA); break;
                case 'L': strcpy(rubik->pieces[s * 9 + i - 1].c, BLUE); break;
                case 'U': strcpy(rubik->pieces[s * 9 + i - 1].c, YELLOW); break;
                case 'R': strcpy(rubik->pieces[s * 9 + i - 1].c, GREEN); break;
                case 'D': strcpy(rubik->pieces[s * 9 + i - 1].c, WHITE); break;
                case 'F': strcpy(rubik->pieces[s * 9 + i - 1].c, RED); break;
            }
        }
    }
}

int getIndex(const char *pos) {
    switch (pos[0]) {
        case 'B': return (pos[1] - '1');
        case 'L': return 9 + (pos[1] - '1');
        case 'U': return 18 + (pos[1] - '1');
        case 'R': return 27 + (pos[1] - '1');
        case 'D': return 36 + (pos[1] - '1');
        case 'F': return 45 + (pos[1] - '1');
        default: return -1;
    }
}

void move(RubikCube *rubik, char *posToChange[], char *posNewVal[]) {
    RubikPiece temp[54];
    memcpy(temp, rubik->pieces, sizeof(RubikPiece) * 54);
    for (int i = 0; i < 20; i++) { // Ajusta a 20 elementos
        int oldIndex = getIndex(posToChange[i]);
        int newIndex = getIndex(posNewVal[i]);
        if (oldIndex >= 0 && oldIndex < 54 && newIndex >= 0 && newIndex < 54) {
            strcpy(rubik->pieces[oldIndex].v, temp[newIndex].v);
            strcpy(rubik->pieces[oldIndex].c, temp[newIndex].c);
        } else {
            fprintf(stderr, "Invalid index: oldIndex=%d, newIndex=%d\n", oldIndex, newIndex);
        }
    }
}

void printRubik(RubikCube *rubik) {
    for (int i = 0; i < 54; i++) {
        printf("%s%s%s%s ", rubik->pieces[i].c, rubik->pieces[i].v, RESET, WHITE);
        if ((i + 1) % 9 == 0) printf("\n");
    }
}

void moveU(RubikCube *rubik) {
    char *posToChange[] = {
        "U1", "U2", "U3", "U6", "U9", "U8", "U7", "U4",
        "R1", "R4", "R7", "F1", "F2", "F3", "L3", "L6", "L9",
        "B7", "B8", "B9"
    };
    char *posNewVal[] = {
        "U7", "U4", "U1", "U2", "U3", "U6", "U9", "U8",
        "B7", "B8", "B9", "R7", "R4", "R1", "F1", "F2", "F3",
        "L9", "L6", "L3"
    };
    move(rubik, posToChange, posNewVal);
}

void moveUP(RubikCube *rubik) {
    char *posToChange[] = {
        "U1", "U2", "U3", "U6", "U9", "U8", "U7", "U4",
        "R1", "R4", "R7", "F1", "F2", "F3", "L3", "L6", "L9",
        "B7", "B8", "B9"
    };
    char *posNewVal[] = {
        "U3", "U6", "U9", "U8", "U7", "U4", "U1", "U2",
        "F3", "F2", "F1", "L3", "L6", "L9", "B9", "B8", "B7",
        "R1", "R4", "R7"
    };
    move(rubik, posToChange, posNewVal);
}

void moveR(RubikCube *rubik) {
    char *posToChange[] = {
        "R1", "R2", "R3", "R6", "R9", "R7", "R8", "R4",
        "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9",
        "U3", "U6", "U9"
    };
    char *posNewVal[] = {
        "R7", "R4", "R1", "R2", "R3", "R9", "R6", "R8",
        "U3", "U6", "U9", "B3", "B6", "B9", "D7", "D4", "D1",
        "F3", "F6", "F9"
    };
    move(rubik, posToChange, posNewVal);
}

void moveRP(RubikCube *rubik) {
    char *posToChange[] = {
        "R1", "R2", "R3", "R6", "R9", "R8", "R7", "R4",
        "B3", "B6", "B9", "D7", "D4", "D1", "F3", "F6", "F9",
        "U3", "U6", "U9"
    };
    char *posNewVal[] = {
        "R3", "R6", "R9", "R8", "R7", "R4", "R1", "R2",
        "D7", "D4", "D1", "F3", "F6", "F9", "U3", "U6", "U9",
        "B3", "B6", "B9"
    };
    move(rubik, posToChange, posNewVal);
}

void moveL(RubikCube *rubik) {
    char *posToChange[] = {
        "L1", "L2", "L3", "L6", "L9", "L7", "L8", "L4",
        "B1", "B4", "B7", "U1", "U4", "U7", "F1", "F4", "F7",
        "D3", "D6", "D9"
    };
    char *posNewVal[] = {
        "L7", "L4", "L1", "L2", "L3", "L9", "L6", "L8",
        "D9", "D6", "D3", "B1", "B4", "B7", "U1", "U4", "U7",
        "F7", "F4", "F1"
    };
    move(rubik, posToChange, posNewVal);
}

void moveLP(RubikCube *rubik) {
    char *posToChange[] = {
        "L1", "L2", "L3", "L6", "L9", "L8", "L7", "L4",
        "B1", "B4", "B7", "U1", "U4", "U7", "F1", "F4", "F7",
        "D3", "D6", "D9"
    };
    char *posNewVal[] = {
        "L3", "L6", "L9", "L8", "L7", "L4", "L1", "L2",
        "U1", "U4", "U7", "F1", "F4", "F7", "D9", "D6", "D3",
        "B7", "B4", "B1"
    };
    move(rubik, posToChange, posNewVal);
}

void moveF(RubikCube *rubik) {
    char *posToChange[] = {
        "F1", "F2", "F3", "F6", "F9", "F7", "F8", "F4",
        "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9",
        "D7", "D8", "D9"
    };
    char *posNewVal[] = {
        "F7", "F4", "F1", "F2", "F3", "F9", "F6", "F8",
        "D7", "D8", "D9", "L7", "L8", "L9", "U7", "U8", "U9",
        "R7", "R8", "R9"
    };
    move(rubik, posToChange, posNewVal);
}

void moveFP(RubikCube *rubik) {
    char *posToChange[] = {
        "F1", "F2", "F3", "F6", "F9", "F8", "F7", "F4",
        "L7", "L8", "L9", "U7", "U8", "U9", "R7", "R8", "R9",
        "D7", "D8", "D9"
    };
    char *posNewVal[] = {
        "F3", "F6", "F9", "F8", "F7", "F4", "F1", "F2",
        "U7", "U8", "U9", "R7", "R8", "R9", "D7", "D8", "D9",
        "L7", "L8", "L9"
    };
    move(rubik, posToChange, posNewVal);
}

void moveB(RubikCube *rubik) {
    char *posToChange[] = {
        "B1", "B2", "B3", "B6", "B9", "B8", "B7", "B4",
        "L1", "L2", "L3", "U1", "U2", "U3", "R1", "R2", "R3",
        "D1", "D2", "D3"
    };
    char *posNewVal[] = {
        "B7", "B4", "B1", "B2", "B3", "B6", "B9", "B8",
        "U1", "U2", "U3", "R1", "R2", "R3", "D1", "D2", "D3",
        "L1", "L2", "L3"
    };
    move(rubik, posToChange, posNewVal);
}

void moveBP(RubikCube *rubik) {
    char *posToChange[] = {
        "B1", "B2", "B3", "B6", "B9", "B8", "B7", "B4",
        "L1", "L2", "L3", "U1", "U2", "U3", "R1", "R2", "R3",
        "D1", "D2", "D3"
    };
    char *posNewVal[] = {
        "B3", "B6", "B9", "B8", "B7", "B4", "B1", "B2",
        "D1", "D2", "D3", "L1", "L2", "L3", "U1", "U2", "U3",
        "R1", "R2", "R3"
    };
    move(rubik, posToChange, posNewVal);
}

void moveD(RubikCube *rubik) {
    char *posToChange[] = {
        "D1", "D2", "D3", "D6", "D9", "D8", "D7", "D4",
        "B1", "B2", "B3", "L1", "L4", "L7", "R3", "R6", "R9",
        "F7", "F8", "F9"
    };
    char *posNewVal[] = {
        "D7", "D4", "D1", "D2", "D3", "D6", "D9", "D8",
        "R3", "R6", "R9", "B3", "B2", "B1", "F9", "F8", "F7",
        "L1", "L4", "L7"
    };
    move(rubik, posToChange, posNewVal);
}

void moveDP(RubikCube *rubik) {
    char *posToChange[] = {
        "D1", "D2", "D3", "D6", "D9", "D8", "D7", "D4",
        "B1", "B2", "B3", "L1", "L4", "L7", "R3", "R6", "R9",
        "F7", "F8", "F9"
    };
    char *posNewVal[] = {
        "D3", "D6", "D9", "D8", "D7", "D4", "D1", "D2",
        "L7", "L4", "L1", "F7", "F8", "F9", "B1", "B2", "B3",
        "R9", "R6", "R3"
    };
    move(rubik, posToChange, posNewVal);
}

void sexyMove(RubikCube *rubik) {
    moveR(rubik);
    moveU(rubik);
    moveRP(rubik);
    moveUP(rubik);
}

char *printPos(RubikCube *rubik, const char *pos) {
    static char formattedPos[54][20]; // Un buffer para cada posición (54 posiciones en total)
    static int callCount = 0; // Llevar un conteo de la posición actual del buffer

    int index = getIndex(pos); // Obtener el índice
    if (index < 0 || index >= 54) {
        snprintf(formattedPos[callCount % 54], sizeof(formattedPos[callCount % 54]), "ERR");
    } else {
        snprintf(formattedPos[callCount % 54], sizeof(formattedPos[callCount % 54]),
                 "%s%s%s", rubik->pieces[index].c, rubik->pieces[index].v, RESET);
    }
    return formattedPos[callCount++ % 54];
}

// Formatear y mostrar el cubo en la terminal
void formatCli(RubikCube *rubik) {
    char buffer[54][20]; // Un buffer para cada posición
    for (int i = 0; i < 54; i++) {
        snprintf(buffer[i], sizeof(buffer[i]), "%s%s%s", rubik->pieces[i].c, rubik->pieces[i].v, RESET);
    }

    printf(
        "         |--------|\n"
        "         |%s-%s-%s|\n"
        "         |%s-%s-%s|\n"
        "         |%s-%s-%s|\n"
        "|--------|--------|--------|--------|\n"
        "|%s-%s-%s|%s-%s-%s|%s-%s-%s|%s-%s-%s|\n"
        "|%s-%s-%s|%s-%s-%s|%s-%s-%s|%s-%s-%s|\n"
        "|%s-%s-%s|%s-%s-%s|%s-%s-%s|%s-%s-%s|\n"
        "|--------|--------|--------|--------|\n"
        "         |%s-%s-%s|\n"
        "         |%s-%s-%s|\n"
        "         |%s-%s-%s|\n"
        "         |--------|\n",
        // Cara B
        buffer[getIndex("B1")], buffer[getIndex("B2")], buffer[getIndex("B3")],
        buffer[getIndex("B4")], buffer[getIndex("B5")], buffer[getIndex("B6")],
        buffer[getIndex("B7")], buffer[getIndex("B8")], buffer[getIndex("B9")],
        // Línea principal: L, U, R, D
        buffer[getIndex("L1")], buffer[getIndex("L2")], buffer[getIndex("L3")],
        buffer[getIndex("U1")], buffer[getIndex("U2")], buffer[getIndex("U3")],
        buffer[getIndex("R1")], buffer[getIndex("R2")], buffer[getIndex("R3")],
        buffer[getIndex("D1")], buffer[getIndex("D2")], buffer[getIndex("D3")],
        buffer[getIndex("L4")], buffer[getIndex("L5")], buffer[getIndex("L6")],
        buffer[getIndex("U4")], buffer[getIndex("U5")], buffer[getIndex("U6")],
        buffer[getIndex("R4")], buffer[getIndex("R5")], buffer[getIndex("R6")],
        buffer[getIndex("D4")], buffer[getIndex("D5")], buffer[getIndex("D6")],
        buffer[getIndex("L7")], buffer[getIndex("L8")], buffer[getIndex("L9")],
        buffer[getIndex("U7")], buffer[getIndex("U8")], buffer[getIndex("U9")],
        buffer[getIndex("R7")], buffer[getIndex("R8")], buffer[getIndex("R9")],
        buffer[getIndex("D7")], buffer[getIndex("D8")], buffer[getIndex("D9")],
        // Cara F
        buffer[getIndex("F1")], buffer[getIndex("F2")], buffer[getIndex("F3")],
        buffer[getIndex("F4")], buffer[getIndex("F5")], buffer[getIndex("F6")],
        buffer[getIndex("F7")], buffer[getIndex("F8")], buffer[getIndex("F9")]
    );
}

// Verifica si dos cubos son iguales
bool compareCubes(RubikCube *c1, RubikCube *c2) {
    for (int i = 0; i < 54; i++) {
        if (strcmp(c1->pieces[i].v, c2->pieces[i].v) != 0 ||
            strcmp(c1->pieces[i].c, c2->pieces[i].c) != 0) {
            return false;
        }
    }
    return true;
}

// Deshabilitar el main si RUBIK_LIB está definido
#ifndef RUBIK_LIB
// Manejo de argumentos y movimientos
int main(int argc, char *argv[]) {
    RubikCube rubik, initialRubik;
    initializeRubik(&rubik);
    initializeRubik(&initialRubik); // Guardamos el cubo inicial para comparaciones

    bool suppressOutput = false; // Variable para manejar la opción -s
    char *sequence = NULL;

    // Procesar argumentos
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "-s") == 0) {
            suppressOutput = true; // Si se pasa -s, activa la bandera
        } else {
            sequence = argv[i]; // La secuencia de movimientos
        }
    }

    // Realizar movimientos si se proporciona una secuencia
    if (sequence) {
        for (int j = 0; j < strlen(sequence); j++) {
            switch (sequence[j]) {
                case 'U': moveU(&rubik); break;
                case 'u': moveUP(&rubik); break;
                case 'F': moveF(&rubik); break;
                case 'f': moveFP(&rubik); break;
                case 'R': moveR(&rubik); break;
                case 'r': moveRP(&rubik); break;
                case 'L': moveL(&rubik); break;
                case 'l': moveLP(&rubik); break;
                case 'D': moveD(&rubik); break;
                case 'd': moveDP(&rubik); break;
                case 'B': moveB(&rubik); break;
                case 'b': moveBP(&rubik); break;
                default:
                    fprintf(stderr, "Unknown move: %c\n", sequence[j]);
                    return 1;
            }
        }
    }

    // Si se pasa -s, verificamos el estado y no imprimimos el cubo
    if (suppressOutput) {
        if (compareCubes(&rubik, &initialRubik)) {
            printf("Ordered\n");
        } else {
            printf("Unordered\n");
        }
    } else {
        // Imprimir el cubo antes y después de los movimientos
        printf("Initial Rubik's Cube:\n");
        formatCli(&initialRubik); // El cubo inicial
        printf("\nAfter moves:\n");
        formatCli(&rubik); // El cubo modificado
    }

    return 0;
}
#endif
