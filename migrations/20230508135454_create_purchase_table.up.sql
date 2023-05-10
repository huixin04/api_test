CREATE EXTENSION IF NOT EXISTS "pg_trgm";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE purchase(
    username VARCHAR NOT NULL,
    purchase_id uuid default uuid_generate_v4()  NOT NULL primary key,
    company_name VARCHAR NOT NULL,
    department VARCHAR NOT NULL,
    product_name VARCHAR NOT NULL,
    purchase_quantity int NOT NULL,
    price int NOT NULL,
    created_by VARCHAR not NULL,
    created_at timestamp default now()  NOT NULL,
    updated_at timestamp null ,
    updated_by VARCHAR null,
    is_deleted bool default false not null
);

-- create unique index purchase_purchase_id_uindex
-- on purchases (purchase_id)

CREATE or replace function purchase_code()
returns text as
$$
declare
    old_id text :=(select  purchase_id  from purchase order by purchase_id desc limit 1);
    id_number char(4) :='0001';
    
    new_id text ;
    num integer;
begin
    if old_id is null then
        new_id:='p'||id_number;
        return new_id;
    end if;
    
   
        num :=cast(right(old_id,4) as integer)+1;
        id_number:=
        case
            when num<10 then '000'||num
            when num<100 then '00'||num
            when num<1000 then '0'||num
            when num<10000 then cast(num as text)
        end;
    
    
    new_id:='p'||id_number;
    return new_id;
end; 
$$
language 'plpgsql';

select purchase_id()
