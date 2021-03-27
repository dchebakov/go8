create table books
(
    book_id bigint unsigned auto_increment,
    title varchar(255) null,
    published_date timestamp not null,
    image_url varchar(255) null,
    description text not null,
    created_at datetime null default now(),
    updated_at datetime null default now(),
    deleted_at datetime null,
    constraint books_pk
        primary key (book_id)
);

