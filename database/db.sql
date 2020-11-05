drop table if exists users;
drop sequence if exists users_id_seq;
create sequence users_id_seq;
create table users (
    id bigint not null default nextval('users_id_seq') primary key,
    name text,
    username text,
    password text
);

drop table if exists books;
drop sequence if exists books_id_seq;
create sequence books_id_seq;
create table books (
    id bigint not null default nextval('books_id_seq') primary key,
    name text
);

drop table if exists user_books;
drop sequence if exists user_books_id_seq;
create sequence user_books_id_seq;
create table user_books (
    id bigint not null default nextval('user_books_id_seq') primary key,
    user_id bigint,
    book_id bigint,
    foreign key (user_id) references users(id),
    foreign key (book_id) references books(id)
);