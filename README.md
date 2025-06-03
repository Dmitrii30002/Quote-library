# Music-library
Music library REST full API

## Описание
REST full API для управления библиотекой цитат с возможностью:
- Добавления/удаления/редактирования цитат
- Получения всех цитат
- Получения случайной цитаты
- Получение цитаты по автору

## Стек
* Язык: Go (стандартная библиотека)
* База данных: PostgreSQL
* Контейнеризация: docker-compose

# Запуск
1. Клонирование репозитория:
``` bash
git clone https://github.com/Dmitrii30002/Quote-library.git
```

2. Настройка конфигурации:<br>
В файле .env находятся данные кофнигурации базы данных, при необходимости поменяйте их. Также в файле config/config.json - находятся конфигурации сервера.

3. Запуск БД: \
   Запуск БД на устройстве, лиюо заапуск контейнера с PostgreSQL:
   ``` bash
   docker run --name postgres -e POSTGRES_DB=postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres:15
   ```

5. Запуск проекта:
``` bash
go build cmd/main.go
./main.exe
```

# Эндпоинты:

|Метод	    |Путь	              |Описание                                 |
|:----------|:------------------|:-------------------------------------------|
|GET	      |/quotes	            |Список всех цитат                          |
|GET	      |/quotes/random	   |Получить случайную цитату                  |
|GET	      |/quotes?author   	|Получить цитату автора по его имени        |
|DELETE	   |/quotes/{id}	      |Удалить цитату                             |
|Post	      |/quotes	            |Добавить цитату                            |

# Структура проекта
* **cmd** -  точка входа в приложение — исполняемая программа
* **config** - загрузка конфигураций в приложение
* **internal** - Основная логика
* **internal/handlers** - реализация хэндлеров
* **internal/models** - реализация моделей
* **internal/repository** - бизнес-логика
* **internal/migrations** - миграции
