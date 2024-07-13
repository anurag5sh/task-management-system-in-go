create TABLE users(
      id INTEGER PRIMARY KEY,
      username varchar(255) UNIQUE ,
      password_hash varchar(255),
      email varchar(255),
      created_at date
);

create TABLE tasks(
      id INTEGER PRIMARY KEY ,
      title varchar(500),
      description varchar(5000),
      status varchar(255),
      created_at date,
      updated_at date,
      user_id INTEGER,
      FOREIGN KEY(user_id) REFERENCES users(id)
);