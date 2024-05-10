create table public.car(
                                id int primary key generated always as identity,
                                reg_num varchar(100) unique,
                                mark varchar(100),
                                model varchar(100)
);


