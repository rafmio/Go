#!/bin/bash

# Массив маршрутов
routes=("/")
routes+=("/about")
routes+=("/contact")

# Количество итераций
iterations=3

for ((i = 0; i < iterations; i++)); do
    for route in "${routes[@]}"; do
        curl -X GET "http://localhost:8080${route}"
        echo ""  # Добавляет пустую строку после каждого запроса
        sleep 1  # Пауза между запросами
    done
done
