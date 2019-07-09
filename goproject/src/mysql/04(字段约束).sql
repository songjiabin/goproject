
-- 不重复的内容 DISTINCT
select DISTINCT ProID from   T_city;




-- 字段约束

CREATE TABLE   person (
  id INT  ,
  NAME VARCHAR(20),
  age int ,
  rmb FLOAT ,
  sex bool,
  word text,

  -- 日期类型
  birthday DATE ,
  -- 生辰八字
  scbz datetime
);


-- 插入数据
insert into person (  NAME,  sex,  birthday, scbz) values ("毛爷爷", TRUE ,18931226,18931226123456);


-- 添加字段的约束
CREATE TABLE   person (
  id INT  PRIMARY KEY auto_increment,
  NAME VARCHAR(20) NOT NULL UNIQUE , -- UNIQUE 唯一的约束
  age int DEFAULT 0,
  rmb FLOAT DEFAULT 0,
  sex bool ,
  word VARCHAR(100) DEFAULT "这个家伙很懒，啥也没留下", --注意text不能有默认值的

  -- 日期类型
  birthday DATE ,
  -- 生辰八字
  scbz datetime
);

-- 插入
insert into person (  NAME,  sex,  birthday, scbz) values ("周总理", FALSE ,18931226,18931226123456);