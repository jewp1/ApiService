# REST API для управления задачами, написанный на Go с использованием фреймворка Fiber.

## CRUD операции для задач:
  - Создание задачи (POST /task Content-Type: application/json)
```
{
  "title": "Название",
  "description": "Описание"
}
```
  - Получение списка задач (GET /tasks)
  - Получение задачи по ID (GET /task/id)
  - Обновление задачи (PUT /task/id
Content-Type: application/json)
```
{
  "title": "Название",
  "description": "Описание"
}
```
  - Удаление задачи (DELETE /task/id)

## Технологии
- Fiber - веб-фреймворк
- godotenv - загрузка .env файлов
- zap - логгер

## Запуск
```bash
go run cmd/main.go
```
Сервис будет доступен по адресу: http://localhost:8080

## Структура 
- cmd/ Точка входа приложения
- config/ Обработка конфига
- dto/ Http ответы 
- logger/ Логгер
- repo/ Работа с хранилицем
- service/ Бизнес логика
- go.mod # Зависимости
- local.env # Конфиг
