<h1 align="center">Тестовое задание на позицию стажера-бекендера</h1>



### Установка
`git clone https://github.com/Woodman5/avito-intern`

### Запуск
`docker-compose -p avito-intern -f deployments/docker-compose.yml up --build`


### Запросы
Получение сведений о балансе пользователя

`127.0.0.1:8080/amount/{user UUID}`

Зачисление на счет / Снятие со счета
`127.0.0.1:8080/amount/56a36a84-fb53-4bfa-a2dd-4307c4b3c988`