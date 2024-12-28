#include <stdio.h>
#include <string.h>   // Para memcpy, memcmp
#include <stdlib.h>   // Para atoi (si lo quisieras)
#include "rubik.h"

// ---------------------------------------------------------------------------
// Macros para indexar rápidamente: faceIndex(face, pos)
#define faceIndex(face, pos) ((face) * 9 + (pos))

// ---------------------------------------------------------------------------
// Inicializa el cubo en estado resuelto.
// Asignamos a cada sticker su color = face, y facePos = 0..8
void initializeRubik(RubikCube *rubik) {
    // Rellenar las 6 caras
    for (int face = 0; face < 6; face++) {
        for (int pos = 0; pos < 9; pos++) {
            int idx = faceIndex(face, pos);
            rubik->pieces[idx].color   = face;
            rubik->pieces[idx].facePos = pos;
        }
    }
}

// ---------------------------------------------------------------------------
// Auxiliar para imprimir 54 stickers en 6 filas de 9
void printRubik(const RubikCube *rubik) {
    for (int i = 0; i < 54; i++) {
        int c = rubik->pieces[i].color; // 0..5
        // facePos no lo usamos aquí, pero podrías mostrarlo si quieres

        // Imprimimos color + [facePos] + RESET:
        printf("%s[%d]%s ", COLOR_TO_ANSI[c], rubik->pieces[i].facePos, RESET);

        if ((i + 1) % 9 == 0) {
            printf("\n");
        }
    }
}

// ---------------------------------------------------------------------------
// Compara 2 cubos byte a byte con memcmp
bool compareCubes(const RubikCube *c1, const RubikCube *c2) {
    return (memcmp(c1->pieces, c2->pieces, sizeof(c1->pieces)) == 0);
}

// ---------------------------------------------------------------------------
// Permutaciones: En esta versión, cada sticker se mueve a otra posición
//                (índice del array [0..53]). Hacemos un array temp[54]
//                y copiamos en un solo paso.
//
// applyPermutation() recibe:
//   - oldIndices[]: las posiciones que se van a "actualizar" en rubik->pieces
//   - newIndices[]: de dónde se copian los datos de temp
//   - count: cuántos stickers permutar
static void applyPermutation(RubikCube *rubik, const int *oldIndices,
                             const int *newIndices, int count)
{
    RubikCube backup;
    memcpy(&backup, rubik, sizeof(RubikCube));

    for (int i = 0; i < count; i++) {
        int oldIdx = oldIndices[i];
        int newIdx = newIndices[i];
        // Copiamos color y facePos
        rubik->pieces[oldIdx].color   = backup.pieces[newIdx].color;
        rubik->pieces[oldIdx].facePos = backup.pieces[newIdx].facePos;
    }
}

// ---------------------------------------------------------------------------
// Definimos macros para las caras, a fin de usar en las permutaciones
// B=0, L=1, U=2, R=3, D=4, F=5 (como en enum RubikColor)
#define B 0
#define L 1
#define U 2
#define R 3
#define D 4
#define F 5

// Rotación U (Up) horario
void moveU(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(U,0), faceIndex(U,1), faceIndex(U,2),
        faceIndex(U,5), faceIndex(U,8), faceIndex(U,7),
        faceIndex(U,6), faceIndex(U,3),
        faceIndex(R,0), faceIndex(R,3), faceIndex(R,6),
        faceIndex(F,0), faceIndex(F,1), faceIndex(F,2),
        faceIndex(L,2), faceIndex(L,5), faceIndex(L,8),
        faceIndex(B,6), faceIndex(B,7), faceIndex(B,8)
    };
    static const int newIdx[20] = {
        faceIndex(U,6), faceIndex(U,3), faceIndex(U,0),
        faceIndex(U,1), faceIndex(U,2), faceIndex(U,5),
        faceIndex(U,8), faceIndex(U,7),
        faceIndex(B,6), faceIndex(B,7), faceIndex(B,8),
        faceIndex(R,6), faceIndex(R,3), faceIndex(R,0),
        faceIndex(F,0), faceIndex(F,1), faceIndex(F,2),
        faceIndex(L,8), faceIndex(L,5), faceIndex(L,2)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// U' (Up inverso)
void moveUP(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(U,0), faceIndex(U,1), faceIndex(U,2),
        faceIndex(U,5), faceIndex(U,8), faceIndex(U,7),
        faceIndex(U,6), faceIndex(U,3),
        faceIndex(R,0), faceIndex(R,3), faceIndex(R,6),
        faceIndex(F,0), faceIndex(F,1), faceIndex(F,2),
        faceIndex(L,2), faceIndex(L,5), faceIndex(L,8),
        faceIndex(B,6), faceIndex(B,7), faceIndex(B,8)
    };
    static const int newIdx[20] = {
        faceIndex(U,2), faceIndex(U,5), faceIndex(U,8),
        faceIndex(U,7), faceIndex(U,6), faceIndex(U,3),
        faceIndex(U,0), faceIndex(U,1),
        faceIndex(F,2), faceIndex(F,1), faceIndex(F,0),
        faceIndex(L,2), faceIndex(L,5), faceIndex(L,8),
        faceIndex(B,8), faceIndex(B,7), faceIndex(B,6),
        faceIndex(R,0), faceIndex(R,3), faceIndex(R,6)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// ---------------------------------------------------------------------------
// R (Right) horario
void moveR(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(R,0), faceIndex(R,1), faceIndex(R,2),
        faceIndex(R,5), faceIndex(R,8), faceIndex(R,7),
        faceIndex(R,6), faceIndex(R,3),
        faceIndex(B,2), faceIndex(B,5), faceIndex(B,8),
        faceIndex(D,6), faceIndex(D,3), faceIndex(D,0),
        faceIndex(F,2), faceIndex(F,5), faceIndex(F,8),
        faceIndex(U,2), faceIndex(U,5), faceIndex(U,8)
    };
    static const int newIdx[20] = {
        faceIndex(R,6), faceIndex(R,3), faceIndex(R,0),
        faceIndex(R,1), faceIndex(R,2), faceIndex(R,5),
        faceIndex(R,8), faceIndex(R,7),
        faceIndex(U,2), faceIndex(U,5), faceIndex(U,8),
        faceIndex(B,2), faceIndex(B,5), faceIndex(B,8),
        faceIndex(D,6), faceIndex(D,3), faceIndex(D,0),
        faceIndex(F,2), faceIndex(F,5), faceIndex(F,8)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// R' (Right inverso)
void moveRP(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(R,0), faceIndex(R,1), faceIndex(R,2),
        faceIndex(R,5), faceIndex(R,8), faceIndex(R,7),
        faceIndex(R,6), faceIndex(R,3),
        faceIndex(B,2), faceIndex(B,5), faceIndex(B,8),
        faceIndex(D,6), faceIndex(D,3), faceIndex(D,0),
        faceIndex(F,2), faceIndex(F,5), faceIndex(F,8),
        faceIndex(U,2), faceIndex(U,5), faceIndex(U,8)
    };
    static const int newIdx[20] = {
        faceIndex(R,2), faceIndex(R,5), faceIndex(R,8),
        faceIndex(R,7), faceIndex(R,6), faceIndex(R,3),
        faceIndex(R,0), faceIndex(R,1),
        faceIndex(D,6), faceIndex(D,3), faceIndex(D,0),
        faceIndex(F,2), faceIndex(F,5), faceIndex(F,8),
        faceIndex(U,2), faceIndex(U,5), faceIndex(U,8),
        faceIndex(B,2), faceIndex(B,5), faceIndex(B,8)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// ---------------------------------------------------------------------------
// F (Front) horario
void moveF(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(F,0), faceIndex(F,1), faceIndex(F,2),
        faceIndex(F,5), faceIndex(F,8), faceIndex(F,7),
        faceIndex(F,6), faceIndex(F,3),
        faceIndex(L,6), faceIndex(L,7), faceIndex(L,8),
        faceIndex(U,6), faceIndex(U,7), faceIndex(U,8),
        faceIndex(R,6), faceIndex(R,7), faceIndex(R,8),
        faceIndex(D,6), faceIndex(D,7), faceIndex(D,8)
    };
    static const int newIdx[20] = {
        faceIndex(F,6), faceIndex(F,3), faceIndex(F,0),
        faceIndex(F,1), faceIndex(F,2), faceIndex(F,5),
        faceIndex(F,8), faceIndex(F,7),
        faceIndex(D,6), faceIndex(D,7), faceIndex(D,8),
        faceIndex(L,6), faceIndex(L,7), faceIndex(L,8),
        faceIndex(U,6), faceIndex(U,7), faceIndex(U,8),
        faceIndex(R,6), faceIndex(R,7), faceIndex(R,8)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// F' (Front inverso)
void moveFP(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(F,0), faceIndex(F,1), faceIndex(F,2),
        faceIndex(F,5), faceIndex(F,8), faceIndex(F,7),
        faceIndex(F,6), faceIndex(F,3),
        faceIndex(L,6), faceIndex(L,7), faceIndex(L,8),
        faceIndex(U,6), faceIndex(U,7), faceIndex(U,8),
        faceIndex(R,6), faceIndex(R,7), faceIndex(R,8),
        faceIndex(D,6), faceIndex(D,7), faceIndex(D,8)
    };
    static const int newIdx[20] = {
        faceIndex(F,2), faceIndex(F,5), faceIndex(F,8),
        faceIndex(F,7), faceIndex(F,6), faceIndex(F,3),
        faceIndex(F,0), faceIndex(F,1),
        faceIndex(U,6), faceIndex(U,7), faceIndex(U,8),
        faceIndex(R,6), faceIndex(R,7), faceIndex(R,8),
        faceIndex(D,6), faceIndex(D,7), faceIndex(D,8),
        faceIndex(L,6), faceIndex(L,7), faceIndex(L,8)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// ---------------------------------------------------------------------------
// L (Left) horario
void moveL(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(L,0), faceIndex(L,1), faceIndex(L,2),
        faceIndex(L,5), faceIndex(L,8), faceIndex(L,6),
        faceIndex(L,7), faceIndex(L,3),
        faceIndex(B,0), faceIndex(B,3), faceIndex(B,6),
        faceIndex(U,0), faceIndex(U,3), faceIndex(U,6),
        faceIndex(F,0), faceIndex(F,3), faceIndex(F,6),
        faceIndex(D,2), faceIndex(D,5), faceIndex(D,8)
    };
    static const int newIdx[20] = {
        faceIndex(L,6), faceIndex(L,3), faceIndex(L,0),
        faceIndex(L,1), faceIndex(L,2), faceIndex(L,8),
        faceIndex(L,5), faceIndex(L,7),
        faceIndex(D,8), faceIndex(D,5), faceIndex(D,2),
        faceIndex(B,0), faceIndex(B,3), faceIndex(B,6),
        faceIndex(U,0), faceIndex(U,3), faceIndex(U,6),
        faceIndex(F,6), faceIndex(F,3), faceIndex(F,0)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// L' (Left inverso)
void moveLP(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(L,0), faceIndex(L,1), faceIndex(L,2),
        faceIndex(L,5), faceIndex(L,8), faceIndex(L,7),
        faceIndex(L,6), faceIndex(L,3),
        faceIndex(B,0), faceIndex(B,3), faceIndex(B,6),
        faceIndex(U,0), faceIndex(U,3), faceIndex(U,6),
        faceIndex(F,0), faceIndex(F,3), faceIndex(F,6),
        faceIndex(D,2), faceIndex(D,5), faceIndex(D,8)
    };
    static const int newIdx[20] = {
        faceIndex(L,2), faceIndex(L,5), faceIndex(L,8),
        faceIndex(L,7), faceIndex(L,6), faceIndex(L,3),
        faceIndex(L,0), faceIndex(L,1),
        faceIndex(U,0), faceIndex(U,3), faceIndex(U,6),
        faceIndex(F,0), faceIndex(F,3), faceIndex(F,6),
        faceIndex(D,8), faceIndex(D,5), faceIndex(D,2),
        faceIndex(B,6), faceIndex(B,3), faceIndex(B,0)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// ---------------------------------------------------------------------------
// D (Down) horario
void moveD(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(D,0), faceIndex(D,1), faceIndex(D,2),
        faceIndex(D,5), faceIndex(D,8), faceIndex(D,7),
        faceIndex(D,6), faceIndex(D,3),
        faceIndex(B,0), faceIndex(B,1), faceIndex(B,2),
        faceIndex(L,0), faceIndex(L,3), faceIndex(L,6),
        faceIndex(R,2), faceIndex(R,5), faceIndex(R,8),
        faceIndex(F,6), faceIndex(F,7), faceIndex(F,8)
    };
    static const int newIdx[20] = {
        faceIndex(D,6), faceIndex(D,3), faceIndex(D,0),
        faceIndex(D,1), faceIndex(D,2), faceIndex(D,5),
        faceIndex(D,8), faceIndex(D,7),
        faceIndex(R,2), faceIndex(R,5), faceIndex(R,8),
        faceIndex(B,2), faceIndex(B,1), faceIndex(B,0),
        faceIndex(F,8), faceIndex(F,7), faceIndex(F,6),
        faceIndex(L,0), faceIndex(L,3), faceIndex(L,6)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// D' (Down inverso)
void moveDP(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(D,0), faceIndex(D,1), faceIndex(D,2),
        faceIndex(D,5), faceIndex(D,8), faceIndex(D,7),
        faceIndex(D,6), faceIndex(D,3),
        faceIndex(B,0), faceIndex(B,1), faceIndex(B,2),
        faceIndex(L,0), faceIndex(L,3), faceIndex(L,6),
        faceIndex(R,2), faceIndex(R,5), faceIndex(R,8),
        faceIndex(F,6), faceIndex(F,7), faceIndex(F,8)
    };
    static const int newIdx[20] = {
        faceIndex(D,2), faceIndex(D,5), faceIndex(D,8),
        faceIndex(D,7), faceIndex(D,6), faceIndex(D,3),
        faceIndex(D,0), faceIndex(D,1),
        faceIndex(L,6), faceIndex(L,3), faceIndex(L,0),
        faceIndex(F,6), faceIndex(F,7), faceIndex(F,8),
        faceIndex(B,0), faceIndex(B,1), faceIndex(B,2),
        faceIndex(R,8), faceIndex(R,5), faceIndex(R,2)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// ---------------------------------------------------------------------------
// B (Back) horario
void moveB(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(B,0), faceIndex(B,1), faceIndex(B,2),
        faceIndex(B,5), faceIndex(B,8), faceIndex(B,7),
        faceIndex(B,6), faceIndex(B,3),
        faceIndex(L,0), faceIndex(L,1), faceIndex(L,2),
        faceIndex(U,0), faceIndex(U,1), faceIndex(U,2),
        faceIndex(R,0), faceIndex(R,1), faceIndex(R,2),
        faceIndex(D,0), faceIndex(D,1), faceIndex(D,2)
    };
    static const int newIdx[20] = {
        faceIndex(B,6), faceIndex(B,3), faceIndex(B,0),
        faceIndex(B,1), faceIndex(B,2), faceIndex(B,5),
        faceIndex(B,8), faceIndex(B,7),
        faceIndex(U,0), faceIndex(U,1), faceIndex(U,2),
        faceIndex(R,0), faceIndex(R,1), faceIndex(R,2),
        faceIndex(D,0), faceIndex(D,1), faceIndex(D,2),
        faceIndex(L,0), faceIndex(L,1), faceIndex(L,2)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// B' (Back inverso)
void moveBP(RubikCube *rubik) {
    static const int oldIdx[20] = {
        faceIndex(B,0), faceIndex(B,1), faceIndex(B,2),
        faceIndex(B,5), faceIndex(B,8), faceIndex(B,7),
        faceIndex(B,6), faceIndex(B,3),
        faceIndex(L,0), faceIndex(L,1), faceIndex(L,2),
        faceIndex(U,0), faceIndex(U,1), faceIndex(U,2),
        faceIndex(R,0), faceIndex(R,1), faceIndex(R,2),
        faceIndex(D,0), faceIndex(D,1), faceIndex(D,2)
    };
    static const int newIdx[20] = {
        faceIndex(B,2), faceIndex(B,5), faceIndex(B,8),
        faceIndex(B,7), faceIndex(B,6), faceIndex(B,3),
        faceIndex(B,0), faceIndex(B,1),
        faceIndex(D,0), faceIndex(D,1), faceIndex(D,2),
        faceIndex(L,0), faceIndex(L,1), faceIndex(L,2),
        faceIndex(U,0), faceIndex(U,1), faceIndex(U,2),
        faceIndex(R,0), faceIndex(R,1), faceIndex(R,2)
    };
    applyPermutation(rubik, oldIdx, newIdx, 20);
}

// ---------------------------------------------------------------------------
// "sexy move": R U R' U'
void sexyMove(RubikCube *rubik) {
    moveR(rubik);
    moveU(rubik);
    moveRP(rubik);
    moveUP(rubik);
}

/*
   Devuelve la etiqueta (por ejemplo, "B1" o "R9") a partir de color y facePos.
   Usamos un buffer estático circular para poder devolver un puntero válido
   aunque se llame varias veces en un mismo printf.
*/
static int gCallCount = 0;

static const char *getStickerLabel(const RubikCube *rubik, int index) {
    static char buffer[54][4];

    int idx = gCallCount % 54;
    gCallCount++;

    int colorVal = rubik->pieces[index].color; 
    int posVal   = rubik->pieces[index].facePos + 1;

    char faceLetter = FACE_LETTER[colorVal];
    snprintf(buffer[idx], sizeof(buffer[idx]), "%c%d", faceLetter, posVal);
    return buffer[idx];
}

// Imprime un sticker (color + etiqueta + RESET)
static void printSticker(const RubikCube *rubik, int index) {
    // Tomamos el color y la posición desde la estructura
    int colorVal  = rubik->pieces[index].color;   // 0..5
    int facePos   = rubik->pieces[index].facePos; // 0..8

    // Construimos la etiqueta en una variable local
    char faceLetter = FACE_LETTER[colorVal];      // 'B','L','U','R','D','F'
    char label[4];                                // Ej. "B1", "R9", etc.
    snprintf(label, sizeof(label), "%c%d", faceLetter, facePos + 1);

    // Imprimimos color ANSI + etiqueta + RESET
    printf("%s%s%s", COLOR_TO_ANSI[colorVal], label, RESET);
}

void formatCli(const RubikCube *rubik)
{
    // Ejemplo: cara B arriba
    printf("         |--------|\n");
    printf("         |");
    // B1, B2, B3
    printSticker(rubik, B1); printf("-");
    printSticker(rubik, B2); printf("-");
    printSticker(rubik, B3); printf("|\n");
    
    printf("         |");
    // B4, B5, B6
    printSticker(rubik, B4); printf("-");
    printSticker(rubik, B5); printf("-");
    printSticker(rubik, B6); printf("|\n");
    
    printf("         |");
    // B7, B8, B9
    printSticker(rubik, B7); printf("-");
    printSticker(rubik, B8); printf("-");
    printSticker(rubik, B9); printf("|\n");
    printf("|--------|--------|--------|--------|\n");
    
    // Ahora la tira del medio: L, U, R, D...
    //  (fila 1): L1-L2-L3, U1-U2-U3, R1-R2-R3, D1-D2-D3
    printf("|");
    printSticker(rubik, L1); printf("-");
    printSticker(rubik, L2); printf("-");
    printSticker(rubik, L3); printf("|");
    printSticker(rubik, U1); printf("-");
    printSticker(rubik, U2); printf("-");
    printSticker(rubik, U3); printf("|");
    printSticker(rubik, R1); printf("-");
    printSticker(rubik, R2); printf("-");
    printSticker(rubik, R3); printf("|");
    printSticker(rubik, D1); printf("-");
    printSticker(rubik, D2); printf("-");
    printSticker(rubik, D3); printf("|\n");

    //  (fila 2): L4-L5-L6, U4-U5-U6, R4-R5-R6, D4-D5-D6
    printf("|");
    printSticker(rubik, L4); printf("-");
    printSticker(rubik, L5); printf("-");
    printSticker(rubik, L6); printf("|");
    printSticker(rubik, U4); printf("-");
    printSticker(rubik, U5); printf("-");
    printSticker(rubik, U6); printf("|");
    printSticker(rubik, R4); printf("-");
    printSticker(rubik, R5); printf("-");
    printSticker(rubik, R6); printf("|");
    printSticker(rubik, D4); printf("-");
    printSticker(rubik, D5); printf("-");
    printSticker(rubik, D6); printf("|\n");

    //  (fila 3): L7-L8-L9, U7-U8-U9, R7-R8-R9, D7-D8-D9
    printf("|");
    printSticker(rubik, L7); printf("-");
    printSticker(rubik, L8); printf("-");
    printSticker(rubik, L9); printf("|");
    printSticker(rubik, U7); printf("-");
    printSticker(rubik, U8); printf("-");
    printSticker(rubik, U9); printf("|");
    printSticker(rubik, R7); printf("-");
    printSticker(rubik, R8); printf("-");
    printSticker(rubik, R9); printf("|");
    printSticker(rubik, D7); printf("-");
    printSticker(rubik, D8); printf("-");
    printSticker(rubik, D9); printf("|\n");
    printf("|--------|--------|--------|--------|\n");

    // Finalmente cara F (inferior)
    printf("         |");
    printSticker(rubik, F1); printf("-");
    printSticker(rubik, F2); printf("-");
    printSticker(rubik, F3); printf("|\n");
    printf("         |");
    printSticker(rubik, F4); printf("-");
    printSticker(rubik, F5); printf("-");
    printSticker(rubik, F6); printf("|\n");
    printf("         |");
    printSticker(rubik, F7); printf("-");
    printSticker(rubik, F8); printf("-");
    printSticker(rubik, F9); printf("|\n");
    printf("         |--------|\n");
}

// ---------------------------------------------------------------------------
// MAIN de prueba (desactivado si compilas con -DRUBIK_LIB, por ejemplo)
#ifndef RUBIK_LIB
int main(int argc, char *argv[]) {
    RubikCube rubik, initial;
    initializeRubik(&rubik);
    initializeRubik(&initial);

    bool suppressOutput = false;
    char *sequence = NULL;

    // Procesar argumentos
    // Ej: ./rubik -s RUru
    for (int i = 1; i < argc; i++) {
        if (strcmp(argv[i], "-s") == 0) {
            suppressOutput = true;
        } else {
            sequence = argv[i]; // secuencia de movimientos
        }
    }

    // Si tenemos secuencia, ejecutar
    if (sequence) {
        for (int j = 0; j < (int)strlen(sequence); j++) {
            switch (sequence[j]) {
                case 'U': moveU(&rubik);  break;
                case 'u': moveUP(&rubik); break;
                case 'R': moveR(&rubik);  break;
                case 'r': moveRP(&rubik); break;
                case 'F': moveF(&rubik);  break;
                case 'f': moveFP(&rubik); break;
                case 'L': moveL(&rubik);  break;
                case 'l': moveLP(&rubik); break;
                case 'D': moveD(&rubik);  break;
                case 'd': moveDP(&rubik); break;
                case 'B': moveB(&rubik);  break;
                case 'b': moveBP(&rubik); break;
                default:
                    fprintf(stderr, "Movimiento desconocido: %c\n", sequence[j]);
                    return 1;
            }
        }
    }

    // Salida
    if (suppressOutput) {
        // Checar si rubik == initial
        if (compareCubes(&rubik, &initial)) {
            printf("Ordered\n");
        } else {
            printf("Unordered\n");
        }
    } else {
        printf("Cubo inicial:\n");
        formatCli(&initial);
        printf("\nCubo después de movimientos:\n");
        formatCli(&rubik);
        printf("\n");

        if (compareCubes(&rubik, &initial)) {
            printf("El cubo volvió al estado inicial (Ordered)\n");
        } else {
            printf("El cubo quedó diferente (Unordered)\n");
        }
    }

    return 0;
}
#endif