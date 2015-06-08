README
======

The nbagame/db package handles syncing data via the stats.nba.com APIs to a local
MySQL database. I recommend syncing data locally if you plan on doing any significant
work with the data.

Setting up
----------

First, you must create a new MySQL database to populate with NBA data.
We assume you name your database 'nbagame':

```
mysql -e "CREATE DATABASE nbagame"
```

Either allow all local users access to this database, or edit the `dbconf.yml` file
with the parameters to connect to your new database.

Then, migrate the database by running

```
goose up
```


