package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	engine := gin.New()

	//post请求 下级门店信息
	postStartOrderShopNext(engine)

	//post请求服务单的信息
	postStatOrderShop(engine)

	//post请求技师
	postStatOrderMassagist(engine);

	//修改预估值比例
	postModifyRevenue(engine)

	//得到服务单列表数据
	postQueryOrder(engine)

	//技师详情列表数据
	postMassagistEvaluate(engine)

	engine.Run()
}

//技师详情界面
func postMassagistEvaluate(engine *gin.Engine) {

	type Result struct {
		Page_count  int         `json:"page_count"`
		Total       int         `json:"total"`
		Total_count int         `json:"total_count"`
		Data        interface{} `json:"data"`
	}

	type Data struct {
		Id             int         `json:"id"`
		Order_no       string      `json:"order_no"`
		Customer_code  string      `json:"customer_code"`
		Massagist_code string      `json:"massagist_code"`
		Star           int         `json:"star"`
		Advise         string      `json:"advise"`
		Anonymous      int         `json:"anonymous"`
		Pic            interface{} `json:"pic"`
		Create_time    string      `json:"create_time"`
	}

	var str = [3]string{"img1", "img2", "img3"}

	d := Data{
		Id:             1,
		Order_no:       "订单号",
		Customer_code:  "客户编码",
		Massagist_code: "按摩师编码",
		Star:           1,
		Advise:         "这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议这个是我的建议",
		Anonymous:      1,
		Pic:            str,
		Create_time:    "评估日期",
	}

	dataList := make([]Data, 0)
	dataList = append(dataList, d)
	dataList = append(dataList, d)

	r := Result{
		Page_count:  4,
		Total:       2,
		Total_count: 7,
		Data:        dataList,
	}

	engine.POST("/massagist_evaluate", func(context *gin.Context) {
		time.Sleep(time.Second * 2)
		context.JSON(http.StatusOK, gin.H{
			"code":   0,
			"msg":    "ok",
			"result": r,
		})
	})
}

func postQueryOrder(engine *gin.Engine) {

	type Result struct {
		Page_count  int         `json:"page_count"`
		Total       int         `json:"total"`
		Total_count int         `json:"total_count"`
		Data        interface{} `json:"data"`
	}

	type Data struct {
		Id                 int         `json:"id"`
		Order_no           string      `json:"order_no"`
		Comment            string      `json:"comment"`
		State              int         `json:"state"`
		Type               int         `json:"type"`
		Price              int         `json:"price"`
		State_reason       string      `json:"state_reason"`
		Device_code        string      `json:"device_code"`
		Anonymous          int         `json:"anonymous"`
		Volume_code        string      `json:"volume_code"`
		Create_name        string      `json:"create_name"`
		Create_id          int         `json:"create_id"`
		Create_time        string      `json:"create_time"`
		Update_name        string      `json:"update_name"`
		Update_id          int         `json:"update_id"`
		Update_time        string      `json:"update_time"`
		Shop               interface{} `json:"shop"`
		Service            interface{} `json:"service"`
		Customer           interface{} `json:"customer"`
		Massagist          interface{} `json:"massagist"`
		Shop_evaluate      interface{} `json:"shop_evaluate"`
		Massagist_evaluate interface{} `json:"massagist_evaluate"`
		Service_evaluate   interface{} `json:"service_evaluate"`
		Questionnaire      interface{} `json:"questionnaire"`
		Ai                 interface{} `json:"ai"`
	}

	type Shop struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Code     string `json:"code"`
		Province string `json:"province"`
		City     string `json:"city"`
		District string `json:"district"`
		Town     string `json:"town"`
		Address  string `json:"address"`
	}

	type Service struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Code string `json:"code"`
	}

	type Customer struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		Code       string `json:"code"`
		Nick       string `json:"nick"`
		Sex        int    `json:"sex"`
		Age        int    `json:"age"`
		Unionid    string `json:"unionid"`
		Openid     string `json:"openid"`
		Addr       string `json:"addr"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		Head_image string `json:"head_image"`
	}

	type Massagist struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		Code       string `json:"code"`
		Nick       string `json:"nick"`
		Sex        int    `json:"sex"`
		Age        int    `json:"age"`
		Unionid    string `json:"unionid"`
		Openid     string `json:"openid"`
		Addr       string `json:"addr"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		Head_image string `json:"head_image"`
	}

	type Shop_evaluate struct {
		Star      int         `json:"star"`
		Advise    string      `json:"advise"`
		Anonymous int         `json:"anonymous"`
		Pic       interface{} `json:"pic"`
	}

	type Massagist_evaluate struct {
		Star      int         `json:"star"`
		Advise    string      `json:"advise"`
		Anonymous int         `json:"anonymous"`
		Pic       interface{} `json:"pic"`
	}

	type Service_evaluate struct {
		Star   int         `json:"star"`
		Advise string      `json:"advise"`
		Pic    interface{} `json:"pic"`
	}

	type Questionnaire struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		Order       int         `json:"order"`
		Question    interface{} `json:"question"`
		Ask         interface{} `json:"ask"`
		Create_name string      `json:"create_name"`
		Create_id   int         `json:"create_id"`
		Create_time string      `json:"create_time"`
	}

	type Question struct {
		Id            int    `json:"id"`
		OptionContent string `json:"option_content"`
		Sn            int    `json:"sn"`
		IsLast        string `json:"is_last"`
	}

	type Ai struct {
		Id         int         `json:"id"`
		SymContent string      `json:"sym_content"`
		ImpPro     interface{} `json:"imp_pro"`
	}

	type ImpPro struct {
		Id            int    `json:"id"`
		Groups        int    `json:"groups"`
		MaxTime       int    `json:"max_time"`
		MinTime       int    `json:"min_time"`
		RightTime     int    `json:"right_time"`
		Content       string `json:"content"`
		Sn            int    `json:"sn"`
		BodyName      string `json:"body_name"`
		BodyChildName string `json:"body_child_name"`
		TreatName     string `json:"treat_name"`
		StrengthName  string `json:"strength_name"`
		ImpProContent string `json:"imp_pro_content"`
	}

	engine.POST("/query_order", func(context *gin.Context) {

		shop := Shop{
			Id:       1,
			Name:     "商家名称",
			Code:     "商家编码",
			Province: "省",
			City:     "市",
			District: "区县",
			Town:     "乡镇",
			Address:  "地址",
		}

		service := Service{
			Id:   1,
			Name: "30分钟：除肩颈外，任选一个部位;",
			Code: "PD00000001",
		}

		customer := Customer{
			Id:         1,
			Name:       "名字",
			Code:       "用户编码",
			Nick:       "昵称",
			Sex:        1,
			Age:        1,
			Unionid:    "微信唯一id",
			Openid:     "公众号用户id",
			Addr:       "地址",
			Phone:      "手机号码",
			Email:      "邮箱",
			Head_image: "https://cdn.pixabay.com/photo/2019/11/22/13/55/background-4644934__340.jpg",
		}

		massagist := Massagist{
			Id:         1,
			Name:       "名字",
			Code:       "用户编码",
			Nick:       "昵称",
			Sex:        1,
			Age:        1,
			Unionid:    "微信唯一id",
			Openid:     "公众号用户id",
			Addr:       "地址",
			Phone:      "手机号码",
			Email:      "邮箱",
			Head_image: "https://cdn.pixabay.com/photo/2019/11/26/21/26/people-4655553__340.jpg",
		}

		picList := make([]string, 0)
		picList = append(picList, "1.jpg")
		picList = append(picList, "2.jpg")
		picList = append(picList, "3.jpg")
		shop_evaluate := Shop_evaluate{
			Star:      1,
			Advise:    "建议",
			Anonymous: 1,
			Pic:       picList,
		}

		massagist_evaluate := Massagist_evaluate{
			Star:      1,
			Advise:    "建议",
			Anonymous: 1,
			Pic:       picList,
		}

		service_evaluate := Service_evaluate{
			Star:   1,
			Advise: "建议",
			Pic:    picList,
		}

		question1 := Question{
			Id:            1,
			OptionContent: "头部",
			Sn:            1,
			IsLast:        "N",
		}

		question2 := Question{
			Id:            2,
			OptionContent: "手肘",
			Sn:            2,
			IsLast:        "N",
		}

		question3 := Question{
			Id:            3,
			OptionContent: "手腕",
			Sn:            3,
			IsLast:        "N",
		}
		questionList := make([]Question, 0)
		questionList = append(questionList, question1)
		questionList = append(questionList, question2)
		questionList = append(questionList, question3)

		questionnaire := Questionnaire{
			Id:          1,
			Name:        "服务部位",
			Order:       1,
			Question:    questionList,
			Ask:         picList,
			Create_name: "创建者名字",
			Create_id:   1,
			Create_time: "创建日期",
		}

		impPro := ImpPro{
			Id:            1,
			Groups:        3,
			MaxTime:       600,
			MinTime:       120,
			RightTime:     330,
			Content:       "确认手的力度为轻柔",
			Sn:            1,
			BodyName:      "颈部",
			BodyChildName: "肩胛提肌",
			TreatName:     "肌筋膜手法",
			StrengthName:  "轻柔",
			ImpProContent: "治疗方法：肌筋膜手法,手法力度：轻柔,治疗部位：颈部,肩胛提肌,治疗时间：5分/组，3组，治疗时间：2分-10分/组，3组",
		}

		ai1 := Ai{
			Id:         18,
			SymContent: "过肩运动，羽毛球，网球，排球",
			ImpPro:     impPro,
		}

		ai2 := Ai{
			Id:         18,
			SymContent: "颈部动作检查:动作正常无疼痛FN",
			ImpPro:     impPro,
		}

		aiList := make([]Ai, 0)
		aiList = append(aiList, ai1)
		aiList = append(aiList, ai2)

		data := Data{
			Id:                 1,
			Order_no:           "ORDER20191101120001123",
			Comment:            "是否匿名评价， 0 = 未评价  1 = 匿名评价 2 = 实名评价",
			State:              1,
			Type:               1,
			Price:              11001,
			State_reason:       "状态原因",
			Device_code:        "DEVICE1234578",
			Anonymous:          1,
			Volume_code:        "消费次卷",
			Create_name:        "创建者名字",
			Create_id:          1,
			Create_time:        "2019-11-20 00：00：00",
			Update_name:        "修改者名字",
			Update_id:          1,
			Update_time:        "修改时间",
			Shop:               shop,
			Service:            service,
			Customer:           customer,
			Massagist:          massagist,
			Shop_evaluate:      shop_evaluate,
			Massagist_evaluate: massagist_evaluate,
			Service_evaluate:   service_evaluate,
			Questionnaire:      questionnaire,
			Ai:                 aiList,
		}
		dataList := make([]Data, 0)
		for i := 0; i < 10; i++ {
			dataList = append(dataList, data)
		}
		dataList = append(dataList, data)

		r := Result{
			Page_count:  4,
			Total:       2,
			Total_count: 7,
			Data:        dataList,
		}

		time.Sleep(2 * time.Second)

		context.JSON(http.StatusOK, gin.H{
			"code":   0,
			"msg":    "ok",
			"result": r,
		})

	})
}

func postModifyRevenue(engine *gin.Engine) {
	engine.POST("/modify_revenue", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"code":   "0",
			"msge":   "成功",
			"result": nil,
		})
	})
}

func postStatOrderMassagist(engine *gin.Engine) {

	type Result struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		Code        string      `json:"code"`
		Nick        string      `json:"nick"`
		Phone       string      `json:"phone"`
		Head_image  string      `json:"head_image"`
		Evaluate    int         `json:"evaluate"`
		Total       int         `json:"total"`
		OrderCount  int         `json:"order_count"`
		Certificate interface{} `json:"certificate"`
	}

	type Certificate struct {
		Id                     int    `json:"id"`
		Certificate_id         string `json:"certificate_id"`
		Certificate_name       string `json:"certificate_name"`
		Certificate_level_id   string `json:"certificate_level_id"`
		Certificate_level_name string `json:"certificate_level_name"`
		Template               string `json:"template"`
		Visible                int    `json:"visible"`
	}

	c := Certificate{
		Id:                     1,
		Certificate_id:         "证书类型ID",
		Certificate_name:       "证书类型名称",
		Certificate_level_id:   "证书等级ID",
		Certificate_level_name: "证书等级名称",
		Template:               "证书模板文件URL",
		Visible:                0,
	}

	type ResponseBean struct {
		Shop_id int    `json:"shop_id" form:"shop_id"`
		Start   string `json:"start" form:"start"`
		End     string `json:"end" form:"end"`
		Types   string `json:"type" form:"type"`
		Sort    string `json:"sort" fomr:"sort"`
	}

	cerList := make([]Certificate, 0)
	cerList = append(cerList, c)
	cerList = append(cerList, c)

	result := Result{
		Id:          1,
		Name:        "名字",
		Code:        "用户编码",
		Nick:        "昵称",
		Phone:       "手机号码",
		Head_image:  "https://jjmima.top/images/rockets-4528194__340.png",
		Evaluate:    0,
		Total:       6000,
		OrderCount:  100,
		Certificate: cerList,
	}

	resultList := make([]Result, 0)

	for i := 0; i < 10; i++ {
		resultList = append(resultList, result);
	}
	engine.POST("/stat_order_massagist", func(c *gin.Context) {

		var response ResponseBean;
		c.ShouldBind(&response)

		fmt.Println("start--->", response.Start)
		fmt.Println("end--->", response.End)
		fmt.Println(response);
		time.Sleep(time.Second * 2)
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"msg":    "ok",
			"result": resultList,
		})
	});
}

func postStatOrderShop(engine *gin.Engine) {
	type Result struct {
		Shop         interface{} `json:"shop"`
		Little_order interface{} `json:"little_order"`
		Shop_order   interface{} `json:"shop_order"`
	}

	type Shop struct {
		Total       int `json:"total"`
		Order_count int `json:"order_count"`
	}

	type Little_order struct {
		Total       int `json:"total"`
		Order_count int `json:"order_count"`
	}

	type Shop_order struct {
		Total       int `json:"total"`
		Order_count int `json:"order_count"`
	}

	engine.POST("/stat_order_shop", func(c *gin.Context) {
		s := Shop{
			Total:       6000,
			Order_count: 100,
		}
		l := Little_order{
			Total:       6000,
			Order_count: 100,
		}
		o := Shop_order{
			Total:       6000,
			Order_count: 100,
		}

		r := Result{
			Shop:         s,
			Little_order: l,
			Shop_order:   o,
		}
		time.Sleep(time.Second * 2)
		c.JSON(http.StatusOK, gin.H{
			"code":   0,
			"msg":    "ok",
			"result": r,
		})
	})

}

func postStartOrderShopNext(engine *gin.Engine) {

	type Result struct {
		Total       int         `json:"total"`
		Order_count int         `json:"order_count"`
		Data        interface{} `json:"data"`
	}

	type Data struct {
		Id          int    `json:"id"`
		Name        string `json:"name"`
		Code        string `json:"code"`
		Address     string `json:"address"`
		Total       int    `json:"total"`
		Order_count int    `json:"order_count"`
	}

	type GetData struct {
		Page      string    `form:"page" json:"page"`
		StartTime string    `form:"startTime" json:"startTime"`
		EndtTime  time.Time `form:"endTime json:"endTime""`
	}

	currentPage := 1;
	//门店业务统计（下级门店）
	engine.POST("/stat_order_shop_next", func(c *gin.Context) {

		page := c.DefaultPostForm("page", "defaultPage")
		fmt.Println("page--->", page)

		currentPage += 1
		var getData GetData
		if c.ShouldBind(&getData) == nil {

			page := getData.Page
			fmt.Println("pagebind-->" + page)
		} else {
			fmt.Println("没成功...")
		}

		infos := make([]Data, 0)
		d := Data{
			Id:          1,
			Name:        "商家名称",
			Code:        "商家编码",
			Address:     "地址",
			Total:       6000,
			Order_count: currentPage,
		}

		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		infos = append(infos, d)
		r := Result{
			Total:       6000,
			Order_count: 100,
			Data:        infos,
		}

		time.Sleep(time.Second * 2)

		c.JSON(http.StatusOK, gin.H{
			"msg":    "成功",
			"code":   0,
			"result": r,
		})
	})
}
