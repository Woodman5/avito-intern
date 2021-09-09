CREATE TABLE IF NOT EXISTS public.usermoney
(
    uuid uuid not null,
    amount bigint default 0 not null
);

create unique index usermoney_uuid_uindex
    on public.usermoney (uuid);

alter table public.usermoney
    add constraint usermoney_pk
        primary key (uuid);

create table public.transactions
(
    uuid uuid not null,
    createdat timestamp not null,
    useruuid uuid not null,
    transactiontype smallint not null,
    amount bigint default 0 not null,
    balance bigint default 0 not null,
    source character varying(100) not null,
    reason text not null
);

create unique index transactions_uuid_uindex
    on transactions (uuid);

alter table transactions
    add constraint transactions_pk
        primary key (uuid);