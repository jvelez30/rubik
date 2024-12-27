from multiprocessing import Pool, cpu_count
import itertools
import subprocess
import time

def validate_permutations_batch(batch):
    """Validate a batch of permutations using an external command."""
    try:
        results = []
        for perm in batch:
            command = ["rubik", "-s", perm]
            #print(f"Command: {command}")  # Debugging log
            process = subprocess.run(
                command, text=True, capture_output=True, check=True
            )
            #print(f"Output: {process.stdout.strip()}")  # Debugging log
            results.append(process.stdout.strip())
        return results
    except subprocess.CalledProcessError as e:
        print(f"Error: {e.stderr}")  # Debugging log
        return [f"Error: {e.stderr.strip()}" for _ in batch]

def main():
    # Configuración
    values = "UuFfRrLlDdBb"
    n = 4  # Cambia esto a la longitud deseada
    batch_size = 100  # Tamaño del lote

    # Generar permutaciones
    permutations = [''.join(p) for p in itertools.product(values, repeat=n)]

    # Medir tiempo de ejecución
    start_time = time.time()

    # Dividir permutaciones en lotes
    batches = [permutations[i:i + batch_size] for i in range(0, len(permutations), batch_size)]

    # Procesar lotes en paralelo
    with Pool(processes=cpu_count()) as pool:
        results = pool.map(validate_permutations_batch, batches)

    # Clasificar resultados
    ordered_results = []
    unordered_results = []

    for batch_results in results:
        for result in batch_results:
            if "Ordered" in result:
                ordered_results.append(result)
            elif "Unordered" in result:
                unordered_results.append(result)

    # Calcular tiempo total de ejecución
    end_time = time.time()
    elapsed_time = end_time - start_time

    # Mostrar resultados finales
    print("\n--- SUMMARY ---")
    print(f"Total Ordered (Valid): {len(ordered_results)}")
    print(f"Total Unordered (Invalid): {len(unordered_results)}")
    print(f"Execution Time: {elapsed_time:.2f} seconds")

if __name__ == "__main__":
    main()