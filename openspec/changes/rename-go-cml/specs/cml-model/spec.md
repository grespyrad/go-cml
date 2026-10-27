## ADDED Requirements

### Requirement: Каноническое имя модуля и пакета
Библиотека SHALL публиковаться как github.com/grespyrad/go-cml v0.3.0.
Корневой Go-пакет MUST называться cml; внешние клиенты SHALL вызывать
cml.Validate, cml.Unmarshal и остальные API без import alias.
Генераторы MUST сохранять имя cml при повторном запуске.

#### Scenario: Импорт без алиаса
- **WHEN** внешний клиент импортирует "github.com/grespyrad/go-cml"
- **THEN** существующий публичный API SHALL быть доступен через cml.
