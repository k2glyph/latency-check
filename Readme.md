# Latency Check
Example:<br>
*config.yaml*
```
mysql
  hostname: 10.13.20.1
  username: root
  password: root
  database: db1
  runsqlfile: ./run.sql
  drainsqlfile: ./drain.sql
```
`./latency --config ./config.yaml mysql check`