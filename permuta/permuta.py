import itertools
import aiohttp
import asyncio
import time

async def send_request(session, base_url, perm):
    """Envía una solicitud HTTP asincrónica con aiohttp."""
    try:
        url = f"{base_url}{perm}"
        async with session.get(url) as response:
            status = await response.text()
            if "Ordered" in status:
                return perm, "Ordered"
            elif "Unordered" in status:
                return perm, "Unordered"
            else:
                return perm, f"Unexpected response: {status}"
    except Exception as e:
        return perm, f"Error: {e}"

async def main():
    # Configuración
    base_url = "http://127.0.0.1:8096/rubik.php?s=1&m="
    values = "UuFfRrLlDdBb"
    n = 4  # Cambia esto a la longitud deseada

    # Generar permutaciones
    permutations = [''.join(p) for p in itertools.product(values, repeat=n)]

    # Medir tiempo de ejecución
    start_time = time.time()

    # Crear un cliente aiohttp
    async with aiohttp.ClientSession() as session:
        tasks = [send_request(session, base_url, perm) for perm in permutations]
        results = await asyncio.gather(*tasks)

    # Clasificar resultados
    ordered_results = [perm for perm, status in results if status == "Ordered"]
    unordered_results = [perm for perm, status in results if status == "Unordered"]

    # Calcular tiempo total de ejecución
    end_time = time.time()
    elapsed_time = end_time - start_time

    # Mostrar resultados finales
    print("\n--- SUMMARY ---")
    print(f"Total Ordered (Valid): {len(ordered_results)}")
    print(f"Total Unordered (Invalid): {len(unordered_results)}")
    print(f"Execution Time: {elapsed_time:.2f} seconds")

if __name__ == "__main__":
    asyncio.run(main())
