# 钉钉消息接入文档

## 1、设计思路

借鉴MQ消息，分为Topic、Listener。

### table

​	ding_talk:钉钉token配置

​	msg_topic:消息主题

​	msg_topic_ding_talk:钉钉监听的消息主题

​	alarm_msg:告警消息

​	db_connection:数据库连接

​	monitor_sql:业务SQL监控

### SQL

```mysql
CREATE TABLE `ding_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `access_token` varchar(255)  DEFAULT NULL COMMENT '授权code',
  PRIMARY KEY (`id`) USING BTREE
) COMMENT='钉钉授权配置';

CREATE TABLE `msg_topic` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sn` varchar(50)  DEFAULT NULL COMMENT '消息主题',
  `remark` varchar(255)  DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
) COMMENT='消息主题配置';

CREATE TABLE `msg_topic_ding_talk` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topic_id` int(11) DEFAULT NULL COMMENT '消息主题id',
  `ding_talk_id` int(11) DEFAULT NULL COMMENT '钉钉ID',
  PRIMARY KEY (`id`) USING BTREE
) COMMENT='消息主题和钉钉配置';

CREATE TABLE `db_connection` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `alias_name` varchar(255)  NOT NULL COMMENT '别名',
  `driver_name` varchar(255)  DEFAULT NULL COMMENT '连接名mysql\\pgsql',
  `data_source` varchar(255)  DEFAULT NULL COMMENT '数据源配置',
  PRIMARY KEY (`id`) USING BTREE
) COMMENT='数据源连接配置';


CREATE TABLE `monitor_sql` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cron` varchar(255)  DEFAULT NULL COMMENT '表达式',
  `sql` varchar(255)  DEFAULT NULL COMMENT 'SQL语句查询列名为 cnt 单条数据',
  `alias_name` varchar(20)  DEFAULT NULL COMMENT '连接别名',
  `status` int(11) DEFAULT NULL COMMENT '监控状态0-关,1-开',
  `note` varchar(255)  DEFAULT NULL COMMENT '说明',
  `db_id` int(11) DEFAULT NULL COMMENT '数据源编号',
  `show_sql` int(11) DEFAULT '0' COMMENT '是否在msg里面拼接SQL，1-拼接',
  `sn` varchar(255)  DEFAULT NULL COMMENT '消息主题sn',
  PRIMARY KEY (`id`) USING BTREE
) COMMENT='业务SQL监控';


CREATE TABLE `alarm_msg` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `msg` text  COMMENT '消息内容',
  `status` int(11) DEFAULT '0' COMMENT '发送状态0-待发送,1-已发送',
  `access_token` varchar(255)  DEFAULT NULL COMMENT 'token',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`) USING BTREE
) COMMENT='告警消息';
```

### 数据库信息

```
RDS-dev 测试库：alarm
RDS-dev 生产库：alarm_prod
```



## 2、接入文档

### 2.1 消息主题和钉钉配置

2.1.1 保存消息主题

```mysql
INSERT INTO `alarm_prod`.`msg_topic`(`sn`, `remark`) VALUES ('5d608a828f45461bb7c8302494f106d7', '业务SQL监控告警');
```

2.1.2 保存钉钉机器人access_token

```mysql
INSERT INTO `alarm_prod`.`ding_talk`(`access_token`) VALUES ('3d958826cf5996563f9c73d34a5789a250ba2805569669424d7802922a66ff74');
```

2.1.3 消息主题和钉钉机器人关联

```mysql
INSERT INTO `alarm_prod`.`msg_topic_ding_talk`(`topic_id`, `ding_talk_id`) VALUES (1, 1);
```



### 2.2 业务SQL监控接入

2.2.1 插入业务监控SQL

```mysql
INSERT INTO `alarm_prod`.`monitor_sql`(`cron`, `sql`, `alias_name`, `status`, `note`, `db_id`, `show_sql`, `sn`) VALUES ( '0/30 * * * ?', 'select count(*) cnt from wb_account_detail where create_time > CURRENT_DATE - 1 AND create_time < (CURRENT_TIMESTAMP - INTERVAL 10 MINUTE) AND txid is NULL;', 'wallet', 1, ' [业务监控] 账户变动明细10分钟以前的数据没有区块ID的数量：', 2, 1, '5d608a828f45461bb7c8302494f106d7');
```

2.2.2 根据SQL编号(id)启动监控

```
post请求：http://18.162.116.178:8900/alarm-api/monitor/start/task

body: {"id":8}
```

2.2.3 根据SQL编号(id)关闭监控

```
post请求：http://18.162.116.178:8900/alarm-api/monitor/stopc/task
body: {"id":8}
```



### 2.3 服务器告警消息接入

2.3.1 应用消息告警接入

```
post请求：http://18.162.116.178:8900/alarm-api/alarm/msg/save
body:{
    	"msg":"我是通过API给钉钉发送消息的~~~",
    	"sn":"5d608a828f45461bb7c8302494f106d7",
    	"requestNo":"test1"
	}
```

