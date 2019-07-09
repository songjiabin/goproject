
-- 插入学生数据
insert into student(name,gender) values ("张三丰",1);
insert into student(name,gender) values ("野间忠一郎",1);
insert into student(name,gender) values ("二郎神",1);
insert into student(name,gender) values ("郭小四",0);
insert into student(name,gender) values ("隔壁老王",1);
insert into student(name,gender) values ("练过的六爷",1);
insert into student(name,gender) values ("洪七公",1);
insert into student(name,gender) values ("香香八婆",0);
insert into student(name,gender) values ("马英九",1);
insert into student(name,gender) values ("十三姨",0);
insert into student(name,gender) values ("山本五十六",1);
insert into student(name,gender) values ("包租婆",0);




-- 插入老师的数据
insert into teacher(name,gender) values ("bill",1);
insert into teacher(name,gender) values ("steve",1);
insert into teacher(name,gender) values ("jackma",1);
insert into teacher(name,gender) values ("robin",1);




-- 插入班级
insert into clazz(name) values ("丐帮"), ("小刀会"), ("斧头帮"), ("天地会");


-- 插入课程
insert into course(name) values ("Python"),("Java"),("HTML5"),("PHP");







-- 给所有的学生进行分班  也就是更新学生表中的 班级 id

UPDATE  student set classid =(SELECT id FROM clazz WHERE name="丐帮") WHERE NAME = "野间忠一郎";
update student set classid=(select id from clazz where clazz.name = '小刀会') where student.name='二郎神';
update student set classid=(select id from clazz where clazz.name = '斧头帮') where student.name='张三丰';
update student set classid=(select id from clazz where clazz.name = '天地会') where student.name='郭小四';
update student set classid=(select id from clazz where clazz.name = '丐帮') where student.name='隔壁老王';
update student set classid=(select id from clazz where clazz.name = '小刀会') where student.name='练过的六爷';
update student set classid=(select id from clazz where clazz.name = '斧头帮') where student.name='洪七公';
update student set classid=(select id from clazz where clazz.name = '天地会') where student.name='香香八婆';
update student set classid=(select id from clazz where clazz.name = '丐帮') where student.name='马英九';
update student set classid=(select id from clazz where clazz.name = '小刀会') where student.name='十三姨';
update student set classid=(select id from clazz where clazz.name = '斧头帮') where student.name='山本五十六';
update student set classid=(select id from clazz where clazz.name = '天地会') where student.name='包租婆';








