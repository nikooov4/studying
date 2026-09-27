# Текущее состояние обучения

## Актуальная точка — дни 11 и 12 завершены, 27.09.2026

- Статус: завершён. Отчёты: `sessions/day-11-2026-09-27.md`, `sessions/day-12-2026-09-26.md`.
- Последняя завершённая по плану тема: день 12 — `WaitGroup`, `Mutex`, `RWMutex`, основы `atomic`; concurrent counter с Mutex и atomic.
- Завершение дня 11: `select`, timeout, nil channel, worker + timeout и fan-in; практика в `practice/day-11/exercise-01` и `exercise-02` зачтена, `go vet` чист.
- Алгоритмы: остаются в долге; не считать долг выполненным.
- Текущий этап: итог сохранён.
- Текущее задание: нет.
- Последний отчёт: `sessions/day-11-2026-09-27.md`; день 11 завершён после дня 12.
- Итог дня 12: concurrent counter с Mutex и atomic дал 10000; `go vet` и race detector без диагностик. Оценка синхронизации — 3/5.
- Актуальные повторы: день 13 — receive vs send после `close`, `range` pipeline; также `WaitGroup.Add` до запуска goroutine, ровно один `Done`, невозможность upgrade `RLock` → `Lock`; день 15 — CPU-bound эксперимент; день 17 — cancellation pipeline через context.
- Следующий шаг: начать день 13 с recall закрытия каналов, fan-in и правил `WaitGroup`; алгоритмический долг вести отдельно.
