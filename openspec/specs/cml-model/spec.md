# cml-model

## Purpose

Независимая библиотека go-cml предоставляет типизированные Go-модели и кодек
исходного CML 6.12.0. CLI находится в отдельном репозитории cml-lint.

## Requirements

### Requirement: Типизированные сущности
go-cml SHALL предоставлять Go structs для всех объектных правил pinned грамматики,
string enums для перечислений и типизированные Ref для cross references.
ContextMappingModel, BoundedContext, Aggregate, Entity, ValueObject и другие
объекты MUST быть доступны без обращения к map полям универсального AST.

#### Scenario: Прямое чтение модели
- **WHEN** Unmarshal декодирует service.cml в ContextMappingModel
- **THEN** пользователь SHALL прочитать BoundedContexts, Aggregates и DomainObjects
  как типизированные Go-поля; Unicode описания MUST сохраняться.

### Requirement: Сериализация CML
Marshal SHALL сериализовать Go-модель в исходный CML. Parse результата SHALL
сохранять модельные данные, включая ссылки, роли, строки, enums и операции.
Неизвестные или противоречивые варианты union MUST давать ошибку.

#### Scenario: Создание модели кодом
- **WHEN** пользователь создаёт BoundedContext с Aggregate и Entity на Go
- **THEN** Marshal SHALL вернуть CML, который принимает официальный parser.

#### Scenario: Round trip
- **WHEN** модель проходит Unmarshal, Marshal и повторный Unmarshal
- **THEN** сравнение модельных данных SHALL показать эквивалентность.

### Requirement: Независимость библиотеки
Runtime MUST использовать только Go standard library. Библиотека MUST не
содержать CLI, запуск процессов, network IO или init side effects.
Parse/Unmarshal SHALL работать в памяти; только явно вызванный ValidateFile
SHALL читать импортируемые файлы. Root package MUST иметь external tests.

#### Scenario: Импорт клиентом
- **WHEN** cml-lint импортирует github.com/grespyrad/go-cml v0.3.0
- **THEN** анализатор SHALL работать без копии grammar/parser в своём репозитории.

### Requirement: Каноническое имя модуля и пакета
Библиотека SHALL публиковаться как github.com/grespyrad/go-cml v0.3.0.
Корневой Go-пакет MUST называться cml; внешние клиенты SHALL вызывать
cml.Validate, cml.Unmarshal и остальные API без import alias.
Генераторы MUST сохранять имя cml при повторном запуске.

#### Scenario: Импорт без алиаса
- **WHEN** внешний клиент импортирует "github.com/grespyrad/go-cml"
- **THEN** существующий публичный API SHALL быть доступен через cml.
