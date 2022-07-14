CREATE DATABASE faraway;

\c faraway;

CREATE TABLE users (
  uid serial PRIMARY KEY,
  name varchar(120) NOT NULL,
  email varchar(50) NOT NULL,
  phone varchar(20) NOT NULL,
  gender varchar(7) NOT NULL
);

INSERT INTO users (name, email, phone, gender) VALUES
  ('Sunny Xo', 'sonne@google.com', '111-111-111', 'male'),
  ('Nadya Zulu', 'nadyanadya@mail.com', '22222222', 'female'),
  ('Vincent Vega', 'vincent_vega@api.com', '6543210', 'male'),
  ('Michelle P', 'michelle@gmail.com', '333-333-333', 'female'),
  ('John Doe', 'famouse@mail.ru', '444-444-444', 'male'),
  ('Mr. Long Long Long Cat', 'loooooooooooooooooongest@kitty.com', '77-777-777', 'female'),
  ('Ilon Dusk', 'ilonchik@yandex.ru', '555-555-555', 'male'),
  ('Ms. Mia Wallace', 'miumiumiu@microsoft.com', '233-322', 'female')