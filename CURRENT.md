# Текущее состояние обучения

## Актуальная точка — день 16 завершён, 02.10.2026

- Статус: завершён. Алгоритмы в день 16 были исключены по явной просьбе ученика; незавершённые live coding и алгоритмический mock дня 14 остаются перенесёнными и не отмечены выполненными.
- Последний отчёт: `sessions/day-16-2026-10-02.md`; предыдущий: `sessions/day-16-2026-10-01.md`.
- Текущий этап: день 16 завершён — goroutine leaks и cancellation propagation.
- Текущее задание: нет; следующий учебный день — день 17, HTTP request/response и JSON.
- Что уже сделано: `practice/day-16/exercise-02`: cancellation-aware `Double` с normal-flow и cancellation-тестами. `gofmt` и `go vet` чистые; `go test -race -count 100 -timeout 5s .` passed (ранее также `-count 500`).
- Пробелы: тестовая синхронизация, различение ожидаемой ветки `select` и timeout; закрепить `Done`/`Err`, причины отмены и отменяемые receive/send.
- Последние оценки: Go context / cancellation — 3/5; Go concurrency / worker pool — 2/5 (день 15).
- Повторить в день 17: cancellation receive/send, ownership закрытия output, смысл `finished`, `Done`/`Err` и причина отмены.
- Следующий шаг: начать день 17 с recall повторений; алгоритмическую часть согласовать отдельно, так как в день 16 она была исключена.
