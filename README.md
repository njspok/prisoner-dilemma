# Prisoner dilemma

[Дилемма заключенного](https://en.wikipedia.org/wiki/Prisoner%27s_dilemma) 
    - фундаментальная проблема в теории игр.
Программа моделирует игру Game с разными стратегиями, которые реализуются игроками Player.
Между всеми игроками проводятся турнир Tournament после которого публикуются результаты игр.

# Описание

Игроки делают ход, не зная как сходил другой игрок. 
Игрок может сходить КООПЕРИРУЮСЬ или ОТКАЗЫВАЮСЬ.
Всего возможно четыре исхода после того как сходили игроки:
- оба сыграли КООПЕРИРУЮСЬ, то получат, например 3 очка
- оба сыграли ОТКАЗЫВАЮСЬ (в кооперации), то получат 1 очко
- первый КООПЕРУЮСЬ, второй ОТКАЗЫВАЮСЬ - первый получит 5 очков, второй 0
- первый ОТКАЗЫВАЮСЬ, второй КООПЕРИРУЮСЬ - первый получит 0 очков, второй 5

Как видно, кооперироваться выгоддно обоим, но если один отказался кооперироваться, то его выигрыш
больше, чем если бы они одновременно скооперировались.

# Запуск и результаты

```bash
# go run main.go
```
с бинарником
```bash
# go build
# ./prisoner-dilemma
```
после чего можно увидеть в консоли табличу с результатами игр, именами игроков, сколько
они заработали очков и сколько набрали очков в сумме
```
   Player 1 |    Score |    Player 2 |    Score  | Total
      Jesus |      300 |       Jesus |      300  | 600
      Jesus |        0 |       Judah |      500  | 500
      Jesus |      153 |       Crazy |      398  | 551
      Jesus |      300 |      Talion |      300  | 600
      Judah |      500 |       Jesus |        0  | 500
      Judah |      100 |       Judah |      100  | 200
      Judah |      324 |       Crazy |       44  | 368
      Judah |      104 |      Talion |       99  | 203
      Crazy |      388 |       Jesus |      168  | 556
      Crazy |       55 |       Judah |      280  | 335
      Crazy |      200 |       Crazy |      255  | 455
      Crazy |      231 |      Talion |      231  | 462
     Talion |      300 |       Jesus |      300  | 600
     Talion |       99 |       Judah |      104  | 203
     Talion |      215 |       Crazy |      215  | 430
     Talion |      300 |      Talion |      300  | 600
```

# Стратегии

Есть несколько базовых стратегий:

- jesus - всегда играет КООПЕРИРУЮСЬ
- judah - всегда играет ОТКАЗЫВАЮСЬ
- crazy - случайно выбирает КООПЕРИРУЮСЬ или ОТКАЗЫВАСЬ с вероятностью 50%
- talion - выбирает стратегию ОКО ЗА ОКО, то есть просто повторяет предыдущий ход противника

