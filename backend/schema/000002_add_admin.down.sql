DELETE FROM lecturers WHERE user_id = (SELECT id FROM users WHERE email = 'anna7909966@yandex.ru');

DELETE FROM users WHERE email = 'anna7909966@yandex.ru';
