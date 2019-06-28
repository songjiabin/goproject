package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `
<!--index.wxml-->
<view class="container">
  <view class="text_h1b3">
    <swiper indicator-dots="{{indicatorDots}}"  interval="5000" duration="500">
      <block wx:for="{{banners}}">
        <swiper-item>
          <image src="{{item.ImageUrl}}" class="slide-image" bindtap='bannerClick' id='{{index}}' />
        </swiper-item>
      </block>
    </swiper>
  </view>
	
  <div>1</div> 
  <div>2</div>
  <div>3</div>
  <div>4
  </div>
  <view class='ScanCode'>
    <view bindtap='btnScan'>
      <text>扫一扫\n<text class='ScanCodeText'>扫码进入会议</text></text>
    </view>
    <view class='ScanCodeB' bindtap='btnShowQRCode'>
      <text>签到码\n<text class='ScanCodeText'>快速活动签到</text></text>
    </view>
  </view>
<!--活动列表-->
<scroll-view scroll-y="true" style='height:100%' bindscrolltolower="scrollLower">
  <view wx:for="{{batchList}}" class="HomePageList">
	<image src="{{item.Logo}}"  class="HomePLImg"/>
	<view class="HomePLWords" bindtap='btnActivityInfo'>
		<text class="ListWordP" style="font-weight: bold;">{{(item.BeginDateTime)}}</text>
		<text class="ListWordP" style="font-size: 15px;">{{item.Name}}</text>
		<text class="ListWordP">{{item.ProvinceName+item.CityName+item.CountyName+item.Address}}</text>

	</view>
  <view class="HomePLPre" wx:if="{{item.IsSignUp==='true'}}">已报名</view>
	<view class="HomePLPre"  id="{{item.ID}}"  bindtap='signUpActivity' wx:if="{{item.IsSignUp==='false'}}">立即报名</view>
</view>
</scroll-view>
</view>`

	//想要里面的<div>标签 但是不会匹配换行的
	//result := regexp.MustCompile(`<div>(.*)</div>`)

	//result := regexp.MustCompile(`<div>(?s:(.*))</div>`)

	result := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	if result == nil {
		//匹配错误
		fmt.Println("error")
		return
	}

	r := result.FindAllStringSubmatch(str, -1)
	fmt.Println(r)

	//过滤
	//[[<div>1</div> 1] [<div>2</div> 2] [<div>3</div> 3] [<div>4
	//</div> 4
	//]]

	for _, v := range r {
		fmt.Println(v[0])
		fmt.Println(v[1])
	}

}
