import itertools
import requests

def generate_permutations(n, values):
    """Genera todas las permutaciones posibles de longitud n con los valores dados."""
    return [''.join(p) for p in itertools.product(values, repeat=n)]

def send_requests(base_url, permutations):
    """Envía solicitudes HTTP a la URL base con cada permutación."""
    ordered_results = []
    unordered_results = []
    
    for perm in permutations:
        try:
            url = f"{base_url}{perm}"
            response = requests.get(url)
            status = response.text.strip()
            
            if "Ordered" in status:
                ordered_results.append(perm)
                print(f"VALID (Ordered): {perm}")
            elif "Unordered" in status:
                unordered_results.append(perm)
                print(f"INVALID (Unordered): {perm}")
            else:
                print(f"Unexpected response for {perm}: {status}")
        except Exception as e:
            print(f"Error with URL {url}: {e}")
    
    return ordered_results, unordered_results

def main():
    # Configuración
    base_url = "http://127.0.0.1:8095/rubik.php?s=1&m="
    values = "UuFfRrLlDdBb"
    n = 4  # Cambia esto a la longitud deseada

    # Generar permutaciones
    permutations = generate_permutations(n, values)

    # Enviar solicitudes y clasificar resultados
    ordered, unordered = send_requests(base_url, permutations)

    # Mostrar resultados finales
    print("\n--- SUMMARY ---")
    print(f"Total Ordered (Valid): {len(ordered)}")
    print(f"Total Unordered (Invalid): {len(unordered)}")
    print(f"Ordered Permutations: {ordered}")
    print(f"Unordered Permutations: {unordered}")

if __name__ == "__main__":
    main()
