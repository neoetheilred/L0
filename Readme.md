# L0

#### Автор: Роман Андриевский
---
### Использование:

- Поднять сервис:
    ```
    $ docker-compose up
    ```
- Послать в nats-streaming фейковые данные:
    ```
    $ cd servce
    # go run ./util/main.go <Количество записей order>
    $ go run ./util/main.go 100
    ```
    Посылает в nats-streaming 100 случайно сгенерированных моделей

- Web клиент доступен по адресу ```localhost:8081```
    
    В поле для ввода нужно вводить ```order_uid```

    Справа список доступных id (закешированных в сервисе), при нажатии на id клиент получает объект из сервиса
    
- База данных (postgres): 
    - по адресу ```localhost:8080``` доступен клиент adminer
        ```
        имя пользователя: root
        пароль: toor
        база данных: root
        ```
    - таблицы:
        ```
        orders - основная таблица, где хранятся данные в соответствии с моделью из условия
        deliveries, items, payments - таблицы для хранения внутренних структур модели
        ```

        подробнее см. ```postgres/1.sql```

---

### Отказоустойчивость, кеш и валидация

- Если в сервис приходит не json - то сервис ничего не делает, продолжает ожидать данные

- Если поступает валидный json, тогда сервис валидирует полученный объект (с помощью ```github.com/go-playground/validator/v10```), если объект не проходит валидацию, то сервис ничего не делает, продолжает ожидать данные

- При создании сервиса из БД в кеш загружаются 100 последних записей (самых новых), впоследствии, одновременно с записью в БД, в кеше сохраняются получаемые данные

- При падении сервиса он перезапускается, при запуске кеш загружается из БД