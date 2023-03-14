package response

// 首页分类数据VO(第三级)
type ThirdLevelCategoryVO struct {
	CategoryId    int    `json:"categoryId"`
	CategoryLevel int    `json:"categoryLevel"`
	CategoryName  string `json:"categoryName" `
	CategoryImg   string `json:"category_img"`
}

type SecondLevelCategoryVO struct {
	CategoryId            int                    `json:"categoryId"`
	ParentId              int                    `json:"parentId"`
	CategoryLevel         int                    `json:"categoryLevel"`
	CategoryName          string                 `json:"categoryName" `
	CategoryImg           string                 `json:"category_img"`
	ThirdLevelCategoryVOS []ThirdLevelCategoryVO `json:"thirdLevelCategoryVOS"`
}

type NewBeeMallIndexCategoryVO struct {
	CategoryId int `json:"categoryId"`
	//ParentId               int                      `json:"parentId"`
	CategoryLevel          int                     `json:"categoryLevel"`
	CategoryName           string                  `json:"categoryName" `
	CategoryImg            string                  `json:"category_img"`
	SecondLevelCategoryVOS []SecondLevelCategoryVO `json:"secondLevelCategoryVOS"`
}
