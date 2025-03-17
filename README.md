# REST API для управления задачами, написанный на Go с использованием фреймворка Fiber и базы данных PostgreSQL.

## CRUD операции для задач:
  - Создание задачи (POST /task Content-Type: application/json)
```json
{
  "userId" : "id пользователя",
  "title": "Название",
  "description": "Описание"
}
```
- Создание пользователя (POST /user Content-Type: application/json)
```json
{
  "username" : "имя пользователя",
  "password": "пароль"
}
```
  - Получение списка задач по username (GET /task/:username)
  - Получение задачи по ID (GET /task/:id)
  - Обновление задачи (PUT /task/:id
Content-Type: application/json)
```json
{
  "userId" : "id пользователя",
  "title": "Название",
  "description": "Описание",
  "status": "new, in_progress ,done"
}
```
  - Удаление задачи (DELETE /task/:id)
  - Удаление пользователя (DELETE /user/:userId)

## Технологии
- Fiber - веб-фреймворк
- godotenv - загрузка .env файлов
- zap - логгер
- pgx - драйвер для postgreSQL

## Запуск
```bash
go run cmd/main.go
```
Сервис будет доступен по адресу: http://localhost:8080

## Создание таблиц
```sql
CREATE TABLE users (
id SERIAL PRIMARY KEY,
username TEXT UNIQUE NOT NULL,
password TEXT NOT NULL, -- для упрощения, в продакшн разработке пароли так не хранятся
created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE tasks (
id SERIAL PRIMARY KEY,
user_id INT REFERENCES users(id) ON DELETE CASCADE,
title TEXT NOT NULL,
description TEXT,
status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
created_at TIMESTAMP DEFAULT now()
);
```
## Структура 
- cmd/ Точка входа приложения
- config/ Обработка конфига
- dto/ Http ответы 
- logger/ Логгер
- repo/ Работа с хранилицем
- sqlRequests/ запросы к бд
- service/ Бизнес логика
- go.mod # Зависимости
- local.env # Конфиг
