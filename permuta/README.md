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
- A web server running on `http://127.0.0.1:8096` with the endpoint `rubik.php` to validate permutations.

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

3. Set up a web server:
   - Place the `rubik.php` file in your web server's root directory.
   - Ensure the server is accessible at `http://127.0.0.1:8096`.

## Usage

1. Run the application from the command line:

   ```bash
   python <script_name>.py
   ```

2. Configure the values in the script as needed:
   - `base_url`: The URL of the server handling permutation validation (default: `http://127.0.0.1:8096/rubik.php?s=1&m=`).
   - `values`: Characters to generate permutations (default: `"UuFfRrLlDdBb"`).
   - `n`: Length of the permutations.
   - `batch_size`: Size of each batch for processing permutations.

3. The program will generate all possible permutations, validate them against the server, and display a summary at the end:
   - Total ordered permutations (`Ordered`).
   - Total unordered permutations (`Unordered`).
   - Total execution time.

## Example

For a set of values `"UuFfRr"` and length `n = 4`:

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

You can modify the validation logic in the `rubik.php` file on the server to adapt it to different or more complex rules.

## Notes

- To improve performance on larger datasets, consider adjusting the batch size (`batch_size`).
- Ensure the web server is running and accessible before executing the script.

## Contributions

Contributions are welcome. If you encounter an issue or have an improvement, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).
