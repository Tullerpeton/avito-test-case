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
