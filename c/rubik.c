#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <stdbool.h>
#include "rubik.h"

/* Macros para mapear "B1" => 0, "B2" => 1, ..., "F9" => 53
   en el mismo orden en que se inicializa el cubo (B, L, U, R, D, F). */

// Cara B (indices [0..8])
#define B1  0
#define B2  1
#define B3  2
#define B4  3
#define B5  4
#define B6  5
#define B7  6
#define B8  7
#define B9  8

// Cara L (indices [9..17])
#define L1  9
#define L2  10
#define L3  11
#define L4  12
#define L5  13
#define L6  14
#define L7  15
#define L8  16
#define L9  17

// Cara U (indices [18..26])
#define U1  18
#define U2  19
#define U3  20
#define U4  21
#define U5  22
#define U6  23
#define U7  24
#define U8  25
#define U9  26

// Cara R (indices [27..35])
#define R1  27
#define R2  28
#define R3  29
#define R4  30
#define R5  31
#define R6  32
#define R7  33
#define R8  34
#define R9  35

// Cara D (indices [36..44])
#define D1  36
#define D2  37
#define D3  38
#define D4  39
#define D5  40
#define D6  41
#define D7  42
#define D8  43
#define D9  44

// Cara F (indices [45..53])
#define F1  45
#define F2  46
#define F3  47
#define F4  48
#define F5  49
#define F6  50
#define F7  51
#define F8  52
#define F9  53


/* 
   Inicializamos un arreglo de 54 etiquetas ("B1", "B2", ..., "F9") y 
   otro de 54 colores (MAGENTA, BLUE, YELLOW, GREEN, WHITE, RED) según la cara.
*/
void initializeRubik(RubikCube *rubik) {
    static const char *labels[54] = {
        // B: [0..8]
        "B1","B2","B3","B4","B5","B6","B7","B8","B9",
        // L: [9..17]
        "L1","L2","L3","L4","L5","L6","L7","L8","L9",
        // U: [18..26]
        "U1","U2","U3","U4","U5","U6","U7","U8","U9",
        // R: [27..35]
        "R1","R2","R3","R4","R5","R6","R7","R8","R9",
        // D: [36..44]
        "D1","D2","D3","D4","D5","D6","D7","D8","D9",
        // F: [45..53]
        "F1","F2","F3","F4","F5","F6","F7","F8","F9"
    };

    // Según el orden B=0..8, L=9..17, U=18..26, R=27..35, D=36..44, F=45..53
    // Asignamos colores ANSI
    static const char *colors[54] = {
        /* B (MAGENTA) 0..8 */   MAGENTA, MAGENTA, MAGENTA, MAGENTA, MAGENTA, MAGENTA, MAGENTA, MAGENTA, MAGENTA,
        /* L (BLUE) 9..17 */     BLUE,    BLUE,    BLUE,    BLUE,    BLUE,    BLUE,    BLUE,    BLUE,    BLUE,
        /* U (YELLOW) 18..26 */  YELLOW,  YELLOW,  YELLOW,  YELLOW,  YELLOW,  YELLOW,  YELLOW,  YELLOW,  YELLOW,
        /* R (GREEN) 27..35 */   GREEN,   GREEN,   GREEN,   GREEN,   GREEN,   GREEN,   GREEN,   GREEN,   GREEN,
        /* D (WHITE) 36..44 */   WHITE,   WHITE,   WHITE,   WHITE,   WHITE,   WHITE,   WHITE,   WHITE,   WHITE,
        /* F (RED) 45..53 */     RED,     RED,     RED,     RED,     RED,     RED,     RED,     RED,     RED
    };

    // Copiamos labels y colores al arreglo rubik->pieces
    for (int i = 0; i < 54; i++) {
        strcpy(rubik->pieces[i].v, labels[i]);
        strcpy(rubik->pieces[i].c, colors[i]);
    }
}


/* Función genérica para aplicar una permutación:
   - oldIndices: las posiciones (en rubik->pieces) que se van a actualizar
   - newIndices: las posiciones de donde se copian los datos
   - count: cuántos elementos permutar
*/
void applyPermutation(RubikCube *rubik, const int *oldIndices, const int *newIndices, int count) {
    RubikPiece temp[54];
    memcpy(temp, rubik->pieces, sizeof(temp)); // Copiamos una sola vez

    for (int i = 0; i < count; i++) {
        int oldIdx = oldIndices[i];
        int newIdx = newIndices[i];
        strcpy(rubik->pieces[oldIdx].v, temp[newIdx].v);
        strcpy(rubik->pieces[oldIdx].c, temp[newIdx].c);
    }
}

/* A continuación definimos, para cada rotación, los arreglos de permutaciones (old->new).
   Cada rotación permuta 20 stickers: 8 del propio "anillo" + 12 alrededor.
   Ejemplo: moveU => rotar la cara U (U1..U9) y modificar los bordes R1..R3, F1..F3, L1..L3, B1..B3 etc. 
   según el orden que manejas en tu código.
*/

/* U (rotación de la cara Up en el sentido horario) */
void moveU(RubikCube *rubik) {
    static const int U_posToChange[20] = {
        U1, U2, U3, U6, U9, U8, U7, U4,
        R1, R4, R7, F1, F2, F3, L3, L6, L9,
        B7, B8, B9
    };
    static const int U_posNewVal[20] = {
        U7, U4, U1, U2, U3, U6, U9, U8,
        B7, B8, B9, R7, R4, R1, F1, F2, F3,
        L9, L6, L3
    };
    applyPermutation(rubik, U_posToChange, U_posNewVal, 20);
}

/* U' (U inverso, en sentido antihorario) */
void moveUP(RubikCube *rubik) {
    static const int U_posToChange[20] = {
        U1, U2, U3, U6, U9, U8, U7, U4,
        R1, R4, R7, F1, F2, F3, L3, L6, L9,
        B7, B8, B9
    };
    static const int U_posNewVal[20] = {
        U3, U6, U9, U8, U7, U4, U1, U2,
        F3, F2, F1, L3, L6, L9, B9, B8, B7,
        R1, R4, R7
    };
    applyPermutation(rubik, U_posToChange, U_posNewVal, 20);
}

/* R */
void moveR(RubikCube *rubik) {
    static const int R_posToChange[20] = {
        R1, R2, R3, R6, R9, R7, R8, R4,
        B3, B6, B9, D7, D4, D1, F3, F6, F9,
        U3, U6, U9
    };
    static const int R_posNewVal[20] = {
        R7, R4, R1, R2, R3, R9, R6, R8,
        U3, U6, U9, B3, B6, B9, D7, D4, D1,
        F3, F6, F9
    };
    applyPermutation(rubik, R_posToChange, R_posNewVal, 20);
}

/* R' */
void moveRP(RubikCube *rubik) {
    static const int R_posToChange[20] = {
        R1, R2, R3, R6, R9, R8, R7, R4,
        B3, B6, B9, D7, D4, D1, F3, F6, F9,
        U3, U6, U9
    };
    static const int R_posNewVal[20] = {
        R3, R6, R9, R8, R7, R4, R1, R2,
        D7, D4, D1, F3, F6, F9, U3, U6, U9,
        B3, B6, B9
    };
    applyPermutation(rubik, R_posToChange, R_posNewVal, 20);
}

/* L */
void moveL(RubikCube *rubik) {
    static const int L_posToChange[20] = {
        L1, L2, L3, L6, L9, L7, L8, L4,
        B1, B4, B7, U1, U4, U7, F1, F4, F7,
        D3, D6, D9
    };
    static const int L_posNewVal[20] = {
        L7, L4, L1, L2, L3, L9, L6, L8,
        D9, D6, D3, B1, B4, B7, U1, U4, U7,
        F7, F4, F1
    };
    applyPermutation(rubik, L_posToChange, L_posNewVal, 20);
}

/* L' */
void moveLP(RubikCube *rubik) {
    static const int L_posToChange[20] = {
        L1, L2, L3, L6, L9, L8, L7, L4,
        B1, B4, B7, U1, U4, U7, F1, F4, F7,
        D3, D6, D9
    };
    static const int L_posNewVal[20] = {
        L3, L6, L9, L8, L7, L4, L1, L2,
        U1, U4, U7, F1, F4, F7, D9, D6, D3,
        B7, B4, B1
    };
    applyPermutation(rubik, L_posToChange, L_posNewVal, 20);
}

/* F */
void moveF(RubikCube *rubik) {
    static const int F_posToChange[20] = {
        F1, F2, F3, F6, F9, F7, F8, F4,
        L7, L8, L9, U7, U8, U9, R7, R8, R9,
        D7, D8, D9
    };
    static const int F_posNewVal[20] = {
        F7, F4, F1, F2, F3, F9, F6, F8,
        D7, D8, D9, L7, L8, L9, U7, U8, U9,
        R7, R8, R9
    };
    applyPermutation(rubik, F_posToChange, F_posNewVal, 20);
}

/* F' */
void moveFP(RubikCube *rubik) {
    static const int F_posToChange[20] = {
        F1, F2, F3, F6, F9, F8, F7, F4,
        L7, L8, L9, U7, U8, U9, R7, R8, R9,
        D7, D8, D9
    };
    static const int F_posNewVal[20] = {
        F3, F6, F9, F8, F7, F4, F1, F2,
        U7, U8, U9, R7, R8, R9, D7, D8, D9,
        L7, L8, L9
    };
    applyPermutation(rubik, F_posToChange, F_posNewVal, 20);
}

/* B */
void moveB(RubikCube *rubik) {
    static const int B_posToChange[20] = {
        B1, B2, B3, B6, B9, B8, B7, B4,
        L1, L2, L3, U1, U2, U3, R1, R2, R3,
        D1, D2, D3
    };
    static const int B_posNewVal[20] = {
        B7, B4, B1, B2, B3, B6, B9, B8,
        U1, U2, U3, R1, R2, R3, D1, D2, D3,
        L1, L2, L3
    };
    applyPermutation(rubik, B_posToChange, B_posNewVal, 20);
}

/* B' */
void moveBP(RubikCube *rubik) {
    static const int B_posToChange[20] = {
        B1, B2, B3, B6, B9, B8, B7, B4,
        L1, L2, L3, U1, U2, U3, R1, R2, R3,
        D1, D2, D3
    };
    static const int B_posNewVal[20] = {
        B3, B6, B9, B8, B7, B4, B1, B2,
        D1, D2, D3, L1, L2, L3, U1, U2, U3,
        R1, R2, R3
    };
    applyPermutation(rubik, B_posToChange, B_posNewVal, 20);
}

/* D */
void moveD(RubikCube *rubik) {
    static const int D_posToChange[20] = {
        D1, D2, D3, D6, D9, D8, D7, D4,
        B1, B2, B3, L1, L4, L7, R3, R6, R9,
        F7, F8, F9
    };
    static const int D_posNewVal[20] = {
        D7, D4, D1, D2, D3, D6, D9, D8,
        R3, R6, R9, B3, B2, B1, F9, F8, F7,
        L1, L4, L7
    };
    applyPermutation(rubik, D_posToChange, D_posNewVal, 20);
}

/* D' */
void moveDP(RubikCube *rubik) {
    static const int D_posToChange[20] = {
        D1, D2, D3, D6, D9, D8, D7, D4,
        B1, B2, B3, L1, L4, L7, R3, R6, R9,
        F7, F8, F9
    };
    static const int D_posNewVal[20] = {
        D3, D6, D9, D8, D7, D4, D1, D2,
        L7, L4, L1, F7, F8, F9, B1, B2, B3,
        R9, R6, R3
    };
    applyPermutation(rubik, D_posToChange, D_posNewVal, 20);
}

/* Ejemplo de un "sexy move": R U R' U' */
void sexyMove(RubikCube *rubik) {
    moveR(rubik);
    moveU(rubik);
    moveRP(rubik);
    moveUP(rubik);
}

/* 
   Función para imprimir el cubo en 6 filas de 9 stickers 
   (similar a tu printRubik original).
*/
void printRubik(RubikCube *rubik) {
    for (int i = 0; i < 54; i++) {
        // pieces[i].c es el color ANSI, pieces[i].v la etiqueta
        printf("%s%s%s%s ", rubik->pieces[i].c, rubik->pieces[i].v, RESET, WHITE);
        if ((i + 1) % 9 == 0) printf("\n");
    }
}

/* Formatear y mostrar el cubo en la terminal con layout 2D */
void formatCli(RubikCube *rubik) {
    // Para abreviar:
    #define P(i) rubik->pieces[(i)]

    // Imprime la cara B (superior)
    printf(
        "         |--------|\n"
        "         |%s%s%s-%s%s%s-%s%s%s|\n"
        "         |%s%s%s-%s%s%s-%s%s%s|\n"
        "         |%s%s%s-%s%s%s-%s%s%s|\n"
        "|--------|--------|--------|--------|\n",
        /* 1era fila de la cara B */
        P(B1).c, P(B1).v, RESET, P(B2).c, P(B2).v, RESET, P(B3).c, P(B3).v, RESET,
        /* 2da fila de la cara B */
        P(B4).c, P(B4).v, RESET, P(B5).c, P(B5).v, RESET, P(B6).c, P(B6).v, RESET,
        /* 3era fila de la cara B */
        P(B7).c, P(B7).v, RESET, P(B8).c, P(B8).v, RESET, P(B9).c, P(B9).v, RESET
    );

    // L, U, R, D en 3 filas
    printf(
        // Fila 1: L1..L3, U1..U3, R1..R3, D1..D3
        "|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|\n"
        // Fila 2: L4..L6, U4..U6, R4..R6, D4..D6
        "|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|\n"
        // Fila 3: L7..L9, U7..U9, R7..R9, D7..D9
        "|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|%s%s%s-%s%s%s-%s%s%s|\n"
        "|--------|--------|--------|--------|\n",
        
        /* -- FILA 1 -- */
        // L1..L3
        P(L1).c, P(L1).v, RESET,
        P(L2).c, P(L2).v, RESET,
        P(L3).c, P(L3).v, RESET,
        // U1..U3
        P(U1).c, P(U1).v, RESET,
        P(U2).c, P(U2).v, RESET,
        P(U3).c, P(U3).v, RESET,
        // R1..R3
        P(R1).c, P(R1).v, RESET,
        P(R2).c, P(R2).v, RESET,
        P(R3).c, P(R3).v, RESET,
        // D1..D3
        P(D1).c, P(D1).v, RESET,
        P(D2).c, P(D2).v, RESET,
        P(D3).c, P(D3).v, RESET,

        /* -- FILA 2 -- */
        // L4..L6
        P(L4).c, P(L4).v, RESET,
        P(L5).c, P(L5).v, RESET,
        P(L6).c, P(L6).v, RESET,
        // U4..U6
        P(U4).c, P(U4).v, RESET,
        P(U5).c, P(U5).v, RESET,
        P(U6).c, P(U6).v, RESET,
        // R4..R6
        P(R4).c, P(R4).v, RESET,
        P(R5).c, P(R5).v, RESET,
        P(R6).c, P(R6).v, RESET,
        // D4..D6
        P(D4).c, P(D4).v, RESET,
        P(D5).c, P(D5).v, RESET,
        P(D6).c, P(D6).v, RESET,

        /* -- FILA 3 -- */
        // L7..L9
        P(L7).c, P(L7).v, RESET,
        P(L8).c, P(L8).v, RESET,
        P(L9).c, P(L9).v, RESET,
        // U7..U9
        P(U7).c, P(U7).v, RESET,
        P(U8).c, P(U8).v, RESET,
        P(U9).c, P(U9).v, RESET,
        // R7..R9
        P(R7).c, P(R7).v, RESET,
        P(R8).c, P(R8).v, RESET,
        P(R9).c, P(R9).v, RESET,
        // D7..D9
        P(D7).c, P(D7).v, RESET,
        P(D8).c, P(D8).v, RESET,
        P(D9).c, P(D9).v, RESET
    );

    // Imprime la cara F (inferior)
    printf(
        "         |%s%s%s-%s%s%s-%s%s%s|\n"
        "         |%s%s%s-%s%s%s-%s%s%s|\n"
        "         |%s%s%s-%s%s%s-%s%s%s|\n"
        "         |--------|\n",
        /* 1era fila de la cara F */
        P(F1).c, P(F1).v, RESET, P(F2).c, P(F2).v, RESET, P(F3).c, P(F3).v, RESET,
        /* 2da fila de la cara F */
        P(F4).c, P(F4).v, RESET, P(F5).c, P(F5).v, RESET, P(F6).c, P(F6).v, RESET,
        /* 3era fila de la cara F */
        P(F7).c, P(F7).v, RESET, P(F8).c, P(F8).v, RESET, P(F9).c, P(F9).v, RESET
    );

    #undef P
}

/* Verifica si dos cubos son iguales */
bool compareCubes(RubikCube *c1, RubikCube *c2) {
    for (int i = 0; i < 54; i++) {
        if (strcmp(c1->pieces[i].v, c2->pieces[i].v) != 0 ||
            strcmp(c1->pieces[i].c, c2->pieces[i].c) != 0) {
            return false;
        }
    }
    return true;
}

#ifndef RUBIK_LIB
int main(int argc, char *argv[]) {
    RubikCube rubik, initialRubik;
    initializeRubik(&rubik);
    initializeRubik(&initialRubik); // Guardamos el cubo inicial

    bool suppressOutput = false;
    char *sequence = NULL;

    // Procesar argumentos
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "-s") == 0) {
            suppressOutput = true;
        } else {
            sequence = argv[i]; // La secuencia de movimientos
        }
    }

    // Realizar movimientos si se proporciona una secuencia
    if (sequence) {
        for (int j = 0; j < (int)strlen(sequence); j++) {
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
                    fprintf(stderr, "Movimiento desconocido: %c\n", sequence[j]);
                    return 1;
            }
        }
    }

    // Si se pasa -s, solo mostramos si está ordenado o no
    if (suppressOutput) {
        if (compareCubes(&rubik, &initialRubik)) {
            printf("Ordered\n");
        } else {
            printf("Unordered\n");
        }
    } else {
        // Mostramos el cubo inicial y el cubo tras los movimientos
        printf("Cubo inicial:\n");
        formatCli(&initialRubik);
        printf("\nCubo después de aplicar movimientos:\n");
        formatCli(&rubik);
        printf("\n");
    }

    return 0;
}
#endif

