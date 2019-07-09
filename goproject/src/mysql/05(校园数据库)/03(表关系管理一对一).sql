


-- 一对一的管理

-- 给班级分配班主任
--
update clazz SET  masterid=(
SELECT id FROM teacher WHERE name="robin"
) WHERE name="丐帮";

update clazz SET  masterid=(
SELECT id FROM teacher WHERE name="jackma"
) WHERE name="小刀会";


update clazz SET  masterid=(
SELECT id FROM teacher WHERE name="bill"
) WHERE name="斧头帮";


update clazz SET  masterid=(
SELECT id FROM teacher WHERE name="steve"
) WHERE name="天地会";





-- 根据班级找班主任
SELECT *FROM teacher
WHERE id =
  (SELECT masterid FROM clazz WHERE name ="天地会");



-- 根据班主任找班级
 SELECT *FROM clazz WHERE masterid =
  (SELECT id FROM teacher
  WHERE name = "bill");