package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hotinfo/app/response"
	"hotinfo/app/schedule/baidu"
	"hotinfo/app/schedule/bilibili"
	"hotinfo/app/schedule/bilibili_rank"
	"hotinfo/app/schedule/douyin"

	"hotinfo/app/schedule/hyper"

	"hotinfo/app/schedule/juejin"

	"hotinfo/app/schedule/lol"

	"hotinfo/app/schedule/tengxun"
	"hotinfo/app/schedule/toutiao"
	"hotinfo/app/schedule/wangyi"
	"hotinfo/app/schedule/weibo"

	"hotinfo/app/schedule/zhihu"
)

func Select(c *gin.Context) {
	newName := c.Query("new_name")
	fmt.Println("new_name", newName)
	switch newName {
	case "微博热搜榜":
		resp := weibo.Refresh()
		response.Ok(resp, "微博热搜榜刷新成功", c)
		return
	case "知乎热搜榜":
		resp := zhihu.Refresh()
		response.Ok(resp, "知乎热搜榜刷新成功", c)
		return
	case "网易热搜榜":
		resp := wangyi.Refresh()
		response.Ok(resp, "网易热搜榜刷新成功", c)
		return
	case "头条热搜榜":
		resp := toutiao.Refresh()
		response.Ok(resp, "头条热搜榜刷新成功", c)
		return
	case "腾讯热搜榜":
		resp := tengxun.Refresh()
		response.Ok(resp, "腾讯热搜榜刷新成功", c)
		return
	case "英雄联盟公告":
		resp := lol.Refresh()
		response.Ok(resp, "lol公告刷新成功", c)
		return
	case "掘金热搜榜":
		resp := juejin.Refresh()
		response.Ok(resp, "掘金热搜榜刷新成功", c)
		return
	case "澎湃新闻热搜榜":
		resp := hyper.Refresh()
		response.Ok(resp, "澎湃新闻热搜榜刷新成功", c)
		return
	case "B站热搜榜":
		resp := bilibili.Refresh()
		response.Ok(resp, "B站热搜榜刷新成功", c)
		return
	case "抖音热搜榜":
		resp := douyin.Refresh()
		response.Ok(resp, "抖音热搜榜刷新成功", c)
		return
	case "B站排行榜":
		resp := bilibili_rank.Refresh()
		response.Ok(resp, "B站排行榜刷新成功", c)
		return
	case "百度热搜榜":
		resp := baidu.Refresh()
		response.Ok(resp, "百度热搜榜", c)
		return

	case "全部刷新":
		type AllInfo struct {
			BaiDu        []baidu.BaiDu         `json:"bai_du"`
			BilibiliHot  []bilibili.Bilibili   `json:"bilibili_hot"`
			BilibiliRank []bilibili_rank.BRank `json:"bilibili_rank"`
			DouYin       []douyin.DouYin       `json:"dou_yin"`
			Hyper        []hyper.Hyper         `json:"hyper"`
			JueJin       []juejin.Juejin       `json:"jue_jin"`
			LoL          []lol.Lol             `json:"lo_l"`
			TengXun      []tengxun.TengXun     `json:"teng_xun"`
			TouTiao      []toutiao.TouTiao     `json:"tou_tiao"`
			WangYi       []wangyi.WangYi       `json:"wang_yi"`
			WeiBo        []weibo.WeiBo         `json:"wei_bo"`
			ZhiHu        []zhihu.ZhiHu         `json:"zhi_hu"`
		}

		baidu := baidu.Refresh()
		bilibili := bilibili.Refresh()
		bilibiliRank := bilibili_rank.Refresh()
		douyin := douyin.Refresh()
		hyper := hyper.Refresh()
		juejin := juejin.Refresh()
		lol := lol.Refresh()
		tengxun := tengxun.Refresh()
		toutiao := toutiao.Refresh()
		wangyi := wangyi.Refresh()
		weibo := weibo.Refresh()
		zhihu := zhihu.Refresh()

		allInfo := AllInfo{
			BaiDu:        baidu,
			BilibiliHot:  bilibili,
			BilibiliRank: bilibiliRank,
			DouYin:       douyin,
			Hyper:        hyper,
			JueJin:       juejin,
			LoL:          lol,
			TengXun:      tengxun,
			TouTiao:      toutiao,
			WangYi:       wangyi,
			WeiBo:        weibo,
			ZhiHu:        zhihu,
		}
		response.Ok(allInfo, "刷新成功", c)
		return

	}

}
