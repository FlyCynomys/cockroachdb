# *配置cockroach 数据库*
***
>## Step 1. 启动第1个节点
        $ cockroach start --insecure --background

   如下输出：

        CockroachDB node starting at 2017-04-20 16:35:27.89629114 -0400 EDT
        build:      beta-20170420 @ 2017/04/20 20:27:23 (go1.8.1)
        admin:      http://localhost:8080
        sql:        postgresql://root@localhost:26257?sslmode=disable
        logs:       cockroach-data/logs
        store[0]:   path=cockroach-data
        status:     initialized new cluster
        clusterID:  {dab8130a-d20b-4753-85ba-14d8956a294c}
        nodeID:     1
   
* 节点之间的通讯通过26257端口交互，是不加密的；8080端口提供给ADMIN来进行访问；
* 节点的数据存储在命令中的 cockroach-data 目录中
* 还有一些辅助信息，像节点ID和SQL的访问URL etc.
* 默认情况下会只使用25%的内存，要改变这个限制的话可以在启动时使用 --cache 参数来更改

***
>## Step 2. 添加节点到集群中

### 在单节点情况下，上述已经足够使用可以开始构建自己的数据库表了，在实际生产中却是至少要3个以上的节点来保证数据的高可用和可恢复性

        1. 自动备份
        2. 自主重新平衡
        3. 失效容忍或者说失效恢复

>### 启动第2个节点
        cockroach start --insecure \
        --background \
        --store=node2 \
        --port=26258 \
        --http-port=8081 \
        --join=localhost:26257

>### 启动第3个节点
        cockroach start --insecure \
        --background \
        --store=node3 \
        --port=26259 \
        --http-port=8082 \
        --join=localhost:26258

        在新添加的节点命令中使用了 --join 的命令,这个命令指定了要加入的初始节点的地址（26257端口）

***

## Step 3. 测试集群

### 启动内建的SQL client以执行SQL语句，无论从集群的那一个节点启动运行，
### 都会以这个节点作为一个GATEWAY网关来对集群进行执行
        $ cockroach sql --insecure
        # Welcome to the cockroach SQL interface.
        # All statements must be terminated by a semicolon.
        # To exit: CTRL + D.
####  剩下就是执行SQL语句了
        在节点1上执行

> *     CREATE DATABASE bank;

> *     CREATE TABLE bank.accounts (id INT PRIMARY KEY, balance DECIMAL);

> *     INSERT INTO bank.accounts VALUES (1, 1000.50);

> *     SELECT * FROM bank.accounts;

        在节点2上执行

> *     SELECT * FROM bank.accounts;
#### 建议在集群部署启动时使用同样的端口和存储结构，这样的话对运维和执行命令的时候负担也轻一些
***

## Step 4. 集群监视器

> *     打开集群的admin页面 [admin](http://localhost:8080)


#### 可以参考

*  自动备份[replication](https://www.cockroachlabs.com/docs/demo-data-replication.html)
*  自动平衡[rebalancing](https://www.cockroachlabs.com/docs/demo-automatic-rebalancing.html)
*  失效处理或者容忍[fault tolerance](https://www.cockroachlabs.com/docs/demo-fault-tolerance-and-recovery.html)


***


## Step 5. 停止集群

        $ cockroach quit --insecure
        停止第一个节点之后，节点2和3会进入失效处理或者容忍状态
        $ cockroach quit --insecure --port=26258
        停止第2个节点，第2个节点停止后，备份机制会失效（只剩最后一个节点了）
        这样停止最后一个节点，就只能使用kill process 这种暴力的手段了

***
