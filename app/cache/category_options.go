/**
 * Copyright 2014 @ z3q.net.
 * name :
 * author : jarryliu
 * date : 2014-02-05 21:53
 * description :
 * history :
 */
package cache

import (
	"bytes"
	"fmt"
	"github.com/jsix/gof/algorithm/iterator"
	"go2o/core/domain/interface/content"
	"go2o/core/domain/interface/product"
	"go2o/core/infrastructure/domain/util"
	"go2o/core/service/rsi"
)

func readToCategoryDropList(mchId int32) []byte {
	categories := rsi.ProductService.GetCategories(mchId)
	buf := bytes.NewBuffer([]byte{})
	var f iterator.WalkFunc = func(v1 interface{}, level int) {
		c := v1.(*product.Category)
		if c.ID != 0 {
			buf.WriteString(fmt.Sprintf(
				`<option class="opt%d" value="%d">%s</option>`,
				level,
				c.ID,
				c.Name,
			))
		}
	}
	util.WalkSaleCategory(categories, &product.Category{ID: 0}, f, nil)
	return buf.Bytes()
}

// 获取销售分类下拉选项
func GetDropOptionsOfSaleCategory(mchId int32) []byte {
	return readToCategoryDropList(mchId)
}

// 获取商品模型下拉选项
func GetDropOptionsOfProModel() string {
	buf := bytes.NewBuffer([]byte{})
	list := rsi.ProductService.GetModels()
	for _, v := range list {
		buf.WriteString(fmt.Sprintf(
			`<option value="%d">%s</option>`,
			v.Id,
			v.Name,
		))
	}
	return buf.String()
}

func readToArticleCategoryDropList() []byte {
	categories := rsi.ContentService.GetArticleCategories()
	buf := bytes.NewBuffer([]byte{})
	var f iterator.WalkFunc = func(v1 interface{}, level int) {
		c := v1.(*content.ArticleCategory)
		if c.Id != 0 {
			buf.WriteString(fmt.Sprintf(
				`<option class="opt%d" value="%d">%s</option>`,
				level,
				c.Id,
				c.Name,
			))
		}
	}
	util.WalkArticleCategory(categories, &content.ArticleCategory{Id: 0},
		f, nil)
	return buf.Bytes()
}

// 获取文章栏目下拉选项
func GetDropOptionsOfArticleCategory() []byte {
	return readToArticleCategoryDropList()
}
