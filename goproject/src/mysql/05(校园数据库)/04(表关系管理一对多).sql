

 -- 一对多关系的管理

 -- 一个班级可以有多个学生  但是一个学生只有有一个班级

 -- 所有学生中要有班级的 id

-- 插入数据
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



-- 根据班级找出所有的学生
-- 找出小刀会的所有学生
SELECT * FROM student WHERE classid =(
  SELECT id FROM clazz WHERE NAME ="小刀会"
);



-- 根据学生找出所有的班级
-- 查询十三姨是哪个班的

SELECT * FROM clazz WHERE id=(SELECT classid FROM student WHERE name="十三姨" );
