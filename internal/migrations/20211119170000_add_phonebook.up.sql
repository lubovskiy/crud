CREATE TABLE phonebook (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    create_date TIMESTAMP DEFAULT (now() AT TIME ZONE 'UTC'),
    update_date TIMESTAMP DEFAULT (now() AT TIME ZONE 'UTC')
);

CREATE FUNCTION update_date_proc()
    returns trigger as
$body$
begin
    new.update_date = now() AT TIME ZONE 'UTC';
    return new;
end;
$body$
    language plpgsql volatile security definer;

CREATE trigger update_date
    BEFORE update
    ON phonebook
    for each row
execute procedure update_date_proc();
