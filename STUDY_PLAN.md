# Подготовка к зимней стажировке Avito Backend Go

## 0. Как использовать этот файл

Этот файл — постоянный учебный контекст для подготовки к зимнему набору на стажировку **Avito Backend Go**.

Основная цель:
- пройти отбор на зимнюю стажировку Avito Backend Go;
- быть готовым не только к алгоритмам, но и к Go Platform, SQL/PostgreSQL, HTTP/сетям, Git/Unix, Docker, тестированию и защите тестового задания;
- уметь объяснять свои решения вслух и защищать архитектурные решения.

Ориентир по нагрузке:
- **4–5 часов в день**;
- алгоритмы не прекращаются во время изучения backend/Go;
- после завершения первых 50 задач — обычно **3 новых + 2 повторных алгоритмических задачи в день**;
- 1 раз в неделю — mock interview / timed session.

Важно:
- даты зимнего набора 2026/27 на момент составления плана официально не подтверждены;
- собственный дедлайн готовности: **к середине–концу октября уже быть готовым проходить отбор**;
- если набор откроется раньше — подаваться сразу, не ждать «идеальной готовности».

---

# 1. Что ожидают на отборе Avito Backend Go

По материалам о собеседовании Avito:

## Скоринговое интервью

Обычно проверяются базовые hard skills:
- алгоритмы и структуры данных;
- Go;
- HTTP;
- SQL;
- Git;
- Unix/Linux.

Это короткий технический фильтр перед большим интервью.

## Основное техническое интервью

Ориентир по формату:
- примерно **2–3 часа**;
- live coding;
- две основные секции:
  1. **Programming**
  2. **Go Platform**
- System Design может быть дополнительной секцией.

## Programming

Ориентир:
- примерно **1 час**;
- обычно одна Easy + одна Medium;
- оцениваются:
  - способность придумать алгоритм;
  - написание и отладка кода;
  - corner cases;
  - структуры данных;
  - временная сложность;
  - сложность по памяти;
  - способность объяснять решение вслух.

Целевой уровень:
- Easy: уверенно за 10–15 минут;
- Medium: в среднем 20–35 минут;
- самостоятельно определять паттерн;
- объяснять Time / Space Complexity;
- самостоятельно придумывать edge cases.

## Go Platform

Особенно важны:
- pointers;
- slices;
- maps;
- interfaces;
- goroutines;
- channels;
- select;
- WaitGroup;
- Mutex / RWMutex;
- atomic;
- race condition;
- deadlock;
- context;
- timeouts / cancellation;
- HTTP;
- сетевые основы;
- graceful shutdown;
- connection reuse / keep-alive;
- базовое понимание Go runtime.

Нужно уметь не просто пересказать определение, а:
- прочитать кусок Go-кода;
- найти проблему;
- объяснить причину;
- исправить;
- объяснить компромисс.

---

# 2. Целевой уровень к подаче

## Алгоритмы — 10/10 по приоритету

Нужно уверенно знать:
- Big O;
- Array / Slice;
- HashMap / Set;
- Strings;
- Two Pointers;
- Sliding Window;
- Prefix Sum;
- Binary Search;
- Stack / Queue;
- Linked List;
- Trees;
- BFS / DFS;
- Heap;
- базовые intervals;
- базовые graph problems.

Второй приоритет:
- backtracking;
- basic dynamic programming;
- более сложные heap/tree/graph задачи.

Не главный приоритет до первого отбора:
- segment tree;
- Fenwick tree;
- сложный DP;
- сложные олимпиадные графовые алгоритмы.

## Go — 10/10

Нужно уверенно:
- syntax;
- arrays/slices;
- map;
- structs;
- methods;
- pointer/value receiver;
- interfaces;
- nil;
- errors;
- errors.Is / errors.As;
- defer;
- panic/recover;
- packages/modules.

Concurrency:
- goroutine;
- channel;
- buffered/unbuffered;
- close;
- select;
- WaitGroup;
- Mutex;
- RWMutex;
- atomic;
- race;
- deadlock;
- goroutine leak;
- go test -race;
- go vet.

Context:
- Context;
- WithCancel;
- WithTimeout;
- deadline;
- ctx.Done();
- зачем вызывать cancel.

Runtime — базово:
- goroutine vs OS thread;
- M:N;
- GOMAXPROCS;
- scheduler;
- GC;
- stack growth;
- netpoll — на концептуальном уровне.

## SQL/PostgreSQL — 9/10

Нужно:
- SELECT;
- WHERE;
- ORDER BY;
- LIMIT/OFFSET;
- INNER JOIN;
- LEFT JOIN;
- несколько JOIN;
- GROUP BY;
- HAVING;
- aggregate functions;
- subqueries;
- CTE;
- window functions — basics;
- ACID;
- transactions;
- COMMIT / ROLLBACK;
- isolation — основы;
- PK/FK;
- UNIQUE;
- NOT NULL;
- indexes;
- composite indexes;
- EXPLAIN ANALYZE;
- pagination;
- атомарные операции.

## HTTP / backend — 9/10

Нужно:
- HTTP request/response;
- methods;
- status codes;
- headers;
- JSON;
- REST;
- CRUD;
- validation;
- middleware;
- auth/JWT basics;
- timeout;
- retry basics;
- idempotency;
- graceful shutdown;
- HTTP client reuse;
- keep-alive;
- connection pool;
- DNS/TCP basics.

## Архитектура — 8/10

Базовая структура:
- handler / transport;
- service / usecase;
- repository;
- DB;
- migrations.

Принципы:
- HTTP-слой не должен содержать всю бизнес-логику;
- бизнес-правила должны быть выражены явно;
- БД должна защищать целостность данных;
- интерфейсы не должны превращаться в god-interface;
- транзакция применяется там, где несколько действий должны быть атомарны.

## Testing — 8/10

Нужно:
- go test;
- table-driven tests;
- unit tests;
- integration tests;
- E2E basics;
- mocks;
- race detector;
- тестировать бизнес-сценарии, а не только гнаться за coverage.

## Docker / Git / Unix — 7/10

Docker:
- Dockerfile;
- multi-stage;
- Docker Compose;
- volumes;
- networks;
- env;
- healthcheck;
- запуск БД + backend + migrations.

Git:
- branches;
- merge;
- rebase;
- reset;
- revert;
- stash;
- remote;
- cherry-pick basics.

Unix:
- processes;
- signals;
- permissions;
- environment;
- curl;
- grep;
- ps;
- ss;
- lsof;
- kill;
- bash basics.

---

# 3. Правила обучения

## Алгоритмы

Для каждой задачи:

1. Прочитать условие.
2. Назвать brute-force.
3. Посчитать сложность brute-force.
4. Найти узкое место.
5. Подобрать структуру данных / паттерн.
6. Написать алгоритм словами.
7. Назвать Time Complexity.
8. Назвать Space Complexity.
9. Выписать corner cases.
10. Только после этого писать Go-код.

Лимит:
- Easy: 15–20 минут;
- Medium: 25–35 минут.

Если решения нет:
1. посмотреть только hint;
2. ещё подумать;
3. посмотреть идею;
4. не копировать готовый код;
5. закрыть подсказку;
6. написать решение самостоятельно;
7. повторить задачу через 1–3 дня.

Статусы:
- 🟢 решил сам и могу объяснить;
- 🟡 решил с подсказкой / слишком долго;
- 🔴 не смог / пришлось смотреть решение.

Повторять в первую очередь 🟡 и 🔴.

## Теория

Нельзя ограничиваться чтением.

После каждой темы:
- 5–15 коротких вопросов;
- 1–3 мини-кода;
- объяснение темы своими словами;
- один практический сценарий.

Например после Mutex:
- что защищает Mutex;
- чем отличается от RWMutex;
- когда RWMutex хуже Mutex;
- где race;
- исправить race в коде;
- проверить через -race.

## Backend

Любую тему закреплять кодом.

Пример:
- изучили transaction;
- написать перевод денег между A и B;
- специально вызвать ошибку после списания;
- проверить rollback.

---

# 4. Первые 50 алгоритмических задач

Все решения писать на Go.

## HashMap / Array

1. Two Sum
2. Contains Duplicate
3. Valid Anagram
4. Majority Element
5. Missing Number
6. Intersection of Two Arrays
7. Group Anagrams
8. Top K Frequent Elements

## Two Pointers

9. Valid Palindrome
10. Merge Sorted Array
11. Remove Duplicates from Sorted Array
12. Move Zeroes
13. Two Sum II — Input Array Is Sorted
14. Container With Most Water
15. 3Sum

## Sliding Window

16. Maximum Average Subarray I
17. Best Time to Buy and Sell Stock
18. Longest Substring Without Repeating Characters
19. Longest Repeating Character Replacement
20. Permutation in String

## Prefix Sum

21. Running Sum of 1d Array
22. Find Pivot Index
23. Range Sum Query - Immutable
24. Subarray Sum Equals K

## Binary Search

25. Binary Search
26. Search Insert Position
27. First Bad Version
28. Search a 2D Matrix
29. Find Minimum in Rotated Sorted Array
30. Search in Rotated Sorted Array

## Stack

31. Valid Parentheses
32. Min Stack
33. Daily Temperatures

## Linked List

34. Reverse Linked List
35. Merge Two Sorted Lists
36. Linked List Cycle
37. Middle of the Linked List
38. Remove Nth Node From End of List

## Trees

39. Maximum Depth of Binary Tree
40. Invert Binary Tree
41. Same Tree
42. Binary Tree Level Order Traversal
43. Validate Binary Search Tree
44. Lowest Common Ancestor of a Binary Search Tree

## Graph / BFS / DFS

45. Flood Fill
46. Number of Islands
47. Clone Graph
48. Rotting Oranges

## Heap

49. Kth Largest Element in an Array
50. K Closest Points to Origin

## Календарь первых 50 алгоритмов — сентябрь 2026

Начало работы по этому расписанию: **7 сентября 2026** (со слов ученика). Расписание начинается с №9; результаты №1–8 и статусы остальных задач нужно уточнить, а не считать выполненными по дате.

| Дата | Задачи и повторения |
|---|---|
| 07.09 | №9–12: Valid Palindrome, Merge Sorted Array, Remove Duplicates from Sorted Array, Move Zeroes |
| 08.09 | №13–15: Two Sum II, Container With Most Water, 3Sum + 2 повтора |
| 09.09 | №16–20: Sliding Window |
| 10.09 | №21–24: Prefix Sum + 1 повтор |
| 11.09 | №25–27: Binary Search + 2 старые 🔴 задачи |
| 12.09 | №28–30: Binary Search + 2 повтора |
| 13.09 | №31–33: Stack + 2 случайные предыдущие задачи |
| 14.09 | №34–38: Linked List |
| 15.09 | №39–41: Trees + 2 повтора |
| 16.09 | №42–44: Trees + 2 случайные предыдущие задачи |
| 17.09 | №45–48: Graph / BFS / DFS |
| 18.09 | №49–50: Heap + 3 повтора |
| 19.09 | 5 случайных Easy/Medium, без заранее указанного паттерна |
| 20.09 | Mock: Easy + Medium за 60 минут, затем ещё 2 задачи |

До 20.09 включительно алгоритмическую нагрузку брать из этой таблицы: она заменяет общие нормы количества задач в дневных разделах, а не добавляется к ним. Даты — ориентир; незавершённые задачи не считать решёнными автоматически. С 21.09 переходить к смешанной практике, если первые 50 действительно завершены; иначе продолжить оставшиеся задачи.

После первых 50:
- не решать задачи блоками;
- брать смешанные Easy/Medium;
- обычно 3 новых + 2 повторных в день;
- 1 раз в неделю: Easy + Medium за 60 минут.

---

# 5. Учебный план по дням

Старт календарного плана: **7 сентября 2026**. День 1 — 07.09, день 2 — 08.09. Алгоритмы до 20.09 идут по календарю раздела 4.

Если отдельная тема уже пройдена — не пропускать день полностью:
- сделать дополнительные mixed algorithms;
- пройти вопросы ITBooster;
- сделать дополнительную практику по теме.

---

## Неделя 1 — 7–13 сентября
### Алгоритмы + Go Core

### День 1 — 07.09
Алгоритмы:
- продолжить текущий прогресс первых 50;
- ориентир: 4–5 задач;
- повторить минимум 1 вчерашнюю 🟡/🔴.

Теория:
- arrays vs slices;
- len/cap;
- append;
- underlying array;
- copy.

Практика:
- написать 5 коротких примеров изменения slice;
- проверить, когда два slice разделяют underlying array;
- посмотреть изменения len/cap.

Самопроверка:
- почему append иногда изменяет старый slice, а иногда нет?
- чем array отличается от slice?

### День 2 — 08.09
Алгоритмы:
- 4–5 задач;
- Two Pointers / текущий блок;
- 1 повтор.

Теория:
- pointers;
- pass-by-value в Go;
- pointer receiver vs value receiver;
- structs/methods.

Практика:
- 5 мини-примеров с указателями;
- изменить struct через pointer;
- объяснить, почему переназначение локальной копии указателя не меняет исходную переменную.

### День 3 — 09.09
Алгоритмы:
- 4–5 задач.

Теория:
- map;
- zero value;
- lookup;
- delete;
- iteration;
- strings / []byte / []rune.

Практика:
- частотный словарь;
- group-by через map;
- подсчёт символов Unicode.

### День 4 — 10.09
Алгоритмы:
- 4–5 задач.

Теория:
- interfaces;
- method set;
- type assertion;
- type switch;
- nil interface.

Практика:
- 3 реализации одного интерфейса;
- пример typed nil внутри interface.

### День 5 — 11.09
Алгоритмы:
- 4–5 задач.

Теория:
- error;
- wrapping;
- errors.Is;
- errors.As;
- defer;
- panic/recover.

Практика:
- цепочка ошибок repository → service → handler;
- корректное errors.Is;
- 3 эксперимента с defer.

### День 6 — 12.09
Алгоритмы:
- 4–5 задач + 2 повтора.

Теория:
- packages;
- modules;
- visibility;
- embedding;
- basic project layout.

Практика:
- создать маленький Go module;
- packages transport/service/repository.

### День 7 — 13.09
Алгоритмы:
- 5 mixed задач;
- 2 задачи повторно без подсказок.

Контроль Go Core:
- 30–45 минут вопросов;
- slices/maps/interfaces/pointers/errors;
- объяснять вслух.

---

## Неделя 2 — 14–20 сентября
### Concurrency

Каждый день:
- 4–5 алгоритмических контактов.

### День 8 — 14.09
Теория:
- goroutine;
- concurrency vs parallelism;
- scheduler basics;
- GOMAXPROCS.

Практика:
- последовательная и concurrent версия одной задачи;
- замер времени.

### День 9 — 15.09
Теория:
- unbuffered/buffered channels;
- blocking semantics.

Практика:
- producer → consumer;
- experiment with buffer size 0/1/10.

### День 10 — 16.09
Теория:
- close(channel);
- range over channel;
- кто должен закрывать канал;
- receive from closed channel.

Практика:
- pipeline из 2–3 stages.

### День 11 — 17.09
Теория:
- select;
- timeout;
- nil channel;
- multiple ready cases.

Практика:
- worker + timeout;
- fan-in из нескольких каналов.

### День 12 — 18.09
Теория:
- WaitGroup;
- Mutex;
- RWMutex;
- atomic basics.

Практика:
- concurrent counter;
- вариант с Mutex;
- вариант с atomic.

### День 13 — 19.09
Теория:
- race condition;
- deadlock;
- go test -race;
- go vet.

Практика:
- специально создать race;
- найти через -race;
- исправить;
- специально создать deadlock;
- объяснить причину.

### День 14 — 20.09
Алгоритмический mock:
- 1 Easy + 1 Medium;
- максимум 60 минут;
- вслух объяснить решение.

Go concurrency mock:
- 45–60 минут;
- goroutine/channel/select/waitgroup/mutex/race.

---

## Неделя 3 — 21–27 сентября
### Context + HTTP + networking

Алгоритмы каждый день:
- 3 новых + 2 повторных.

### День 15 — 21.09
- context.Context;
- WithCancel;
- WithTimeout;
- Deadline;
- Done;
- cancel.

Практика:
- функция с таймаутом;
- отмена нескольких worker.

### День 16 — 22.09
- goroutine leaks;
- cancellation propagation.

Практика:
- создать leak;
- исправить через context.

### День 17 — 23.09
- HTTP request/response;
- methods;
- status codes;
- JSON.

Практика:
- CRUD API in-memory.

### День 18 — 24.09
- handlers;
- validation;
- middleware;
- error mapping.

Практика:
- logging middleware;
- recovery middleware;
- request validation.

### День 19 — 25.09
- http.Client;
- client timeout;
- request context.

Практика:
- параллельно вызвать 10 URL;
- общий timeout.

### День 20 — 26.09
- keep-alive;
- connection reuse;
- Transport;
- MaxIdleConns;
- IdleConnTimeout.

Практика:
- reusable HTTP client.

### День 21 — 27.09
- OS signals;
- graceful shutdown.

Практика:
- корректно остановить HTTP server через signal + context.

---

## Неделя 4 — 28 сентября – 4 октября
### SQL / PostgreSQL

Алгоритмы:
- 3 новых + 2 повторных ежедневно.

### День 22 — 28.09
- SELECT;
- WHERE;
- ORDER BY;
- LIMIT/OFFSET.

Практика:
- минимум 20 запросов.

### День 23 — 29.09
- INNER JOIN;
- LEFT JOIN;
- multiple JOIN.

Практика:
- минимум 15 запросов с 3–4 таблицами.

### День 24 — 30.09
- GROUP BY;
- HAVING;
- SUM/COUNT/AVG/MIN/MAX.

Практика:
- минимум 15 запросов.

### День 25 — 01.10
- subqueries;
- CTE;
- window functions basics.

Практика:
- 10 CTE;
- 5 window queries.

### День 26 — 02.10
- ACID;
- transaction;
- COMMIT;
- ROLLBACK;
- isolation basics.

Практика:
- transfer A → B;
- намеренно упасть посередине;
- проверить rollback.

### День 27 — 03.10
- indexes;
- B-tree;
- composite indexes;
- selectivity.

Практика:
- таблица 100k+ строк;
- запросы до/после index.

### День 28 — 04.10
- EXPLAIN ANALYZE;
- Seq Scan;
- Index Scan.

Практика:
- минимум 10 запросов;
- обосновать индекс словами.

---

## Неделя 5 — 5–11 октября
# Avito Backend Trainee Assignment Winter 2025

Не смотреть официальный Solution/Advice до своей первой законченной версии.

Алгоритмы:
- ежедневно 3 новых + 2 повторных.

### День 29 — 05.10
- прочитать ТЗ;
- выписать use cases;
- выписать invariants;
- DB schema;
- API design.

### День 30 — 06.10
- architecture;
- handler → service/usecase → repository;
- migrations;
- PostgreSQL.

### День 31 — 07.10
- auth/JWT;
- user endpoints.

### День 32 — 08.10
- sendCoin;
- transaction;
- concurrent transfers;
- atomicity.

### День 33 — 09.10
- merchandise purchase;
- inventory/info.

### День 34 — 10.10
- unit tests;
- integration/E2E;
- race detector.

### День 35 — 11.10
- Docker Compose;
- README;
- clean startup;
- lint;
- затем открыть официальный Solution/Advice;
- выписать свои ошибки;
- составить refactoring list.

---

## Неделя 6 — 12–18 октября
# Avito Backend Trainee Assignment Spring 2025

Алгоритмы:
- 3 новых + 2 повторных ежедневно.

### День 36 — 12.10
- domain model;
- state transitions;
- DB schema.

### День 37 — 13.10
- registration/login;
- JWT;
- authorization/roles.

### День 38 — 14.10
- PVZ/reception/product endpoints.

### День 39 — 15.10
- LIFO deletion;
- close reception;
- transactions.

### День 40 — 16.10
- filters;
- pagination;
- indexes.

### День 41 — 17.10
- unit/integration tests;
- go test -race.

### День 42 — 18.10
- Docker;
- migrations;
- README;
- затем официальный Solution/Advice;
- refactoring.

К концу этого дня состояние должно быть:
**READY TO APPLY**.

---

## Неделя 7 — 19–25 октября
# Autumn 2025 + Go Platform

Алгоритмы:
- 5 mixed contacts в день.

### День 43 — 19.10
Autumn 2025:
- schema;
- architecture.

### День 44 — 20.10
- teams/users.

### День 45 — 21.10
- PR;
- reviewer assignment.

### День 46 — 22.10
- reassignment;
- idempotent merge;
- business invariants.

### День 47 — 23.10
- tests;
- Docker;
- README.

### День 48 — 24.10
Timed algorithms:
- Easy + Medium ≤60 минут;
- затем 3 review.

Go Platform:
- slices/maps/interfaces/pointers.

### День 49 — 25.10
Timed algorithms:
- Easy + Medium ≤60 минут;
- затем 3 review.

Go Platform:
- goroutines;
- channels;
- context;
- mutex;
- race.

---

## Неделя 8 — 26 октября – 1 ноября
# Autumn 2026 + system thinking

Алгоритмы:
- 5 mixed contacts в день.

### День 50 — 26.10
- CJM пользователя.

### День 51 — 27.10
- CJM заведения;
- domain model.

### День 52 — 28.10
- API;
- OpenAPI.

### День 53 — 29.10
- PostgreSQL schema;
- migrations.

### День 54 — 30.10
- основной Go service.

### День 55 — 31.10
- service заведения;
- integration.

### День 56 — 01.11
- C4 L2/L3;
- architecture defense;
- 60–90 минут объяснять проект вслух.

---

## Неделя 9 — 2–8 ноября
# Режим отбора

Алгоритмы:
- 3 новых + 2–3 повторных ежедневно.

### День 57 — 02.11
Programming Mock:
- Easy + Medium;
- 60 минут.

### День 58 — 03.11
Go Platform Mock:
- pointers;
- slices;
- maps;
- interfaces.

### День 59 — 04.11
Go Platform Mock:
- goroutines;
- channels;
- synchronization.

### День 60 — 05.11
Backend Mock:
- HTTP;
- context;
- networking.

### День 61 — 06.11
SQL Mock:
- JOIN;
- CTE;
- indexes;
- transactions;
- EXPLAIN.

### День 62 — 07.11
Test assignment defense:
- выбрать Winter/Spring;
- 60 минут отвечать «почему так?».

### День 63 — 08.11
Полное mock technical interview:
- 2–3 часа;
- Programming;
- Go Platform;
- backend questions;
- test assignment defense.

---

# 6. Старые тестовые Avito

Официальный репозиторий:
https://github.com/avito-tech/tech-internship/tree/main/Tech%20Internships/Backend

Изучать:

1. Backend Trainee Assignment Winter 2025
2. Backend Trainee Assignment Spring 2025
3. Backend Trainee Assignment Autumn 2025
4. Backend Trainee Assignment Autumn 2026

## Правильный порядок работы

1. Прочитать ТЗ.
2. НЕ смотреть чужие решения.
3. Выписать бизнес-инварианты.
4. Спроектировать API.
5. Спроектировать DB.
6. Спроектировать слои.
7. Сделать MVP.
8. Добавить tests.
9. Docker Compose.
10. README.
11. Проверить clean startup.
12. Только затем прочитать официальный Solution/Advice, если он есть.
13. Составить список ошибок.
14. Провести рефакторинг.

Цель не просто «сделать 4 проекта».

Цель:
- четыре раза пройти цикл проектирования;
- научиться принимать решения;
- научиться объяснять их;
- увидеть реальные ошибки участников прошлых наборов.

---

# 7. Материалы

## Avito

Backend internships:
https://github.com/avito-tech/tech-internship/tree/main/Tech%20Internships/Backend

Avito Tech Playbook:
https://github.com/avito-tech/playbook

## Go

A Tour of Go:
https://go.dev/tour/

Effective Go:
https://go.dev/doc/effective_go

Go documentation:
https://go.dev/doc/

context:
https://pkg.go.dev/context

sync:
https://pkg.go.dev/sync

net/http:
https://pkg.go.dev/net/http

Race detector:
https://go.dev/doc/articles/race_detector

Database access:
https://go.dev/doc/database/

## PostgreSQL

Tutorial:
https://www.postgresql.org/docs/current/tutorial.html

Indexes:
https://www.postgresql.org/docs/current/indexes.html

EXPLAIN:
https://www.postgresql.org/docs/current/using-explain.html

Transactions:
https://www.postgresql.org/docs/current/tutorial-transactions.html

## ITBooster

Главная база:
https://itbooster.ru/database/

Использовать в первую очередь:
- Algorithms;
- SQL;
- Networks;
- Docker;
- Git.

ITBooster использовать как:
- интервью-вопросы;
- self-check;
- повторение;
- поиск пробелов.

Не заменять им практику.

---

# 8. Как использовать Codex как наставника

## Главный принцип

Codex не должен превращаться в генератор готовых ответов.

Он должен:
1. дать теорию;
2. задать вопросы;
3. дать небольшую практику;
4. заставить ученика самому решить;
5. провести code review;
6. предложить улучшение;
7. вернуть к теме через несколько дней.

## Правила подсказок

Для алгоритмов:
- первые 20–30 минут не давать готовое решение;
- сначала задать наводящий вопрос;
- потом дать hint;
- потом объяснить идею;
- полный код давать только по явному запросу или после самостоятельной попытки.

Для backend:
- не генерировать целиком тестовое Avito;
- сначала просить ученика предложить:
  - API;
  - DB;
  - architecture;
  - invariants;
- затем ревьюить.

## Формат каждого учебного дня

Codex должен начать с:

1. «Что ты успел сделать вчера?»
2. «Какие задачи были 🟢/🟡/🔴?»
3. Короткий recall quiz вчерашней темы.
4. Теория сегодняшней темы.
5. Мини-практика.
6. Алгоритмические задачи.
7. Основная backend-практика.
8. Финальный self-check.
9. Список, что повторить завтра.

---

# 9. Шаблон дневного отчёта

Дата:

Время обучения:

## Алгоритмы

| Задача | Статус | Время | Ошибка / паттерн |
|---|---|---:|---|
| | 🟢/🟡/🔴 | | |

Что сегодня не увидел:
- 

## Теория

Изучено:
- 

Могу объяснить без подсказки:
- 

Нужно повторить:
- 

## Практика

Что написал:
- 

Что сломалось:
- 

Что исправил:
- 

## Вопросы

1.
2.
3.

## План повторения

Завтра:
-

Через 3 дня:
-

Через неделю:
-

---

# 10. Еженедельная проверка

Каждые 7 дней:

## Algorithms
- 1 Easy + 1 Medium;
- 60 минут;
- неизвестные заранее темы;
- обязательно объяснять вслух.

## Go
10–20 вопросов + 2 code review snippets.

## SQL
5 теоретических вопросов + 5 запросов.

## Backend
одна практическая задача:
- transaction;
- HTTP;
- concurrency;
- caching;
- idempotency;
- architecture.

## Итог недели

Оценить 0–5:
- Algorithms:
- Go Core:
- Concurrency:
- Context:
- HTTP:
- SQL:
- PostgreSQL:
- Testing:
- Docker:
- Git/Unix:
- Architecture:

Следующая неделя должна сильнее давить на самые слабые пункты.

---

# 11. Checklist «готов подаваться в Avito»

## Algorithms
- [ ] Easy стабильно ≤15–20 минут
- [ ] Medium часто ≤30–35 минут
- [ ] называю Big O
- [ ] называю Space Complexity
- [ ] сам нахожу corner cases
- [ ] могу объяснять вслух
- [ ] умею map/two pointers/sliding window/binary search/BFS/DFS

## Go
- [ ] slices
- [ ] maps
- [ ] pointers
- [ ] interfaces
- [ ] errors
- [ ] goroutines
- [ ] channels
- [ ] select
- [ ] WaitGroup
- [ ] Mutex/RWMutex
- [ ] race/deadlock
- [ ] context
- [ ] HTTP client/server
- [ ] graceful shutdown

## SQL
- [ ] multiple JOIN
- [ ] GROUP BY / HAVING
- [ ] CTE
- [ ] transactions
- [ ] ACID
- [ ] indexes
- [ ] composite indexes
- [ ] EXPLAIN ANALYZE

## Backend
- [ ] REST CRUD
- [ ] validation
- [ ] middleware
- [ ] JWT basics
- [ ] idempotency
- [ ] transactions
- [ ] handler/service/repository
- [ ] migrations

## Engineering
- [ ] unit tests
- [ ] integration tests
- [ ] go test -race
- [ ] Dockerfile
- [ ] Docker Compose
- [ ] Git
- [ ] Linux basics

## Avito practice
- [ ] Winter 2025
- [ ] Spring 2025
- [ ] Autumn 2025
- [ ] Autumn 2026 analysis/project
- [ ] могу защищать решения по каждому проекту

---

# 12. Если зимний набор откроется раньше

НЕ ждать завершения плана.

Сразу податься.

После подачи:
- продолжать алгоритмы;
- определить ближайший этап отбора;
- перераспределить подготовку.

Если выдали тестовое:
- алгоритмы оставить 30–60 минут в день;
- 70–80% времени отдавать тестовому;
- не позволять AI делать проект вместо кандидата;
- любой AI-generated код кандидат должен понимать и уметь защитить.

---

# 13. Главная стратегия

Нужно прийти к отбору не с состоянием:

> «Я прочитал Go и решил несколько LeetCode».

А с состоянием:

> «Я умею решать Easy/Medium, понимаю Go concurrency/context, пишу backend с PostgreSQL, понимаю транзакции и индексы, умею тестировать и контейнеризировать сервис, прошёл несколько прошлых тестовых Avito и могу защитить каждое своё решение».

Это и есть основная цель программы.

