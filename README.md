# Dungeon Time API

## Using just
Install just from:
https://github.com/casey/just

Then run:
```
just build
just run
just test
```

## Database migrations
Uses https://github.com/golang-migrate/migrate/ 

``` bash
just init
```

``` bash
just migrate <up, down, drop>
```