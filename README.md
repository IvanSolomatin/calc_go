set "PORT=8087" & "main.exe"
curl -X POST http://localhost:8080/api/v1/calculate -d "{\"expression\": \"1\"}"
curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1\"}"
400 [HTTP] 2024/12/20 18:52:45 aplication.go:80: Bad Request: invalid character '\r' in string literal
go get github.com/IvanSolomatin/calc_go/internal/application



# Веб-сервис для вычисления арифметических выражений

## Описание
Этот проект реализует веб-сервис, который вычисляет арифметические выражения, переданные пользователем через HTTP-запрос.

## Структура проекта

- `cmd/` — точка входа приложения.
- `internal/application` — веб-сервис.
- `pkg/calculation ` — вспомогательные пакеты и утилиты.


## Запуск

1. Установите [Go](https://go.dev/dl/).
2. Установите [Git](https://git-scm.com/downloads).
3. Склонируйте проект с GitHub используя командную строку:
    git clone https://github.com/IvanSolomatin/calc_go
4. Перейдите в папку проекта и запустите сервер(calc_go):

    go run ./cmd/main.go
    
5. Сервис будет доступен по адресу: [http://localhost:8080/api/v1/calculate](http://localhost:8080/api/v1/calculate).

### Альтернативный запуск
Вы можете использовать скрипты для сборки и запуска:
- **Для Linux/MacOS:**
    ```bash
    ./build/build.sh
    ```
- **Для Windows:**
    ```powershell
    .\build\build.bat
    ```

## Эндпоинты

### `POST /api/v1/calculate`

#### Описание
Эндпоинт принимает JSON с математическим выражением.

#### Пример запроса с использованием PowerShell

```powershell
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "2+2*2"}'
Пример успешного ответа
json
Копировать код
{
  "result": "6.000000"
}
Пример ошибки 500
Если выражение содержит некорректный символ $, сервер вернёт ошибку 500:

Пример запроса
powershell
Копировать код
Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
-Method POST `
-Headers @{"Content-Type"="application/json"} `
-Body '{"expression": "1+$2"}'
Пример ответа
json
Копировать код
{
  "error": "Некорректное выражение"
}
Тестирование
Для запуска тестов выполните:

bash
Копировать код
go test ./...
Примечания
Для работы API требуется установленный Go (версии 1.18 и выше).
Все зависимости проекта управляются через go mod. Убедитесь, что в корне проекта находятся go.mod и go.sum.