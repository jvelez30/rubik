#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include "rubik.h"

void generate_permutations(const char *values, int n, char *current, int pos, RubikCube *initial, int *ordered, int *unordered, long long *counter) {
    if (pos == n) {
        current[pos] = '\0';
        (*counter)++; // Incrementar el contador de permutaciones generadas

        RubikCube rubik;
        initializeRubik(&rubik);

        // Aplica la secuencia de movimientos
        applySequence(&rubik, current);

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
        generate_permutations(values, n, current, pos + 1, initial, ordered, unordered, counter);
    }
}

int main(int argc, char *argv[]) {
    const char *values = "UuFfRrLlDdBb";
    int n = 4;  // Longitud de las permutaciones

    if (argc > 1) {
        n = atoi(argv[1]);
    }

    RubikCube initialRubik;
    initializeRubik(&initialRubik);

    int ordered = 0, unordered = 0;
    long long counter = 0; // Contador para rastrear todas las permutaciones generadas
    char *current = malloc(n + 1);

    if (current == NULL) {
        fprintf(stderr, "Error allocating memory.\n");
        return 1;
    }

    clock_t start = clock();

    generate_permutations(values, n, current, 0, &initialRubik, &ordered, &unordered, &counter);

    clock_t end = clock();
    double elapsed_time = (double)(end - start) / CLOCKS_PER_SEC;

    // Mostrar resultados
    printf("\n--- SUMMARY ---\n");
    printf("Total Permutations Generated: %lld\n", counter);
    printf("Total Ordered (Valid): %d\n", ordered);
    printf("Total Unordered (Invalid): %d\n", unordered);
    printf("Execution Time: %.2f seconds\n", elapsed_time);

    free(current);
    return 0;
}
