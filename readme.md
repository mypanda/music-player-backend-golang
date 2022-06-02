### 未完成，搁置

```shell
# 启动镜像
docker run \
--name mongodb_music_player \
-v /root/gopath/src/github.com/mypanda/music-player-backend/databases/mongodb:/data/db \
--rm \
-d \
-p 27017:27017 \
mongo
# -e MONGODB_INITDB_ROOT_USERNAME=example-user 
# -e MONGODB_INITDB_ROOT_PASSWORD_FILE=/run/secrets/mongo-root-pw | example-user 

# 创建管理员账号失败，可通过设置环境变量密码登录
# 进入容器
# docker exec -it mongodb_music_player mongo admin
# 创建管理员 
# use admin
# db.createUser({user:"admin",pwd:"123456",roles:[{"role":"userAdminAnyDatabase","db":"admin"}]})
# db.auth('admin', '123456')
# 测试管理员登录
# docker exec -it mongodb_music_player mongo admin -u amdin -p 123456 --authenticationDatabase admin

# 创建数据库和所属数据库账号
docker exec -it mongodb_music_player mongo admin
use music_player_db
db.ta.insert({}) #这里插入一条空数据是为了有了collection存在，数据库才会存在
# 在数据库创建文件
db.createUser({ user:'music_admin',pwd:'123456',roles:[ { role:'readWrite', db: 'music_player_db'}, "readWrite"]})
# db.createUser({ user:'admin',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});
# 连接
db.auth('music_admin', '123456')
# 使用创建的用户登录
docker exec -it mongodb_music_player mongo music_player_db -u music_admin -p 123456 --authenticationDatabase music_player_db

# 数据持久化卷
# 创建卷
docker volume create mongo_local_data
# 运行景相融弄个器
docker run --name my_mongo -v /mongo_local_data:/data/db --rm -d -p 27018:27017 mongo

https://blog.csdn.net/u013302168/article/details/121111750
```

```
docker build -t docker-syn-image .
docker run -d -p 12345:12345 -e “MONGODB_URI=YOUR_URI_HERE” docker-syn-image
```

### 
```
configs
controllers
models
db
middlewares
public
routers
test
service
```

### 参考
https://github.dev/uiters/gin-example
https://dev.to/hackmamba/build-a-rest-api-with-golang-and-mongodb-gin-gonic-version-269m
https://blog.csdn.net/weixin_43881017/article/details/113701802
https://github.com/liuhongdi/digv30