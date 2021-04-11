create table "Readers"
(
    reader_id int not null
        constraint readers_pk
            primary key,
    occupation_id int NOT NULL,
    city_id int NOT NULL,
    reader_name char (50) NOT NULL,
    reader_surname char (50) NOT NULL,
    debtor boolean NOT NULL
);
insert into "Readers" values (1, 1, 1, 'Dima', 'Putkou', false);
insert into "Readers" values (2, 2, 2, 'Ivan', 'Ivanov', true);
insert into "Readers" values (3, 2, 1, 'Danik', 'Domaskanou', true);
insert into "Readers" values (4, 5, 3, 'Artem', 'Menshikou', false);
insert into "Readers" values (5, 4, 5, 'Nikita', 'Miladouski', true);



create table "Books"
(
    book_id int not null,
    author_id int not null,
    publisher_id int not null,
    name_of_book char(50) not null,
    year_of_publication date not null,
    book_volume int not null,
    number int not null
);
create unique index books_book_id_uindex
	on "Books" (book_id);
alter table "Books"
    add constraint books_pk
        primary key (book_id);
insert into "Books" values (1,1,1, 'Van Helsing', '16-05-2019',10, 4);
insert into "Books" values (2,2,4, 'Romeo and Juliet', '16-05-2018',10, 4);
insert into "Books" values (3,5,2, 'Three Musketeers', '16-05-2017',10, 4);
insert into "Books" values (4,3,1, 'Captains daughter', '16-05-2016',10, 4);
insert into "Books" values (5,4,3, 'Dubrovsky', '16-05-2015',10, 4);



create table "Authors"
(
    author_id int not null
        constraint authors_pk
            primary key,
    name_of_author char(50) not null,
    surname char(50) not null
);
INSERT INTO "Authors" VALUES (1, 'Alexander','Pushkin');
INSERT INTO "Authors" VALUES (2, 'Lev','Tolstoy');
INSERT INTO "Authors" VALUES (3, 'Alexander','Blok');
INSERT INTO "Authors" VALUES (4, 'Nikolay','Nekrasov');
INSERT INTO "Authors" VALUES (5, 'Anton','Chekhov');



create table "Occupation"
(
    occupation_id int not null
        constraint occupation_pk
            primary key,
    name_of_occupation char(50) not null
);
insert into "Occupation" values (1, 'student');
insert into "Occupation" values (2, 'schoolboy');
insert into "Occupation" values (3, 'worker');
insert into "Occupation" values (4, 'pensioner');
insert into "Occupation" values (5, 'vip');


create table "Cities of people"
(
    city_id int not null
        constraint "cities of people_pk"
            primary key,
    name_of_city char(50) not null
);
insert into "Cities of people" values (1, 'Gomel');
insert into "Cities of people" values (2, 'Minsk');
insert into "Cities of people" values (3, 'Vitebsk');
insert into "Cities of people" values (4, 'Moscow');
insert into "Cities of people" values (5, 'Peter');



create table "Issuing a book"
(
    issuing_id int not null
        constraint "issuing a book_pk"
            primary key,
    reader_id int not null,
    book_id int not null,
    issue_date_of_the_book date not null,
    book_return_date date not null
);
insert into "Issuing a book" values (1, 1, 1, '08-09-2021', '10-10-2021');
insert into "Issuing a book" values (2, 2, 2, '06-09-2020', '10-10-2021');
insert into "Issuing a book" values (3, 3, 2, '08-05-2021', '10-10-2021');
insert into "Issuing a book" values (4, 1, 3, '08-03-2021', '10-11-2021');
insert into "Issuing a book" values (5, 4, 5, '08-02-2021', '10-12-2021');


create table "Publishers"
(
    publisher_id int not null
        constraint publishers_pk
            primary key,
    city_id int not null,
    name_of_publisher char(50) not null
);
insert into "Publishers" values (1, 1, 'Moscow');
insert into "Publishers" values (2, 2, 'Vitebsk');
insert into "Publishers" values (3, 3, 'Gomel');
insert into "Publishers" values (4, 4, 'Peter');
insert into "Publishers" values (5, 5, 'Peter');


create table "Cities of publishers"
(
    city_id int not null
        constraint "cities of publishers_pk"
            primary key,
    name_of_city char(50) not null
);
insert into "Cities of publishers" values (1, 'Moscow');
insert into "Cities of publishers" values (2, 'Vitebsk');
insert into "Cities of publishers" values (3, 'Gomel');
insert into "Cities of publishers" values (4, 'Peter');
insert into "Cities of publishers" values (5, 'Grodno');


