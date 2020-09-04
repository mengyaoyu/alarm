# 水桥-钉钉消息接入文档

## 1、设计思路

借鉴MQ消息，分为Topic、Listener。

### table

​	alarm_notice_listener:接收者配置(钉钉token)

​	alarm_notice_topic:消息主题

​	alarm_notice_topic_listener:接收者监听的消息主题

​	alarm_notice_msg:告警消息

​	alarm_db_connection:数据库连接

​	alarm_notice_monitor_sql:业务SQL监控

### SQL

```mysql
-- ----------------------------
-- Table structure for alarm_db_connection
-- ----------------------------
DROP TABLE IF EXISTS `alarm_db_connection`;
CREATE TABLE `alarm_db_connection`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `alias_name` varchar(255) CHARACTER  NOT NULL COMMENT '别名',
  `driver_name` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '连接名mysql\\pgsql',
  `data_source` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '数据源配置',
  PRIMARY KEY (`id`) USING BTREE
)COMMENT = '数据源连接配置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for alarm_notice_monitor_sql
-- ----------------------------
DROP TABLE IF EXISTS `alarm_notice_monitor_sql`;
CREATE TABLE `alarm_notice_monitor_sql`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cron` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '表达式',
  `sql` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT 'SQL语句查询列名为 cnt 单条数据',
  `alias_name` varchar(20) CHARACTER  NULL DEFAULT NULL COMMENT '连接别名',
  `status` int(11) NULL DEFAULT NULL COMMENT '监控状态0-关,1-开',
  `note` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '说明',
  `db_id` int(11) NULL DEFAULT NULL COMMENT '数据源编号',
  `show_sql` int(11) NULL DEFAULT 0 COMMENT '是否在msg里面拼接SQL，1-拼接',
  `sn` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '消息主题sn',
  PRIMARY KEY (`id`) USING BTREE
) COMMENT = '业务SQL监控' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for alarm_notice_listener
-- ----------------------------
DROP TABLE IF EXISTS `alarm_notice_listener`;
CREATE TABLE `alarm_notice_listener`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `access_token` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '授权code',
  `notice_type` varchar(20) CHARACTER  NULL DEFAULT NULL COMMENT '类型：DT、WX、SMS、PHONE',
  PRIMARY KEY (`id`) USING BTREE
)  COMMENT = '授权配置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for alarm_notice_msg
-- ----------------------------
DROP TABLE IF EXISTS `alarm_notice_msg`;
CREATE TABLE `alarm_notice_msg`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `msg` text CHARACTER  NULL COMMENT '消息内容',
  `status` int(11) NULL DEFAULT 0 COMMENT '发送状态0-待发送,1-已发送',
  `access_token` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT 'token',
  `create_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `notice_type` varchar(10) CHARACTER  NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) COMMENT = '告警消息' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for alarm_notice_topic
-- ----------------------------
DROP TABLE IF EXISTS `alarm_notice_topic`;
CREATE TABLE `alarm_notice_topic`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sn` varchar(50) CHARACTER  NULL DEFAULT NULL COMMENT '消息主题',
  `remark` varchar(255) CHARACTER  NULL DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`) USING BTREE
)   COMMENT = '消息主体配置' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Table structure for alarm_notice_topic_listener
-- ----------------------------
DROP TABLE IF EXISTS `alarm_notice_topic_listener`;
CREATE TABLE `alarm_notice_topic_listener`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topic_id` int(11) NULL DEFAULT NULL COMMENT '消息主题id',
  `listener_id` int(11) NULL DEFAULT NULL COMMENT '监听者ID',
  PRIMARY KEY (`id`) USING BTREE
)  COMMENT = '消息主题和监听者' ROW_FORMAT = DYNAMIC;
```

### 数据库信息

```
RDS-dev 测试库：alarm
RDS-dev 生产库：alarm_prod
```



## 2、接入文档

### 2.1 消息主题和接收者配置

**2.1.1 保存消息主题**

```mysql
INSERT INTO `alarm`.`alarm_notice_topic`(`sn`, `remark`) VALUES ('5d608a828f45461bb7c8302494f106d7', '业务SQL监控告警');
```

**2.1.2 保存接收者配置**

```mysql
INSERT INTO `alarm`.`alarm_notice_listener`(`access_token`, `notice_type`) VALUES ('xxx', 'DT');
```

**2.1.3 消息主题和接收人关联**

```mysql
INSERT INTO `alarm`.`alarm_notice_topic_listener`(`topic_id`, `listener_id`) VALUES (1, 1);
```

### 2.2 业务SQL监控接入

**2.2.1 业务数据库配置SQL**

```mysql
INSERT INTO `alarm_prod`.`alarm_db_connection`(`id`, `alias_name`, `driver_name`, `data_source`) VALUES (1, 'flow', 'mysql', 'x:x*@tcp(x.com:3306)/Flow?charset=utf8');
```

**2.2.1 插入业务监控SQL**

```mysql
INSERT INTO `alarm`.`alarm_monitor_sql`(`cron`, `sql`, `alias_name`, `status`, `note`, `db_id`, `show_sql`, `sn`) VALUES ( '0/30 * * * ?', 'select count(*) cnt from wb_account_detail where create_time > CURRENT_DATE - 1 AND create_time < (CURRENT_TIMESTAMP - INTERVAL 10 MINUTE) AND txid is NULL;', 'wallet', 1, ' [业务监控] 账户变动明细10分钟以前的数据没有区块ID的数量：', 2, 1, '5d608a828f45461bb7c8302494f106d7');
```

**2.2.2 根据SQL编号(id)启动监控**

```
post请求：http://localhost:8900/alarm-api/monitor/start/task

body: {"id":8}
```

**2.2.3 根据SQL编号(id)关闭监控**

```
post请求：http://localhost:8900/alarm-api/monitor/stopc/task
body: {"id":8}
```



### 2.3 服务器告警消息接入

**2.3.1 应用消息告警接入**

```
post请求：http://localhost:8900/alarm-api/alarm/msg/save
body:{
    	"msg":"我是通过API给钉钉发送消息的~~~",
    	"sn":"5d608a828f45461bb7c8302494f106d7",
    	"requestNo":"test1"
	}
```

