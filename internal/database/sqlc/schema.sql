create TABLE IF NOT EXISTS users(
      id INTEGER PRIMARY KEY,
      username varchar(255) UNIQUE NOT NULL ,
      password_hash varchar(255) NOT NULL ,
      email varchar(255),
      created_at date NOT NULL
);

create TABLE IF NOT EXISTS tasks(
      id INTEGER PRIMARY KEY ,
      title varchar(500) NOT NULL ,
      description varchar(5000) NOT NULL ,
      status varchar(255) NOT NULL ,
      created_at date NOT NULL ,
      updated_at date ,
      user_id INTEGER NOT NULL ,
      FOREIGN KEY(user_id) REFERENCES users(id)
);