# worker_pull
Тестовое задание на бэкенд-разработчика Go 
### Запуск
```
go run main.go
```
### Работа с утилитой
-  Запуск задачи (джобы)
```
job input
```
текущая функция делает реверс строки. Можно передать другую функцию.
Также сейчас стоит sleep на 10 секунд для имитации "работы".
- Добавление воркеров
```
add number
```
- Удаление воркеров
```
del number
```
Если удалить воркеров больше, чем есть на данный момент, то их удаление отложиться до появления
(то есть следующие созданные воркеры сразу удалятся 😀).
- Выход
```
ext
```
### Пример работы
```
job abc
worker_id: 1, job_id: 0, res: [cba]
job abcd
worker_id: 2, job_id: 1, res: [dcba]
job abcdsaad
worker_id: 3, job_id: 2, res: [daasdcba]
job aboba
job abcasd
worker_id: 1, job_id: 3, res: [aboba]
worker_id: 2, job_id: 4, res: [dsacba]
del 3
worker_id: 3 done
worker_id: 1 done
worker_id: 2 done
job aaabbbccc
job abcabcabc
job abcdefghi
job anton
job helloworld
add 5
worker_id: 5, job_id: 5, res: [cccbbbaaa]
worker_id: 6, job_id: 7, res: [ihgfedcba]
worker_id: 4, job_id: 6, res: [cbacbacba]
worker_id: 7, job_id: 8, res: [notna]
worker_id: 8, job_id: 9, res: [dlrowolleh]
ext
```
