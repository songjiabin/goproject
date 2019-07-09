
-- 多对多
-- 一个学生可以有多个学科，一个学科可以对应多个学生  这个就是多对多



/*
  创建课程中间表
 */

create table student_course (
  -- 学生id
  sid int not null,
  -- 课程的id
  cid int not NULL,

  -- 联合主键
  -- 这种组合只能出现一次
  PRIMARY KEY (sid,cid)

);

-- 先进行插入数据
-- 十三姨选择golang
insert into student_course (sid, cid) values((SELECT id FROM student WHERE name="十三姨"),(SELECT id FROM course WHERE name="golang"));
-- 十三姨选择Java
insert into student_course (sid, cid) values((SELECT id FROM student WHERE name="十三姨"),(SELECT id FROM course WHERE name="Java"));
-- 十三姨选择Python
insert into student_course (sid, cid) values((SELECT id FROM student WHERE name="十三姨"),(SELECT id FROM course WHERE name="Python"));



insert into student_course (sid, cid) values((SELECT id FROM student WHERE name="山本五十六"),(SELECT id FROM course WHERE name="golang"));

insert into student_course (sid, cid) values((SELECT id FROM student WHERE name="香香八婆"),(SELECT id FROM course WHERE name="golang"));




-- 能根据学生查询其所有的选课
-- 查看十三姨都选择了什么课程

SELECT *FROM course WHERE id in
(
SELECT cid FROM  student_course WHERE sid =
(
  select id from student WHERE NAME ="十三姨"
)
);




-- 能根据选课查询所有的学生
-- 查看golang 都哪些学生选择过


SELECT * FROM student WHERE id IN
(
  SELECT sid FROM student_course WHERE cid =
  (
    SELECT id FROM course WHERE NAME ="golang"
  )
);
