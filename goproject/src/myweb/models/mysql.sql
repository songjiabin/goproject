



create table if not exists beego_test(
  id INT PRIMARY KEY auto_increment, --唯一的主键 自增长 auto_increment
  name VARCHAR(20) ,--字符串 上线是20个字符
)



insert into  beego_test(name) values ("许美琳")



select * from beego_test
where (
 id=1
)