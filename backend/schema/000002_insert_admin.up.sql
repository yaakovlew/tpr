INSERT INTO users (name, surname, email, password)
SELECT 'Анна', 'Тихомирова', 'anna7909966@yandex.ru', '$2a$10$exbLpg4c85R02cWZ8WqrLejPuW4IUOKl5eOmOCfcRR0JiqqYK47o6'
WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = 'anna7909966@yandex.ru');

INSERT INTO lecturers (user_id)
SELECT id FROM users WHERE email = 'anna7909966@yandex.ru'
AND NOT EXISTS (SELECT 1 FROM lecturers WHERE user_id = (SELECT id FROM users WHERE email = 'anna7909966@yandex.ru'));
