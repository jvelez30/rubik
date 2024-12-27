# Permutation Validator

## Description

This application validates permutations generated from a set of characters based on predefined rules for a Rubik's cube. The main goal is to classify permutations as `Ordered` or `Unordered` according to predefined logic, previously handled by a PHP server, now implemented locally in Python.

## Features

- Efficient generation of permutations.
- Validation of permutations using replicated server logic.
- Batch processing to optimize memory and execution time.
- Final report with the count of ordered and unordered permutations, along with the total execution time.

## Requirements

- Python 3.7 or higher.
- Standard libraries (no additional dependencies required).

## Installation

1. Clone this repository:

   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. Ensure Python is installed:

   ```bash
   python --version
   ```

## Usage

1. Run the application from the command line:

   ```bash
   python <script_name>.py
   ```

2. Configure the values in the script as needed:
   - `values`: Characters to generate permutations (default: `"UuFfRrLlDdBb"`).
   - `n`: Length of the permutations.
   - `batch_size`: Size of each batch for processing permutations.

3. The program will generate all possible permutations, validate them, and display a summary at the end:
   - Total ordered permutations (`Ordered`).
   - Total unordered permutations (`Unordered`).
   - Total execution time.

## Example

For a set of values `"UuFfRr"` and length `n = 3`:

```bash
python permutation_validator.py
```

Expected output:

```bash
Batch processed: 12 Ordered, 36 Unordered.
...
--- FINAL SUMMARY ---
Total Ordered (Valid): 120
Total Unordered (Invalid): 2160
Execution Time: 2.34 seconds
```

## Customization

You can modify the validation logic in the `is_ordered` function in the main file to adapt it to different or more complex rules.

## Notes

- To improve performance on larger datasets, consider adjusting the batch size (`batch_size`).
- If you plan to use this logic in another context, ensure the rules in `is_ordered` align with your requirements.

## Contributions

Contributions are welcome. If you encounter an issue or have an improvement, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
