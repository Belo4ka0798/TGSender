\c tgsender;

create type state as enum('Open', 'Closed');

create table questions
(
    id      bigserial primary key unique not null,
    email   text                      not null,
    header  text                      not null,
    message text                      not null,
    answer  text default 'Empty',
    date    time default current_time,
    status state default 'Open'
);