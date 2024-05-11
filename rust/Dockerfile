# Usar la imagen oficial de Rust como base
FROM rust:1.54.0

# Crear un directorio para el código de la aplicación
WORKDIR /usr/src/app

# Copiar el código de la aplicación al contenedor
COPY . .

# Compilar la aplicación
RUN cargo build --release

# Ejecutar la aplicación cuando el contenedor se inicie
# CMD ["cargo", "run"]
CMD ["sleep", "infinity"]