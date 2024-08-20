create table users(
id bigserial not null Primary key,
name varchar(100) not null,
surname varchar(100) not null,
email varchar(50) not null,
password varchar(100) not null
);


create table groups(
id bigserial not null Primary key,
name varchar(10) not null,
is_archive bool not null
);

create table students(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
group_id bigint not null,
FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

create table seminarians(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

create table disciplines(
id bigserial not null Primary key,
name varchar(50) not null,
name_en varchar(50) not null,
seminar_visiting_mark bigint,
lesson_visiting_mark bigint,
exam_mark bigint not null
);

create table curriculums(
discipline_id bigint not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE,
group_id bigint not null,
FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
is_archive bool not null
);

create table lecturers(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

Create table sections(
id bigserial not null Primary key,
name varchar(50) not null,
name_en varchar(50) not null,
discipline_id bigint not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE
);

create table tests(
id bigserial not null Primary key,
name varchar(150) not null,
name_en varchar(50) not null,
minutes_duration int not null,
task_description varchar(200),
task_description_en varchar(200),
default_mark int not null
);

Create table section_test(
section_id bigint not null,
FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE,
test_id bigint not null,
FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE
);

Create table questions(
id bigserial not null Primary key,
name varchar(500) not null,
name_en varchar(150),
is_variable int not null
);

CREATE TABLE answers(
id bigserial not null Primary key,
name varchar(500) not null,
name_en varchar(500),
is_right bool not null,
question_id bigint not null,
FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

Create table themes(
id bigserial not null Primary key,
name varchar(150) not null,
weight bigint not null
);

Create table theme_questions(
theme_id bigint not null,
FOREIGN KEY (theme_id) REFERENCES themes(id) ON DELETE CASCADE,
question_id bigint not null,
FOREIGN KEY (question_id) REFERENCES questions(id) ON DELETE CASCADE
);

create table external_laboratory_works(
id bigserial not null Primary key,
name varchar(50) not null,
name_en varchar(50) not null,
task_description varchar(200),
task_description_en varchar(200),
link text not null,
token text
);

create table laboratory_works (
id bigserial not null Primary key,
default_mark int not null,
external_laboratory_id bigint not null,
FOREIGN KEY (external_laboratory_id) REFERENCES external_laboratory_works(id) ON DELETE CASCADE
);

Create table section_lab(
section_id bigint not null,
FOREIGN KEY (section_id) REFERENCES sections(id) ON DELETE CASCADE,
laboratory_id bigint not null,
FOREIGN KEY (laboratory_id) REFERENCES laboratory_works(id) ON DELETE CASCADE
);


create table tests_date(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
test_id bigint not null,
FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE,
closed_date BIGINT NOT NULL,
is_done bool NOT NULL
);


create table laboratory_date(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
laboratory_id bigint not null,
FOREIGN KEY (laboratory_id) REFERENCES laboratory_works(id) ON DELETE CASCADE,
closed_date BIGINT NOT NULL,
is_done bool NOT NULL
);


create table test_marks(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
test_id bigint not null,
FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE,
mark int
);

create table test_themes(
test_id bigint not null,
FOREIGN KEY (test_id) REFERENCES tests(id) ON DELETE CASCADE,
theme_id bigint not null,
FOREIGN KEY (theme_id) REFERENCES themes(id) ON DELETE CASCADE,
count int not null
);

create table laboratory_marks(
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
laboratory_id bigint not null,
FOREIGN KEY (laboratory_id) REFERENCES laboratory_works(id) ON DELETE CASCADE,
mark int
);

create table lessons(
id bigserial not null Primary key,
discipline_id bigint not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE,
name varchar(30) not null
);

create table lessons_date(
id bigserial not null Primary key,
date bigint not null,
lesson_id bigint not null,
FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE,
group_id bigint not null,
FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE
);

create table seminars(
id bigserial not null Primary key,
discipline_id bigint not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE,
name varchar(30) not null,
group_id bigint,
FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
date bigint not null
);


create table lesson_visiting(
lesson_id bigint not null,
FOREIGN KEY (lesson_id) REFERENCES lessons(id) ON DELETE CASCADE,
is_absent bool not null,
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id)
);


create table seminar_visiting(
seminar_id bigint not null,
FOREIGN KEY (seminar_id) REFERENCES seminars(id) ON DELETE CASCADE,
is_absent bool not null,
user_id bigint not null,
FOREIGN KEY (user_id) REFERENCES users(id)
);


Create table digital_materials(
id bigserial not null Primary key,
name varchar(50) not null,
name_en varchar(50) not null,
description varchar(400),
description_en varchar(400)
);

CREATE TABLE digital_discipline(
digital_material_id bigint not null,
FOREIGN KEY (digital_material_id) REFERENCES digital_materials(id) ON DELETE CASCADE,
discipline_id bigint not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE
);

CREATE TABLE file_path(
id BIGSERIAL not null Primary key,
digital_material_id bigint not null,
FOREIGN KEY (digital_material_id) REFERENCES digital_materials(id) ON DELETE CASCADE,
name varchar(150) not null
);

create table seminarian_groups(
user_id bigint not null,
group_id bigint,
discipline_id bigint not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE,
FOREIGN KEY (group_id) REFERENCES groups(id) ON DELETE CASCADE,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

create table exams(
user_id bigint not null,
discipline_id bigint not null,
mark int not null,
FOREIGN KEY (discipline_id) REFERENCES disciplines(id) ON DELETE CASCADE,
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

ALTER TABLE users ADD CONSTRAINT unique_email UNIQUE (email);
ALTER TABLE theme_questions ADD CONSTRAINT unique_theme_questions UNIQUE (theme_id, question_id);
ALTER TABLE groups ADD CONSTRAINT unique_name UNIQUE (name);
ALTER TABLE disciplines ADD CONSTRAINT unique_discipline UNIQUE (name);
