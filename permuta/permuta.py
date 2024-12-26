import itertools
import requests
import time
from concurrent.futures import ThreadPoolExecutor

def generate_permutations(n, values):
    """Genera todas las permutaciones posibles de longitud n con los valores dados."""
    return [''.join(p) for p in itertools.product(values, repeat=n)]

def send_request(base_url, perm):
    """Envía una solicitud HTTP a la URL base con una permutación específica."""
    try:
        url = f"{base_url}{perm}"
        response = requests.get(url)
        status = response.text.strip()
        if "Ordered" in status:
            return perm, "Ordered"
        elif "Unordered" in status:
            return perm, "Unordered"
        else:
            return perm, f"Unexpected response: {status}"
    except Exception as e:
        return perm, f"Error: {e}"

def main():
    # Configuración
    base_url = "http://127.0.0.1:8095/rubik.php?s=1&m="
    values = "UuFfRrLlDdBb"
    n = 4  # Cambia esto a la longitud deseada
    max_workers = 10  # Número de hilos para paralelizar

    # Generar permutaciones
    permutations = generate_permutations(n, values)

    # Medir tiempo de ejecución
    start_time = time.time()

    # Enviar solicitudes en paralelo
    ordered_results = []
    unordered_results = []

    with ThreadPoolExecutor(max_workers=max_workers) as executor:
        futures = {executor.submit(send_request, base_url, perm): perm for perm in permutations}
        for future in futures:
            perm, status = future.result()
            if status == "Ordered":
                ordered_results.append(perm)
                #print(f"VALID (Ordered): {perm}")
            elif status == "Unordered":
                unordered_results.append(perm)
                #print(f"INVALID (Unordered): {perm}")
            else:
                print(f"Unexpected response for {perm}: {status}")

    # Calcular tiempo total de ejecución
    end_time = time.time()
    elapsed_time = end_time - start_time

    # Mostrar resultados finales
    print("\n--- SUMMARY ---")
    print(f"Total Ordered (Valid): {len(ordered_results)}")
    print(f"Total Unordered (Invalid): {len(unordered_results)}")
    print(f"Ordered Permutations: {ordered_results}")
    #print(f"Unordered Permutations: {unordered_results}")
    print(f"Execution Time: {elapsed_time:.2f} seconds")

if __name__ == "__main__":
    main()
