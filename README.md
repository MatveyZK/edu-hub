# edu-hub 
## commands

Build project:

```bash
make
```

Start server (flag is optional):

```bash
./server -config-path path
```

Start container:

```bash
sudo docker compose up -d
```

See logs:

```bash
sudo docker logs --tail 100 goose_migrations
```

SQL console:

```bash
sudo docker exec -it edu-hub psql -U user -d db
```

Stop container:

```bash
sudo docker compose down
```

See all SQL tables:

```
\dt
```

