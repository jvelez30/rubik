#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "rubik.h"

void generate_permutations(const char *values, int n, char *current, int pos, RubikCube *initial, int *ordered, int *unordered) {
    if (pos == n) {
        current[pos] = '\0';
        RubikCube rubik;
        initializeRubik(&rubik);

        // Aplica la secuencia de movimientos
        for (int i = 0; i < n; i++) {
            switch (current[i]) {
                case 'U': moveU(&rubik); break;
                case 'u': moveUP(&rubik); break;
                case 'R': moveR(&rubik); break;
                case 'r': moveRP(&rubik); break;
                case 'L': moveL(&rubik); break;
                case 'l': moveLP(&rubik); break;
                case 'F': moveF(&rubik); break;
                case 'f': moveFP(&rubik); break;
                case 'D': moveD(&rubik); break;
                case 'd': moveDP(&rubik); break;
                case 'B': moveB(&rubik); break;
                case 'b': moveBP(&rubik); break;
            }
        }

        // Compara el cubo resultante con el inicial
        if (compareCubes(&rubik, initial)) {
            (*ordered)++;
        } else {
            (*unordered)++;
        }
        return;
    }

    for (int i = 0; values[i] != '\0'; i++) {
        current[pos] = values[i];
        generate_permutations(values, n, current, pos + 1, initial, ordered, unordered);
    }
}

int main(int argc, char *argv[]) {
    // Configuración
    const char *values = "UuFfRrLlDdBb";
    int n = 4;  // Cambiar según necesidad

    if (argc > 1) {
        n = atoi(argv[1]);
    }

    RubikCube initialRubik;
    initializeRubik(&initialRubik);

    int ordered = 0, unordered = 0;
    char *current = malloc(n + 1);

    if (current == NULL) {
        fprintf(stderr, "Error allocating memory.\n");
        return 1;
    }

    // Medir tiempo
    clock_t start = clock();

    // Generar permutaciones y validar
    generate_permutations(values, n, current, 0, &initialRubik, &ordered, &unordered);

    clock_t end = clock();
    double elapsed_time = (double)(end - start) / CLOCKS_PER_SEC;

    // Mostrar resultados
    printf("\n--- SUMMARY ---\n");
    printf("Total Ordered (Valid): %d\n", ordered);
    printf("Total Unordered (Invalid): %d\n", unordered);
    printf("Execution Time: %.2f seconds\n", elapsed_time);

    free(current);
    return 0;
}
