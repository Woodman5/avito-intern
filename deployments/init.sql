CREATE TABLE IF NOT EXISTS public.user_moneys
(
    uuid uuid not null,
    amount bigint default 0 not null
);

create unique index user_moneys_uuid_uindex
    on public.user_moneys (uuid);

alter table public.user_moneys
    add constraint user_moneys_pk
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