CREATE TABLE IF NOT EXISTS users
(
  id       serial not null unique,
  username varchar(255) not null unique,
  name     varchar(255) not null,
  password varchar(255) not null,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS advertisements
(
  id          serial not null unique,
  title       varchar(255) not null,
  user_id     int references users (id) on delete cascade not null,
  description varchar(255) not null,
  images      varchar[] DEFAULT null,
  created_at  timestamp NOT NULL DEFAULT NOW(),
  updated_at  timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS comments
(
  id         serial not null unique,
  message    varchar(255) not null,
  user_id    int references users(id) not null,
  ads_id     int references advertisements(id) on delete cascade not null,
  created_at timestamp NOT NULL DEFAULT NOW(),
  updated_at timestamp NOT NULL DEFAULT NOW()
);