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

INSERT INTO public.usermoney (uuid, amount) VALUES
    ('e6030c5a-219b-451c-bef2-344c128ef08f', 7400000),
    ('d791ed34-b8db-4e4f-a9a7-228f56c8ac48', 45000),
    ('58bb92ab-4051-4baa-bc0a-52bad53458d6', 857400),
    ('56a36a84-fb53-4bfa-a2dd-4307c4b3c981', 203700);

INSERT INTO public.transactions (uuid, createdat, useruuid, transactiontype, amount, balance, source,reason) VALUES
    ('40f27f19-b0f5-49db-a2d1-76a787ac69f6', '2021-09-01 04:11:58.000000', 'e6030c5a-219b-451c-bef2-344c128ef08f', 1, 23000, 7410000, 'Vasya', 'Coffee'),
    ('e900108b-6761-47f3-81ed-afbf874f98f7', '2021-09-05 02:10:54.000000', '56a36a84-fb53-4bfa-a2dd-4307c4b3c981', 1, 200000, 203700, 'Tolyan', 'Car repair'),
    ('ac592f97-6045-45ef-a7af-ecd5699aaffe', '2021-09-09 12:31:23.000000', 'e6030c5a-219b-451c-bef2-344c128ef08f', 0, 10000, 7400000, 'Boss', 'Fee'),
    ('7361fb43-36c0-428c-98d8-6737cae73201', '2021-09-10 03:47:08.000000', 'd791ed34-b8db-4e4f-a9a7-228f56c8ac48', 0, 500000, 45000, 'Wife', 'Some goods');
