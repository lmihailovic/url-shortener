# URL Shortener (golang)
The second project in the [Gophercises](https://gophercises.com/) course for
practicing and mastering Go.

I need to be honest - this one gave me some headaches while I wrapped my head
around how Go handles web routing (differences between `http.Handle`, `http.Handler`, `http.HandleFunc`, `http.HandlerFunc`,
and what even is a `mux`), and worst of all: YAML parsing.

Even worse than YAML is the fact that the code looks dead simple to me now -
an interesting contrast to the feeling I had until my synapses connected in
such a way to allow me to write the damn thing. There were several _Eureka!_
moments indeed.

## How it works
Without arguments, the program uses its default map of paths and URLs to which
they point. Optionally, a JSON or YAML file can be passed using the `-f` flag,
or a database can be used via the `-db` flag being set to `true`.

Upon starting the program, visit `localhost:8080`, and specify a path which
redirects to a URL. For example, path `localhost:8080/yaml-godoc` will lead
to `https://godoc.org/gopkg.in/yaml.v2`.

## Arguments
```
 -db bool
        use database
 -f string
        path to file with data
```

### Creating a test database
1. Start the docker container
```
% docker-compose up -d
```
2. Connect to MariaDB
```
% docker exec -it mariadb mariadb -u root -p
```
3. Enter the password `root` when prompted
4. Inside MariaDB shell, create the database
```
MariaDB [(none)]> CREATE DATABASE urlshort;
```
5. Connect to the database with
```
MariaDB [(none)]> USE urlshort;
```
6. Create a table with paths and URLs they redirect to
```
CREATE TABLE redirect (
    id INT NOT NULL AUTO_INCREMENT,
    path VARCHAR(255) NOT NULL,
    url VARCHAR(255) NOT NULL,
    PRIMARY KEY (id)
);
```
7. Insert values into table
```
INSERT INTO redirect (path, url) VALUES
    ('/short', 'https://go.dev/'),
    ('/repo', 'https://github.com/lmihailovic/url-shortener');
```
8. Test that data was inserted correctly
```
MariaDB [urlshort]> SELECT * FROM redirect;
```
## Bonus
- [x] Update the main/main.go source file to accept a YAML file as a flag and then
load the YAML from a file rather than from a string.
- [x] Build a JSONHandler that serves the same purpose, but reads from JSON data.
- [x] Build a Handler that doesn't read from a map but instead reads from a database.
Whether you use BoltDB, SQL, or something else is entirely up to you.