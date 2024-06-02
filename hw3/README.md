Реализуйте интерфейс Formatter с методом Format, который возвращает
отформатированную строку.

Определите структуры, удовлетворяющие интерфейсу Formatter: 
1. обычный текст(как есть)
2. жирным шрифтом(** **)
3. код(` `)
4. курсив(_ _)

Опционально: иметь возможность задавать цепочку модификаторов
chainFormatter.AddFormatter(plainText)
chainFormatter.AddFormatter(bold)
chainFormatter.AddFormatter(code)