# API Блога на Go

Простой API для управления записями блога с хранением данных в памяти.

## Особенности

- CRUD операции для записей блога
- Валидация данных (проверка пустых заголовков и контента)
- Хранение данных в памяти (переменной)
- JSON формат для запросов и ответов
- Логирование операций
- Использование только стандартной библиотеки Go
- Поддержка CORS для фронтенд-интеграции
- RESTful API дизайн

## Структура проекта

```
golang-blog-web/
├── go.mod
├── main.go
├── README.md
├── frontend/               # Фронтенд приложение
│   ├── index.html          # Основная HTML страница
│   ├── css/                # Стили
│   │   └── styles.css
│   └── js/                 # JavaScript
│       └── main.js
└── internal/
    ├── models/
    │   └── post.go          # Модель записи блога
    ├── services/
    │   └── post_service.go  # Бизнес-логика
    ├── handlers/
    │   └── post_handlers.go # HTTP хендлеры
    └── middleware/
        └── cors.go          # CORS middleware
```

## Запуск проекта

1. Убедитесь, что у вас установлен Go
2. Клонируйте репозиторий или скопируйте файлы проекта
3. Запустите сервер:

```bash
go run main.go
```

Сервер будет доступен по адресу: `http://localhost:8080`

#### Конфигурация через переменные окружения

Вы можете настроить сервер и CORS через переменные окружения:

```bash
# Запуск на хосте 0.0.0.0
SERVER_HOST=0.0.0.0 go run main.go

# Запуск на кастомном порту
SERVER_PORT=8000 go run main.go

# Настройка CORS origin
CORS_ALLOWED_ORIGIN="http://localhost:5500" go run main.go

# Комбинированный запуск
SERVER_HOST=0.0.0.0 SERVER_PORT=8080 CORS_ALLOWED_ORIGIN="http://localhost:3000" go run main.go
```

По умолчанию используется:
- Порт: 8080
- Хост: localhost
- CORS origin: * (разрешены все адреса)

## API Эндпоинты

### Создание записи
```
POST /api/posts/
Content-Type: application/json

{
  "title": "Заголовок записи",
  "content": "Содержимое записи"
}
```

### Получение записей
```
GET /api/posts/          # Все записи
GET /api/posts/{id}      # Конкретная запись по ID
```

### Обновление записи
```
PUT /api/posts/{id}
Content-Type: application/json

{
  "title": "Новый заголовок",
  "content": "Новое содержимое"
}
```

### Удаление записи
```
DELETE /api/posts/{id}
```

## Примеры использования

### Создание записи
```bash
curl -X POST http://localhost:8080/api/posts/ \
  -H "Content-Type: application/json" \
  -d '{"title":"Первая запись","content":"Привет, мир!"}'
```

### Получение всех записей
```bash
curl http://localhost:8080/api/posts/
```

### Получение конкретной записи
```bash
curl http://localhost:8080/api/posts/1
```

### Обновление записи
```bash
curl -X PUT http://localhost:8080/api/posts/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Обновленный заголовок","content":"Обновленное содержимое"}'
```

### Удаление записи
```bash
curl -X DELETE http://localhost:8080/api/posts/1
```

## Валидация

- Заголовок не может быть пустым
- Заголовок не может быть длиннее 500 символов
- Контент не может быть пустым
- Контент не может быть длиннее 10000 символов

При попытке создать или обновить запись с пустыми полями, сервер вернет ошибку 400 Bad Request.

## CORS Поддержка

API поддерживает CORS (Cross-Origin Resource Sharing) для интеграции с фронтенд-приложениями.

### Настройка CORS

- **Разрешенный источник**: Настраивается через переменную окружения `CORS_ALLOWED_ORIGIN`
- **Значение по умолчанию**: * (разрешены все адреса)
- **Разрешенные методы**: GET, POST, PUT, DELETE, OPTIONS
- **Разрешенные заголовки**: Content-Type, Authorization

### Конфигурация через переменные окружения

Вы можете настроить CORS origin через переменную окружения:

```bash
# Для фронтенда на другом порту
CORS_ALLOWED_ORIGIN="http://localhost:5500" go run main.go

# Для production
CORS_ALLOWED_ORIGIN="https://yourdomain.com" go run main.go

# Для нескольких доменов (разделяйте запятыми)
CORS_ALLOWED_ORIGIN="https://yourdomain.com,https://api.yourdomain.com" go run main.go
```

### Как изменить разрешенный источник

1. **Рекомендуемый способ**: Используйте переменную окружения
   ```bash
   export CORS_ALLOWED_ORIGIN="https://yourdomain.com"
   go run main.go
   ```

2. **Или измените значение по умолчанию** в файле `internal/config/config.go`

### Пример для разработки

По умолчанию разрешены запросы с любого источника (только для разработки)

> ⚠️ **Внимание**: Использование `*` в production не рекомендуется по соображениям безопасности. Всегда указывайте конкретные домены для production окружения.

## Логирование

Все операции логируются в стандартный вывод. Вы увидите сообщения о:
- Создании, обновлении и удалении записей
- Ошибках валидации
- Запросах к API

## Разработка

Проект использует стандартную библиотеку Go без внешних зависимостей. Для добавления новых функций:

1. Обновите модель в `internal/models/post.go`
2. Добавьте методы в сервис `internal/services/post_service.go`
3. Создайте новые хендлеры в `internal/handlers/post_handlers.go`
4. Настройте маршруты в `main.go`

## Лицензия

[LICENSE](LICENSE)
