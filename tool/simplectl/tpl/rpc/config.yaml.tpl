rpc-service:
  service-addr: 127.0.0.1
  port: 9001

rpc:
  register-type: etcdv3
  register-addr:
    - 127.0.0.1:23791
    - 127.0.0.1:23792
    - 127.0.0.1:23793
  base-path: local

redis:
  addr: 127.0.0.1:6379
  password:
  db: 0

db-list:
  - disabled: true # 是否启用
    type: mysql # 数据库的类型,目前支持mysql、pgsql
    alias-name: upms # 数据库的名称,注意: alias-name 需要在db-list中唯一
    path: 127.0.0.1
    port: 3306
    config: charset=utf8mb4&parseTime=True&loc=Local
    db-name: upms
    username: root
    password: root
    max-idle-conns: 10
    max-open-conns: 100
    log-mode: error
    log-zap: true