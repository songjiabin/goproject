
-- join 联合查询


-- 中国各省的id,省份的名字
SELECT ProId,ProName FROM T_province;

-- 河北所有的地级市 id,名字
SELECT CityID,CityName FROM T_city WHERE ProID =
(
 SELECT ProID FROM T_Province WHERE ProName="河北省"
);


-- 联合查询 将两个表关联到一起   前提是两个表查询处理的结果是一样的列数
-- 使用上面的表头
SELECT ProId,ProName FROM T_province
UNION
-- 河北所有的地级市 id,名字
SELECT CityID,CityName FROM T_city WHERE ProID =
(
 SELECT ProID FROM T_Province WHERE ProName="河北省"
);




-- 查询中国河北有多少地级市
SELECT COUNT(CityID) FROM T_city WHERE ProId=
(
 SELECT ProId FROM T_province WHERE  ProName="河北省"
);



-- 联合查询
-- 前提是必须有相同的光联字段
select ProName,CityName  from (
 T_Province JOIN T_City ON T_Province.ProId =T_City.ProId
) WHERE ProName="河北省";


-- 统计各省地级市的数量，输出省名，地级市数量
SELECT  T_City.ProID,  ProName, COUNT(CityId) cc FROM
(
  T_Province JOIN T_City ON T_Province.ProId =T_City.ProId
)
GROUP BY ProId
ORDER BY cc DESC;



-- 每个省份中最大的城市id及其名称
-- 用子查询做
SELECT  ProID,max(CityId)  FROM  T_City  GROUP BY ProId;

-- -用联合查询
select  CityName,max(CityId)  from
(
  T_Province JOIN T_City ON T_Province.ProId =T_City.ProId
) GROUP  BY T_City.ProId;





--- 查询地级市最多的省份取前10名
-- 使用子查询
SELECT cityName, COUNT(ProId) FROM T_City GROUP by ProId ORDER BY COUNT(ProId) DESC limit 10 ;


--  使用联合查询
SELECT ProName,COUNT(T_City.ProId),CityName FROM
(T_Province JOIN T_City ON T_Province.ProId =T_City.ProId)
GROUP by T_City.ProId
ORDER BY COUNT(T_City.ProId) DESC limit 15 ;






-- 查询拥有20个以上区县的城市
SELECT T_City.CityName, COUNT(T_District.Id) AS disCount FROM
(
  T_District JOIN T_City ON T_District.CityID =T_City.CityID
)
GROUP BY T_District.CityId
HAVING  disCount>20
ORDER BY disCount DESC ;




--- 区县最多的城市(直辖市不要,特别行政区不要)是哪个省的什么市，查询结果包含省、市名、区县数量
SELECT * FROM
(SELECT ProId,  CityName,COUNT(Id) as ci FROM
  (
    T_District JOIN T_City ON T_District.CityID =T_City.CityID
  )
GROUP BY T_District.CityId --按照cityId进行分组
ORDER BY  ci DESC
limit 5
) tcp
JOIN T_Province tp ON tcp.ProID=tp.ProID
WHERE tp.ProRemark !="直辖市" ;


-- 内联系
select ProName,CityName from
(
    T_Province tp join T_City tc on tp.ProID=tc.ProID
);


-- 插入数据 到省表  独有的部分
insert into  T_Province(ProName) values ("日本省");



-- 插入市独有的数据
insert into T_City (CityName) values ("洛杉矶市");





-- 左联 将省表中特有的数据 也进行显示了
select ProName,CityName from
(
    T_Province tp left join T_City tc on tp.ProID=tc.ProID
);



-- 右联 将市中特有的数据 也进行了显示
select ProName,CityName from
(
 T_Province tp right join T_City tc on tp.ProID=tc.ProID
)





-- 2019 宜居城市： ("青岛","昆明","大连","威海","苏州","珠海","厦门","重庆","深圳")
-- 求各省分别拥有的宜居城市，各有几个  降序，省名、宜居城市数量、
-- 因为要数量，所以要进行分组

select ProName, count(tc.ProID) ct from
(
select ProID ,CityName  from T_City where CityName in ("青岛市","昆明市","大连市","威海市","苏州市","珠海市","厦门市","重庆市","深圳市")
)tc join T_Province tp on tc.ProID=tp.ProID
group by tc.ProID
order by ct desc ;








-- 你家乡所在的省个拥有哪些宜居城市？










