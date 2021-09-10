# Тестовое задание на позицию бэкенд стажёра в Avito
[![Build And Deploy Status](https://github.com/Tullerpeton/avito-test-case/actions/workflows/docker-build.yaml/badge.svg)](https://github.com/Tullerpeton/avito-test-case/actions/workflows/docker-build.yaml)
[![Linter](https://github.com/Tullerpeton/avito-test-case/actions/workflows/linter.yaml/badge.svg)](https://github.com/Tullerpeton/avito-test-case/actions/workflows/linter.yaml)
[![codecov](https://codecov.io/gh/Tullerpeton/avito-test-case/branch/main/graph/badge.svg?token=33oVnqFbzY)](https://codecov.io/gh/Tullerpeton/avito-test-case)

## 1. Документация API
Документация API представлена в [swagger](https://app.swaggerhub.com/apis-docs/Tullerpeton/Avito-test-case/0.1.1)

# 2. Запуск приложения
Для запуска приложения:
```shell
docker-compose up
```

# 3. Структура БД
В качестве базы данных используется Postgresql
Таблицы БД, которые используются в проекте:
```SQL
CREATE TABLE balances (
    user_id SERIAL NOT NULL PRIMARY KEY,
    balance DECIMAL(12,2) NOT NULL,

    CONSTRAINT balance_value CHECK (balance >= 0)
);

CREATE TABLE transaction_types (
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL
);

CREATE TABLE transactions (
    id SERIAL NOT NULL PRIMARY KEY,
    type_operation INTEGER NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    user_id INTEGER NOT NULL,
    context TEXT NOT NULL,
    value DECIMAL(12,2) NOT NULL,

    FOREIGN KEY (user_id) REFERENCES balances (user_id),
    FOREIGN KEY (type_operation) REFERENCES transaction_types (id),
    CONSTRAINT operation_value CHECK (transactions.value >= 0)
);
```

# 4. Примеры тестовых запросов
### 4.1 Пополнение баланса
##### 4.1.1 Пользовательский баланс отсутсвует, происходит первое зачисление
Запрос: `http://localhost:8080/api/v1/user/balance/improve`

Тело запроса:
```json
{
    "id": 1,
    "value": 10.13
}
```

Ответ [code = 200]:
```json
{
    "balance": 10.13,
    "currency": "RUB"
}
```

##### 4.1.2 Пользовательский баланс уже существует, происходит пополнение
Запрос: `http://localhost:8080/api/v1/user/balance/improve`

Тело запроса:
```json
{
    "id": 1,
    "value": 7.07
}
```

Ответ [code = 200]:
```json
{
    "balance": 17.2,
    "currency": "RUB"
}
```

### 4.2 Снятие баланса
##### 4.2.1 Пользовательский баланс отсутсвует, нечего снимать
Запрос: `http://localhost:8080/api/v1/user/balance/withdraw`

Тело запроса:
```json
{
    "id": 3,
    "value": 10.13
}
```

Ответ [code = 403]:
```json
{
  "message": "User balance not found"
}
```

##### 4.2.2 Пользовательский баланс уже существует, происходит уменьшение
Запрос: `http://localhost:8080/api/v1/user/balance/withdraw`

Тело запроса:
```json
{
    "id": 1,
    "value": 2.10
}
```

Ответ [code = 200]:
```json
{
  "balance": 15.1,
  "currency": "RUB"
}
```

### 4.3 Получение пользовательского баланса
##### 4.3.1 Пользовательский баланс отсутсвует
Запрос: `http://localhost:8080/api/v1/user/balance/2`

Тело запроса:
```json
{
  
}
```

Ответ [code = 403]:
```json
{
  "message": "User balance not found"
}
```

##### 4.3.2 Пользовательский баланс уже существует, баланс успешно показан
Запрос: `http://localhost:8080/api/v1/user/balance/2`

Тело запроса:
```json
{

}
```

Ответ [code = 200]:
```json
{
  "balance": 15.1,
  "currency": "RUB"
}
```

##### 4.3.3 Пользовательский баланс уже существует, баланс успешно показан в валюте
Запрос: `http://localhost:8080/api/v1/user/balance/2?currency=USD`

Тело запроса:
```json
{

}
```

Ответ [code = 200]:
```json
{
  "balance": 0.2,
  "currency": "USD"
}
```

### 4.4 Перевод с одного баланса на другой
##### 4.4.1 Перевод успешно выполнен
Запрос: `http://localhost:8080/api/v1/user/balance/transfer`

Тело запроса:
```json
{
  "sender_id": 4,
  "receiver_id": 3,
  "value": 2.1
}
```

Ответ [code = 200]:
```json
{
  "receiver": {
    "balance": 5.1,
    "currency": "RUB"
  },
  "sender": {
    "balance": 0.9,
    "currency": "RUB"
  }
}
```

##### 4.4.2 У получателя отсутсвует первоначальный баланс
Запрос: `http://localhost:8080/api/v1/user/balance/transfer`

Тело запроса:
```json
{
  "sender_id": 4,
  "receiver_id": 300,
  "value": 0.1
}
```

Ответ [code = 200]:
```json
{
  "receiver": {
    "balance": 0.1,
    "currency": "RUB"
  },
  "sender": {
    "balance": 0.8,
    "currency": "RUB"
  }
}
```

##### 4.4.2 У отправителя недостаточно средств
Запрос: `http://localhost:8080/api/v1/user/balance/transfer`

Тело запроса:
```json
{
  "sender_id": 4,
  "receiver_id": 300,
  "value": 1000
}
```

Ответ [code = 409]:
```json
{
  "message": "User can not make operation"
}
```

### 4.5 Получение всех транзакций пользователя
##### 4.5.1 Только пагинация и сортировка по умолчанию
Запрос: `http://localhost:8080/api/v1/user/transaction/4?page=1&records=10`

Тело запроса:
```json
{

}
```

Ответ [code = 200]:
```json
[
  {
    "type_operation": "improve",
    "context": "user [id=4] single operation",
    "created_at": "2021-09-10T23:01:43.802558Z",
    "value": 5.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=3]",
    "created_at": "2021-09-10T23:09:22.058143Z",
    "value": 2.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=3]",
    "created_at": "2021-09-10T23:09:28.576632Z",
    "value": 2.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=300]",
    "created_at": "2021-09-10T23:11:54.435142Z",
    "value": 0.1
  }
]
```

##### 4.5.2 Только пагинация и сортировка по размеру транзакции
Запрос: `http://localhost:8080/api/v1/user/transaction/4?page=1&records=10&sort=sum`

Тело запроса:
```json
{

}
```

Ответ [code = 200]:
```json
[
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=300]",
    "created_at": "2021-09-10T23:11:54.435142Z",
    "value": 0.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=3]",
    "created_at": "2021-09-10T23:09:22.058143Z",
    "value": 2.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=3]",
    "created_at": "2021-09-10T23:09:28.576632Z",
    "value": 2.1
  },
  {
    "type_operation": "improve",
    "context": "user [id=4] single operation",
    "created_at": "2021-09-10T23:01:43.802558Z",
    "value": 5.1
  }
]
```

##### 4.5.3 Только пагинация и сортировка по размеру транзакции (по убыванию)
Запрос: `http://localhost:8080/api/v1/user/transaction/4?page=1&records=3&sort=sum&direction=desc`

Тело запроса:
```json
{

}
```

Ответ [code = 200]:
```json
[
  {
    "type_operation": "improve",
    "context": "user [id=4] single operation",
    "created_at": "2021-09-10T23:01:43.802558Z",
    "value": 5.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=3]",
    "created_at": "2021-09-10T23:09:22.058143Z",
    "value": 2.1
  },
  {
    "type_operation": "transfer",
    "context": "user [id=4] sent money to user [id=3]",
    "created_at": "2021-09-10T23:09:28.576632Z",
    "value": 2.1
  }
]
```

# 5. Пример конфига

Путь к конфигу: `configs/app/balance_service.yaml`

```yaml
postgresql-data:
  user: "postgres"
  password: "postgres"
  db-name: "balance_service_db"
  port: "5432"
  host: "localhost"
  sslmode: "disable"

server-data:
  host: ""
  port: "8080"

currency-converter:
  access-secret: {{ your_key }}
  base-currency: "RUB"
  format: "1"
  http-url: "http://api.exchangeratesapi.io/v1/latest"
```