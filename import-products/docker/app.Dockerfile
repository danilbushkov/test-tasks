FROM php:8.2-fpm-alpine

RUN apk add --no-cache \
    libzip-dev libpng-dev postgresql-dev

RUN docker-php-ext-install zip gd pdo pdo_pgsql

WORKDIR /var/www/app





