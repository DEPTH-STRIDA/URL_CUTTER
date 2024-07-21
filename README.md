# URL Shortener Microservice / Микросервис для сокращения URL

[![Русская версия](#русская-версия)](#русская-версия) | [![English Version](#english-version)](#english-version)

---

# Русская версия

## Внимание!!!
### Это тестовый код. Он имеет множество ошибок и опасных моментов.

## Описание микросервиса для сокращения URL

Этот микросервис создан в соответствии с **техническим заданием (ТЗ)**, чтобы обеспечить функционал сокращения URL-адресов. Основная цель микросервиса - создание коротких URL, которые могут быть использованы вместо длинных URL-адресов.

### Основные требования

- Длина сокращенного URL-адреса должна быть *как можно короче*.
- Сокращенный URL может содержать цифры (`0-9`) и буквы (`a-z`, `A-Z`).

### Эндпоинты

#### 1. Создание сокращенного URL

- **Метод**: POST
- **URL**: [http://localhost:8080/](http://localhost:8080/)
- **Request (body)**: 
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

- **Response**: 
http://localhost:8080/qtj5opu

#### 2. Получение полного URL по сокращенному URL

- **Метод**: GET
- **URL**: [http://localhost:8080/qtj5opu](http://localhost:8080/qtj5opu)
- **Request (url query)**: 
http://localhost:8080/qtj5opu

- **Response (body)**: 
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

### Хранение данных

Микросервис имеет возможность хранить информацию в *памяти* или в базе данных *PostgreSQL* в зависимости от флага запуска.

### Язык программирования

Этот микросервис написан на языке программирования *Go* (Golang), обеспечивая *быстродействие* и *эффективное использование ресурсов*.

### Зависимости

Для работы микросервиса требуется наличие *Go* версии X.X.X и установленных зависимостей, указанных в файле `go.mod`.

### Использование

Чтобы запустить микросервис, выполните следующие шаги:

1. Клонируйте репозиторий на локальную машину.
2. Убедитесь, что все зависимости установлены с помощью `go mod tidy`.
3. Используйте `go build` для сборки проекта.
4. Запустите исполняемый файл, указав необходимые параметры (например, флаг хранения данных).

### Автор

Если у вас есть какие-либо вопросы или предложения по улучшению функционала, не стесняйтесь создавать *issues* или отправлять *pull requests* в репозиторий проекта. *Ваш вклад приветствуется!*

**Автор:** Тихомиров Максим Русланович  
**Контактная информация:** [awesome.gail@yandex.ru](mailto:awesome.gail@yandex.ru), [Telegram](https://t.me/Tichomirov2003)

---

# English Version

## Attention!!!
### This is test code. It has many errors and dangerous issues.

## Description of the URL Shortener Microservice

This microservice is created according to the **technical specification (TS)** to provide URL shortening functionality. The main goal of the microservice is to create short URLs that can be used instead of long URLs.

### Main Requirements

- The length of the shortened URL should be *as short as possible*.
- The shortened URL may contain digits (`0-9`) and letters (`a-z`, `A-Z`).

### Endpoints

#### 1. Create a shortened URL

- **Method**: POST
- **URL**: [http://localhost:8080/](http://localhost:8080/)
- **Request (body)**: 
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

- **Response**: 
http://localhost:8080/qtj5opu

#### 2. Get the full URL by the shortened URL

- **Method**: GET
- **URL**: [http://localhost:8080/qtj5opu](http://localhost:8080/qtj5opu)
- **Request (url query)**: 
http://localhost:8080/qtj5opu

- **Response (body)**: 
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

### Data Storage

The microservice can store information either *in memory* or in a *PostgreSQL database* depending on the startup flag.

### Programming Language

This microservice is written in the *Go* programming language (Golang), ensuring *high performance* and *efficient resource usage*.

### Dependencies

The microservice requires *Go* version X.X.X and the dependencies listed in the `go.mod` file to run.

### Usage

To run the microservice, follow these steps:

1. Clone the repository to your local machine.
2. Ensure all dependencies are installed using `go mod tidy`.
3. Use `go build` to build the project.
4. Run the executable, specifying the necessary parameters (e.g., data storage flag).

### Author

If you have any questions or suggestions for improving functionality, feel free to create *issues* or submit *pull requests* to the project repository. *Your contribution is welcome!*

**Author:** Tikhomirov Maxim Ruslanovich  
**Contact Information:** [awesome.gail@yandex.ru](mailto:awesome.gail@yandex.ru), [Telegram](https://t.me/Tichomirov2003)

</body>
</html>
