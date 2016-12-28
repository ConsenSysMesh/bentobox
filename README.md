## Event Crawler

It reads the JSON RPC API of an ethereum node and dumps its results into
a PSQL database.

### Run a PSQL database using docker

Just use the command

```
make run-psql
```

Will mount the dir `$HOME/.psql` and have a DB for you.

You need to setup your database if it is the first time using it.
Use the instructions below.

### Setup the database

You can have your database already, or you can use docker as above.
This script will set you up with your database.

Get in a console able to use `psql` and Run

```
psql -U postgres
```

*TIP*: If you are running PSQL with docker,
to access bash inside the container do

```
docker exec -ti <psql-container> /bin/bash
```

Once inside you create your database with

```
CREATE DATABASE bentobox;
\q
```

Now you need to restore the schema. Do from the comamnd line

```
psql -U postgres bentobox < database.sql
```

You are good to go.

### Retrieving program dependencies

```
make boostrap
```

### Compiling the program

```
make crawler
```

Will put your program in `./build/bin/crawler`. This is compiled for debian
though. To run it...

### Running the program

```
make run-crawler
```

Which will run the program inside the dev container.

#### Adding arguments to the executable inside the container

To add arguments, do, for example

```
ARGS="--help" make run-crawler
```

### Herman, I want to compile and run the program without docker

Fine. Do,

```
glide install # You need to have glide (https://github.com/Masterminds/glide)
go build -o build/bin/crawler main.go
./build/bin/crawler
```
