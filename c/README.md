# C Rubik

## Compile C File

```bash
gcc -o rubik rubik.c
```

## Execute Code

### Basic Execution

```bash
./rubik
```

This will initialize a Rubik's cube, apply no moves, and display the initial state of the cube.

### Passing a Sequence of Moves

```bash
./rubik "URuLf"
```

This will apply the specified sequence of moves to the cube:

- `U`: Rotate the upper face clockwise.
- `R`: Rotate the right face clockwise.
- `u`: Rotate the upper face counterclockwise.
- `L`: Rotate the left face clockwise.
- `f`: Rotate the front face counterclockwise.

It will display the cube's state before and after applying the moves.

### Using the `-s` Option

The `-s` (state) option suppresses the visual representation of the cube and instead determines if the cube is in its ordered (solved) state.

#### Example: Check the Cube's State Without Moves

```bash
./rubik -s
```

Output:

- `Ordered`: If the cube is in its initial (solved) state.
- `Unordered`: If the cube is not in its initial state.

#### Example: Check the Cube's State After Moves

```bash
./rubik -s "URuLf"
```

This will apply the specified sequence of moves to the cube and then output:

- `Ordered`: If the cube returns to its solved state after the moves.
- `Unordered`: If the cube remains unsolved after the moves.

## Permutations

### Compile app

```bash
gcc -O3 -DRUBIK_LIB -o rubik_permuta rubik.c rubik_permuta.c
```

### Run app

```bash
./rubik_permuta 6
```

### Expected output

```bash
--- SUMMARY ---
Total Ordered (Valid): 10464
Total Unordered (Invalid): 2975520
Execution Time: 31.31 seconds
```
