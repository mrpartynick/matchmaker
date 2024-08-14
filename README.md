  Здравствуйте! Большое спасибо за внимание к моему тестовому заданию. Хочу сказать, что само задание мне очень понравилось, поскольку основывалось не на машинальном применении стандартных 
  индустриальных технологий, а на том, чтобы подумать и прежде всего найти подходящий математический алгоритм. Надеюсь вам также понравится мое решение :) 

  # Содержание 
  1. [Описание дизайна](#Описание-дизайна)

     [Алтернативное решение](#альтернативное-решение)
3. [Example2](#example2)
4. [Third Example](#third-example)
5. [Fourth Example](#fourth-examplehttpwwwfourthexamplecom)

  # Описание дизайна 
  Мое решение основывается на распределении: каждый раз, когда в пуле появляется новый игрок, он либо отправляется в существующую подходящую группу, либо, если таковой 
  не найдется, создается новая группа, куда отправляется игрок. 
  
  Каким образом определяется подходящая для игрока группа? Начнем с того, что когда создается новая группа, skill и latency первого добавленного в неё игрока принимаются как skill и latency 
  группы. Далее, для того, чтобы новый игрок был распределен в эту группу, необходимо, чтобы его skill и latency отличались от skill и latency группы не более чем на 
  skillTolerance и latencyTolerance ( значения обоих этих параметров берутся из конфига ). Таким образом получается, что разница в уровне и задержке между игроками в группе будет минимальной.
  И не будет выходить за рамки установленных администратором пределов. 

  Я пытался высчитать значение параметров skillTolerance и latencyTolerance математически, однако понял, что это не имеет смысла и значение этих параметров должно определяться администратором
, исходя из контекста. Также при необходимости можно завести ИИ, которое, учитывая такие факторы, как кол-во игроков, их среднюю разницу, время ожидания и т.д будет устанавливать значение параметров. 

  ## Альтернативное решение 
  Другим решением, которое я придумал в первую очередь, было следующее: отсортировать имеющихся в пуле игроков и последовательно брать первых n ( n = groupSize ). 
  Оставался вопрос о том, как выполнить многокритериальную сортировку. Я решил прибегнуть к методу TOPSIS. 
  
  Таким образом получалось, что достигалось желаемое: разница в показателях была минимальной, однако легко было представить себе ситуацию, в которой в одной группе окажется игрок 
  с навыком 10 условных единиц и игрок с навыком в 100 условных единиц, так как второй игрок имеет следующий минимальный скилл после первого игрока. И если в таких 
  играх, как Battlefield, это не является критичным, так как уровень является лишь показателем того, сколько времени игрок провел в игре, то в таких играх, как Word of Tanks или War thunder, 
  уровень является отражением материального оснащения игрока. И я решил, что лучше бы я пдольше подождал, чем меня с бипланчиком закинет в матч против ребят на реактивных самолетах. 

  Поэтому я отказался от этой реализации в пользу более хорошего user experience :)

  # Архитектура проекта 
  ## Структура конфига 
  ```
  env: -- Уровень окружения. Влияет на логи. 
    env: debug
  server: -- Параметры для сервера. Хост и порт. 
    host: localhost
    port: 8080
  matchmaker: -- Параметры самого матчмейкера
    group_size: 5 -- Размер группы 
    skill_tolerance: 2 -- Толерантность к скиллу 
    latency_tolerance: 100 -- Толерантность к задержке 

  ```
  ## Структура проекта 
  В качестве архитектурного подхода я решил воспользоваться Clean Archetecture с одним маленьким уточнением: я не стал выделять слой Domain Services, поскольку для такого маленького 
  проекта это было бы overengineering. 

  Файловая структура проекта отображена здесь: 
  ```
  .
  ├── cmd
  │   └── matchmaker
  │       └── main.go -- entrypoint сервиса
  ├── config 
  │   └── config.go -- конфиг-модуль сервиса
  │ 
  ├── internal 
  │   ├── matchmaker
  │   │   ├── matchmaker.go -- матчмейкер
  │   │   └── group-methods.go -- методы для управления группами
  │   │
  │   ├── models -- модели данных
  │   └── server -- сервер
  │ 
  ├── pkg
  │   └── logging
  │       └── logging.go -- сервис логирования
  └── config.yaml -- конфиг-файл
  ```
  В качестве фремворка для построения HTTP-сервера был использован Gin 

  # Затраченные ресурсы 
  Время на разработку дизайна - 4 часа 
  Время на реализацию - 2,5 часа 
