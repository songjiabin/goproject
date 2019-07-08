


-- and  or like
-- 找出不叫 区县市 的区县

-- LIKE 以...区

select * from T_District WHERE (DisName not LIKE "%区") AND (DisName not LIKE "%县") AND (DisName not LIKE "%市");





-- 将新疆的名字改成正规的格式

UPDATE T_Province SET ProName="新疆维吾尔族自治区" WHERE ProName LIKE "新%疆%";


-- 找到新疆和内蒙古的id


select * from T_Province  WHERE ProName LIKE "新疆%" OR ProName LIKE "内蒙%";

-- 找出内蒙和新疆的城市
select * from T_City WHERE (ProID=5) OR (ProID=24);




-- order by , limit

-- 查询河北所有的省  按照城市id降序
-- DESC 降序
-- ASC 升序
select * FROM  T_City WHERE ProId=3 ORDER BY  CityId DESC ;



-- limit 截取其中的  并得到其中的前 5名

select * from T_City WHERE ProId=3    ORDER BY  CityId DESC limit 5 ;



-- in   在某个范围以内

-- 2019 宜居城市： ("青岛","昆明","大连","威海","苏州","珠海","厦门","重庆","深圳")

-- 查询广东有几个这个城市
select * from T_City WHERE CityName in ("青岛","昆明","大连","威海","苏州","珠海","厦门","重庆","深圳") and  ProID =19;

-- 上面的效率没有上面的快
select * from T_City
WHERE
ProID = (select  ProId FROM T_Province WHERE ProName="山东省")
AND
CityName IN ("青岛市","昆明市","大连市","威海市","苏州市","珠海市","厦门市","重庆市","深圳市");





-- between   之间  查询省ID 在 10 到20之间
select * from T_Province WHERE ProId BETWEEN 10 AND 20;





-- 分组 group by, having
-- 查询中国各省分别有多少个地级市

select ProId,count(CityID) from T_City
GROUP BY ProID;  -- 在city表中以ProId进行分组。


-- 查询中国各省分别有多少个地级市 降序取前5名
select ProId,count(CityID) as cities from T_City
GROUP BY ProID
ORDER BY COUNT(CityID) DESC
limit 5 ;



-- 查询中国各省分别有多少个地级市 取地级市超过10的省份
select ProId,count(CityID) as cities from T_City
GROUP BY ProID
HAVING cities > 10
ORDER BY COUNT(CityID) DESC



-- 查询中国区县最多的10个城市
select CityId,  COUNT(ID) as cities from T_District
GROUP BY CityID
ORDER BY cities DESC
limit 10;

