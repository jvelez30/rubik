#ifndef RUBIK_H
#define RUBIK_H

#include <stdbool.h>
#include <stdio.h>

// Colores
#define RED "\e[31m"
#define GREEN "\e[32m"
#define YELLOW "\e[33m"
#define BLUE "\e[34m"
#define MAGENTA "\e[95m"
#define WHITE "\e[97m"
#define RESET "\e[0m"

// Estructuras
typedef struct {
    char v[3];
    char c[10];
} RubikPiece;

typedef struct {
    RubikPiece pieces[54];
} RubikCube;

// Declaración de funciones
void initializeRubik(RubikCube *rubik);
bool compareCubes(RubikCube *c1, RubikCube *c2);
void moveU(RubikCube *rubik);
void moveUP(RubikCube *rubik);
void moveR(RubikCube *rubik);
void moveRP(RubikCube *rubik);
void moveL(RubikCube *rubik);
void moveLP(RubikCube *rubik);
void moveF(RubikCube *rubik);
void moveFP(RubikCube *rubik);
void moveD(RubikCube *rubik);
void moveDP(RubikCube *rubik);
void moveB(RubikCube *rubik);
void moveBP(RubikCube *rubik);

#endif
