#ifndef RUBIK_H
#define RUBIK_H

#include <stdbool.h>

// ==================== Colores ANSI ====================
// (puedes ajustarlos a tu gusto)
#define RESET   "\x1b[0m"
#define WHITE   "\x1b[37m"
#define YELLOW  "\x1b[33m"
#define GREEN   "\x1b[32m"
#define BLUE    "\x1b[34m"
#define MAGENTA "\x1b[35m"
#define RED     "\x1b[31m"

// ======================================================
// Asignaremos un entero por cada color de cara.
// Por convención: 
//   0 -> B (Back)
//   1 -> L (Left)
//   2 -> U (Up)
//   3 -> R (Right)
//   4 -> D (Down)
//   5 -> F (Front)
enum RubikColor {
    COLOR_B = 0,
    COLOR_L = 1,
    COLOR_U = 2,
    COLOR_R = 3,
    COLOR_D = 4,
    COLOR_F = 5
};

// Cada sticker/pieza tiene:
//   - un color (0..5)  -> indica de qué cara "nativa" es
//   - facePos (0..8)   -> posición dentro de esa cara
typedef struct {
    int color;    // 0..5
    int facePos;  // 0..8
} RubikPiece;

// Un cubo se compone de 54 stickers (6 caras * 9 = 54)
typedef struct {
    RubikPiece pieces[54];
} RubikCube;

#define B1  0
#define B2  1
#define B3  2
#define B4  3
#define B5  4
#define B6  5
#define B7  6
#define B8  7
#define B9  8

#define L1  9
#define L2  10
#define L3  11
#define L4  12
#define L5  13
#define L6  14
#define L7  15
#define L8  16
#define L9  17

#define U1  18
#define U2  19
#define U3  20
#define U4  21
#define U5  22
#define U6  23
#define U7  24
#define U8  25
#define U9  26

#define R1  27
#define R2  28
#define R3  29
#define R4  30
#define R5  31
#define R6  32
#define R7  33
#define R8  34
#define R9  35

#define D1  36
#define D2  37
#define D3  38
#define D4  39
#define D5  40
#define D6  41
#define D7  42
#define D8  43
#define D9  44

#define F1  45
#define F2  46
#define F3  47
#define F4  48
#define F5  49
#define F6  50
#define F7  51
#define F8  52
#define F9  53

// Mapeo color -> letra de la cara
static const char FACE_LETTER[6] = { 'B', 'L', 'U', 'R', 'D', 'F' };

// Mapeo color -> secuencia ANSI
static const char *COLOR_TO_ANSI[6] = {
    MAGENTA, // B
    BLUE,    // L
    YELLOW,  // U
    GREEN,   // R
    WHITE,   // D
    RED      // F
};

// Para imitar "P(B1).c" y "P(B1).v", definimos:
typedef struct {
    const char *c;  // secuencia ANSI
    const char *v;  // etiqueta tipo "B1"
} PrintPiece;

// ======================================================
// Prototipos de funciones
// ======================================================

// Inicializa el cubo en estado resuelto
void initializeRubik(RubikCube *rubik);

// Imprime el cubo en 6 filas de 9 stickers (forma simple)
void printRubik(const RubikCube *rubik);

// Compara 2 cubos (bit a bit) con memcmp
bool compareCubes(const RubikCube *c1, const RubikCube *c2);

// Movimientos
void moveU(RubikCube *rubik);
void moveUP(RubikCube *rubik);
void moveR(RubikCube *rubik);
void moveRP(RubikCube *rubik);
void moveF(RubikCube *rubik);
void moveFP(RubikCube *rubik);
void moveL(RubikCube *rubik);
void moveLP(RubikCube *rubik);
void moveD(RubikCube *rubik);
void moveDP(RubikCube *rubik);
void moveB(RubikCube *rubik);
void moveBP(RubikCube *rubik);

// Función auxiliar: el "sexy move" (R U R' U')
void sexyMove(RubikCube *rubik);

// Funcion applySequence
int applySequence(RubikCube *rubik, char *sequence);

static void printSticker(const RubikCube *rubik, int index);
// Declaración de formatCli
void formatCli(const RubikCube *rubik);

#endif // RUBIK_H
