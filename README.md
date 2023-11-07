# ETU Shop Service

Этот репозиторий содержит все необходимые материалы для поднятия фронтенд- и бэкенд-части нашего решения.

## Стек

### Дизайн

[Ссылка на проект в figma](https://www.figma.com/file/q107fwWYtziBQ45aUTQE0x/LETI-STORE)

Создание дизайна проходило в 3 этапа:
1. Аналитический
2. Прототипирование
3. Дизайн

#### Аналитический этап
Первостепенно для создания сайта необходимо понимать целевую аудиторию проекта. В нашем случае это студенты ЛЭТИ, реже - сотрудники. После определения ЦА research референсев, в процессе которого были подобраны оптимальные стилистические решения, которые в дальнейшем будут отображены в сайте. Был составлен мудборд проекта. После определения ЦА, стиля, составления мудборда, мы приступили к проработке структуры сайта в программе Xmind

#### Создание прототипа
На следующем этапе на основании структуры и мудборда был создан прототип основных страниц сайта. Прототип - отображение элементов, их размещения и размеров, без использования контента и цветов.
P.S. в дальнейшем будут созданы прототипы для второстепенных страниц, на основании которых уже будет сделан их дизайн

#### Дизайн
Заключительным этапом является создание дизайна, которое подразумевает за собой "разукрашивание" прототипа. В процессе дизайна возможны незначительные изменения позиционирования и размеров, а также добавление дополнительных декоративных элементов

### Фронтенд

| Технология | Задача             |
|------------|--------------------|
| **Vite**   | Сборка приложения  |
| **Vue**    | Фреймворк          |
| **Pinia**  | Менеджер состояний |
| **Vue Router**  | Маршрутизация |

Vite - один из инструментов для настройки сред разработки. Для обеспечения плавной и быстрой разработки с минимальными предварительными настройками, Vite использует самые продвинутые функции современного JavaScript, такие как модули ES. Такие принципы позволяют делать как компиляцию фреймворка, так и использовать уже готовые пресеты. 

Vue.js нужен для быстрой перерисовки того или иного участка интерфейса. Вместо работы напрямую с DOM, мы работаем с данными (Data), а все манипуляции с разметкой происходят уже автоматически благодаря реактивности и директивам. То есть, темплейты во Vue.js - это такой себе способ описать поведение вашей разметки с привязкой к данным - как меняется разметка при изменении данных, какие методы вызываются при том или ином событии (это тоже определяется в темплейтах).

Pinia - легковесная библиотека управления состояниями для Vue.js. Она использует новую систему реактивности во Vue 3 для создания интуитивно понятной и полностью типизированной библиотеки управления состояниями.

Vue Router — официальный маршрутизатор для Vue.js. Он глубоко интегрируется с ядром Vue.js, что упрощает создание одностраничных приложений (SPA) с помощью Vue.js.

Composition API — это набор API, который позволяет нам создавать компоненты Vue, используя импортированные функции вместо объявления параметров. Это общий термин, охватывающий следующие API:
1. API реактивности, например. ref() и reactive(), которые позволяют нам напрямую создавать реактивное состояние, вычисляемое состояние и наблюдателей.
2. Перехватчики жизненного цикла, например. onMounted() и onUnmounted(), которые позволяют нам программно подключиться к жизненному циклу компонента.
3. Внедрение зависимостей, то есть provide() и inject(), которые позволяют нам использовать систему внедрения зависимостей Vue при использовании API-интерфейсов реактивности.

### Настройка проекта
```sh
npm install
```

### Dev-сборка:
```sh
npm run dev
```

### Prod-сборка:
```sh
npm run build
```

### Бэкенд

| Технология     | Задача                                     |
|----------------|--------------------------------------------|
| **Golang**     | Язык разработки, сборка приложения         |
| **PostgreSQL** | База данных                                |

Golang является современным языком разработки бэкенд-приложений. Его компилируемость и строгая типизация позволяют добиться высокой производительности и небольшого размера скомпилированного приложения. На нем можно реализовать как микросервис, который отвечал бы за конкретную область сайта, так и монолитный проект, содержащий в себе все необходимые модули.

В угоду скорости разработки и простоты развертывания был выбран подход с монолитом. Итоговый докер-образ содержит как логическую API-часть, так и статичный HTML+CSS+JS контент.

PostgreSQL была выбрана как одна из реляционных баз данных, обеспечивая более строгий контроль над данными и прекрасно сочетаясь с языком с такой же строгой типизацией.

### Поддержка

| Технология         | Задача                                                       |
|--------------------|--------------------------------------------------------------|
| **Docker**         | Контейнеризация всего проекта                                |
| **GitHub Actions** | Облачная сборка проекта; выполнение миграций на production-сервере |
| **Goose**          | Инструмент миграции базы данных                              |

Для повышения портативности проекта и возможности его запуска на разных платформах был использован Docker. Сборка проходит в три этапа:
1. Сборка бэкенда (стадия `builder`)
2. Сборка фронтенда (стадия `builder_frontend`)
3. Компоновка

Такой подход (multi-stage build) позволил добиться небольшого размера докер-образа, поскольку в него не входит исходный код Go-библиотек (оставшихся на стадии `builder`) и код NodeJS библиотек (оставшихся на стадии `builder_frontend`).

GitHub Actions позволяет автоматизировать сборку и загрузку докер-образа в общедоступный репозиторий. Также в скрипт добавлен отдельный job, выполняющий миграции на production-сервере.

Goose позволяет безболезненно вносить изменения в базу данных (и откатывать их же). В его "распоряжении" находится папка `migrations`. В ней хранятся SQL-скрипты, последовательно выполняя которые, структура базы данных приходит в актуальное состояние.

## Docker-образ

Для запуска Docker-образ может использовать следующие переменные окружения

| Название            | Описание                                   | Значение по умолчанию |
|---------------------|--------------------------------------------|-----------------------|
| `STARTUP`           | **(обязательно)** Команда запуска сервера* |                       |
| `SERVER_PORT`       | Порт, на котором запустится сервер         | `9000`                |
| `POSTGRES_HOST`     | Адрес базы данных**                        | `localhost`           |
| `POSTGRES_PORT`     | Порт базы данных                           | `5432`                |
| `POSTGRES_USER`     | Имя пользователя базы данных               | `shop`                |
| `POSTGRES_PASSWORD` | Пароль пользователя базы данных            | `shop`                |
| `POSTGRES_DB`       | Название базы данных                       | `shop`                |

(*) Эта переменная присутствует только в Production-сборке (тег `main`) и добавлен для того, чтобы в менеджерах контейнеров по типу Pterodactyl или Portainer была возможность изменить строку запуска (например, вывести содержимое какого-то файла вместо запуска сервера или параллельный запуск нескольких процессов).\
При использовании Production-сборки, чтобы сервер запустился, нужно явно указать значение `/app` (при необходимости экранировав слэш в начале).\
Если собрать локальную версию проекта или использовать Dev-версию (тег `main-dev`), то эта переменная не используется.

(**) Подразумевает наличие работающей базы данных. Если использовать Production-сборку (тег `main`), то база данных также должна быть в актуальном состоянии. В локальной и Dev-сборке (тег `main-dev`) образ также содержит инструмент автоматической миграции.

### Сборка локальной версии
```shell
docker build -t ghcr.io/merchleti/service:local-dev -f dev.Dockerfile .
```

### Запуск из командной строки

#### Production-сборка:
```shell
docker run --env STARTUP='\/app' --env SERVER_PORT=8081 --env POSTGRES_HOST=host.docker.internal -p 8081:8081 -it ghcr.io/merchleti/service:main
```

### Dev-сборка:
```shell
docker run --env SERVER_PORT=8081 --env POSTGRES_HOST=host.docker.internal -p 8081:8081 -it ghcr.io/merchleti/service:main-dev
```

#### Локальная версия:
```shell
docker run --env SERVER_PORT=8081 --env POSTGRES_HOST=host.docker.internal -p 8081:8081 -it ghcr.io/merchleti/service:local-dev
```

### Пример docker-compose
```yaml
version: '3.3'

services:
  postgres:
    image: postgres:16-alpine
    environment:
      POSTGRES_DB: shop
      POSTGRES_USER: shop
      POSTGRES_PASSWORD: shop
    healthcheck:
      test: ["CMD-SHELL", "pg_isready"]
      interval: 10s
      timeout: 5s
      retries: 5
    ports:
      - "5432:5432"
  shop:
    image: ghcr.io/merchleti/service:main-dev
    depends_on:
      postgres:
        condition: service_healthy
    environment:
      SERVER_PORT: 8081
      POSTGRES_HOST: postgres
    ports:
      - "8081:8081"
```

## Тестовые данные

Пример данных для БД (в последней версии идентичны тем, что находятся на production-сервере) находится в [SQL-дампе](https://raw.githubusercontent.com/MerchLeti/service/main/test/data.sql).
