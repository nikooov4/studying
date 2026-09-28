# Текущее состояние обучения

## Актуальная точка — день 13 завершён, 28.09.2026

- Статус: завершён. Отчёты: `sessions/day-11-2026-09-27.md`, `sessions/day-12-2026-09-26.md`, `sessions/day-13-2026-09-28.md`.
- Последняя завершённая по плану тема: день 13 — race condition, deadlock, `go test -race`, `go vet`.
- Завершение дня 11: `select`, timeout, nil channel, worker + timeout и fan-in; практика в `practice/day-11/exercise-01` и `exercise-02` зачтена, `go vet` чист.
- Алгоритмы: остаются в долге; не считать долг выполненным.
- Текущий этап: итог сохранён.
- Текущее задание: нет.
- Последний отчёт: `sessions/day-13-2026-09-28.md`.
- Итог дня 13: race на счётчике создан и обнаружен через `-race`, затем исправлен `Mutex`; намеренный channel deadlock воспроизведён и объяснён. Оценка concurrency diagnostics — 3/5.
- Актуальные повторы: день 15 — receive vs send после `close`, `range` pipeline; `WaitGroup.Add` до запуска goroutine; невозможность upgrade `RLock` → `Lock`; CPU-bound эксперимент; день 17 — cancellation pipeline через context.
- Следующий шаг: день 14 — алгоритмический и Go concurrency mock по плану; алгоритмический долг вести отдельно.
