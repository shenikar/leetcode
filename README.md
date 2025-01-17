# LeetCode Solutions Repository

![LeetCode](https://img.shields.io/badge/LeetCode-Practice-orange)
![Language](https://img.shields.io/badge/Language-Go-blue)
![Contributions](https://img.shields.io/badge/Contributions-Welcome-green)

Этот репозиторий содержит мои решения задач с [LeetCode](https://leetcode.com/). Решения реализованы на языке **Go**, с использованием оптимальных алгоритмов и структур данных.

## 📁 Структура репозитория

Репозиторий организован следующим образом:

- **`/easy`**: Решения задач уровня "Easy".
- **`/medium`**: Решения задач уровня "Medium".
- **`/hard`**: Решения задач уровня "Hard".

Каждый файл содержит:
- Имя задачи.
- Само решение.

## 🚀 Как использовать

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/shenikar/leetcode.git
   ```

2. Перейдите в папку с решением:
    ```bash
    cd leetcode/<уровень сложности>
    ```

3. Запустите решение с помощью Go:
    ```bash
    go run <имя файла>.go
    ```

## 🛠️ Особенности

- Решения написаны с упором на:
  - Читаемость.
  - Производительность.
- Используются структуры данных:
  - Массивы.
  - Хэш-таблицы.
  - Деревья.
- Применяются встроенные функции Go для упрощения и оптимизации кода.

## 📝 Пример задачи

### Min Cost Climbing Stairs

**Ссылка на задачу**: [LeetCode](https://leetcode.com/problems/min-cost-climbing-stairs/)

**Описание**:  
Задача заключается в нахождении минимальной стоимости достижения вершины лестницы. На каждом шаге вы можете либо подняться на одну ступень, либо перепрыгнуть через одну.

**Код**:
```go
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    dp := make([]int, n+1)
    for i := 2; i <= n; i++ {
        dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
    }
    return dp[n]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
```

## 👩‍💻 Как вносить вклад

1. **Форкните репозиторий**  
   Создайте копию репозитория, нажав кнопку `Fork` на GitHub.

2. **Склонируйте репозиторий**  
   Скачайте репозиторий на ваш компьютер с помощью команды:
   ```bash
   git clone https://github.com/your-username/leetcode.git
   ```
3. **Создайте новую ветку.**

    Переключитесь на новую ветку для добавления своего решения:
    ```bash
    git checkout -b feature/new-solution
    ```
4. **Добавьте свое решение**
    
    Поместите ваш код в соответствующую папку (/easy, /medium, /hard).
5. **Закоммитьте изменения**
    ```bash
    git add .
    git commit -m "Добавлено решение задачи <название задачи>"
    ```
6. **Запушьте изменения**
    
    Отправьте изменения в ваш форк:
    ```bash
    git push origin feature/new-solution
    ```
7. **Создайте Pull Request**

    Зайдите на страницу вашего форка на GitHub и нажмите кнопку New Pull Request, чтобы предложить свои изменения.

---

## 💡 Контакты

Если у вас есть вопросы, предложения или замечания, свяжитесь со мной одним из следующих способов:

- **Через GitHub Issues**:  
  [Открыть новый Issue](https://github.com/shenikar/leetcode/issues) для обсуждения проблемы или предложения улучшений.

- **Email**:  
  Напишите мне на почту: <ishmatov.ilfar@gmail.com>.

---

### ⭐️ Если вам понравился этот репозиторий, поставьте ⭐️!

Поддержка репозитория помогает мотивировать на создание новых решений и улучшений.



