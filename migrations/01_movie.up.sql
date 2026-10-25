create table movies(
  id serial primary key,
  tmdbid Int not null,
  imdbid varchar(11) not null,
  details jsonb not null
);
