FROM php:8.1-fpm

# Instalar Nginx
RUN apt-get update && apt-get install -y nginx && apt-get clean

# Copiar configuraciones de Nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf

# Establecer el directorio de trabajo
WORKDIR /app

# Exponer el puerto
EXPOSE 8095

# Comando para iniciar PHP-FPM y Nginx
CMD ["sh", "-c", "php-fpm & nginx -g 'daemon off;'"]
