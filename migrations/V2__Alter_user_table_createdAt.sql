alter table "User"
add column  createdAt timestamp with time zone not null default CURRENT_TIMESTAMP,
    add column updatedAt timestamp with time zone not null default CURRENT_TIMESTAMP,
    add column deletedAt timestamp with time zone;