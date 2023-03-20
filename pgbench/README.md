# pgbench

```bash
# Create a DB
docker-compose exec db psql -U postgres -c "CREATE DATABASE benchmarking;"
```

```bash
docker-compose exec bench

# Init benchmarking data
# -s database size is N * 16M
pgbench -h db -p 5432 -U postgres benchmarking -i -s 1

# Run benchmark
time pgbench -h db -p 5432 -U postgres -c 5 -j 2 -t 1000 benchmarking
```

