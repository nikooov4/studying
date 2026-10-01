# Текущее состояние обучения

## Актуальная точка — день 16 в процессе, 01.10.2026

- Статус: в процессе; алгоритмы исключены по явной просьбе ученика. Незавершённые live coding и алгоритмический mock дня 14 остаются перенесёнными и не отмечены выполненными.
- Последний отчёт: `sessions/day-16-2026-10-01.md`; предыдущий: `sessions/day-16-2026-09-30.md`.
- Текущий этап: backend-практика — каркас cancellation-aware pipeline stage.
- Текущее задание: создать `practice/day-16/exercise-02` с `go.mod` и объявить API `Double(ctx, in) (out, finished)`; реализацию пока не писать.
- Что уже сделано: recall и теория дня 16; `StartSend` с отменяемой отправкой и двумя тестами. `gofmt -d` без diff; `go test -count 50 -timeout 5s .` passed. Lifecycle pipeline спроектирован.
- Пробелы: `Done`/`Err`, причины отмены и отменяемые send требовали адресных подсказок; надо закрепить cancellation-путь тестом.
- Последние оценки: Go context — 2/5; Go concurrency / worker pool — 2/5 (день 15).
- Повторить в день 16: `Done`/`Err`, причина отмены vs сигнал, cancellation receive/send, ownership закрытия output, Time `O(J)` / Space `O(W)`.
- Следующий шаг: проверить lifecycle pipeline stage, затем выдать фиксированное ТЗ backend-практики.
