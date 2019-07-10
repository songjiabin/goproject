

-- 登录数据库
mysql - u root - p

--显示所有的数据库
show databases

-- 切换到哪个数据库
use dbname

-- 显示某个数据库下的所有表
show tables

--- 退出
quit




/*基本增删改查*/


-- 创建数据库  名字为 mytest
create database mytest charset=utf8

-- 删除数据库
drop database mytest


-- 创建表
CREATE  TABLE  person(
  id INT PRIMARY KEY auto_increment, --唯一的主键 自增长 auto_increment
  name VARCHAR(20) ,--字符串 上线是20个字符
  age INT ,
  sex bool,
  rmb FLOAT ,
  word text -- 文本类型
);


-- 查看表的数据格式
DESC  person


-- 插入数据
INSERT INTO person(name,age,rmb) VALUES(
  "张三丰",90,123450000
);

INSERT INTO person(name,age,rmb,word) VALUES (
  "郭小四",20,112323,"我是文坛巨人哦"
);

-- 批量插入数据
INSERT INTO person(name,age,rmb) VALUES
("张无忌",34,1132323),
("金毛狮王",54,5555323),
("大辉夜",122,1104042323);


--查看表里面的数据
SELECT * FROM  person;

-- 查询里面的个别字段
select name,age  from person ;

-- 删除表person中名字为 "郭小四" 的数据
DELETE FROM person WHERE  name="郭小四";


-- 修改
UPDATE person SET  name ="张大侠" WHERE  name="张三丰";



-- 删除一列
-- 删除T_city中的CitySrot列
ALTER  TABLE T_city DROP CitySort;