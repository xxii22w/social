## 数据库迁移
```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest tags可以替换为其他数据库
migrate create -seq -ext sql -dir ./cmd/migrate/migrations create_users
-ext sql：这个选项指定了迁移脚本文件的扩展名是 .sql

postgres 登陆
psql -U admin -d social
\l      # 列出数据库
\dt     # 列出当前数据库中的所以的表
\d+ table_name # 查看表结构
\c # 查看数据库版本
\q # 退出
\c database_name # 切换数据库
\i file_path # 执行sql文件
\s # 查看当前命令行历史
http://localhost:8080/v1/swagger/index.html
```
