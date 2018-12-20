package server

import "personal/moudle"

func CheckAndAdd(category string)error{
	var cate moudle.Category
	cates, err := cate.List()
	if err != nil {
		return err
	}

	for _, v := range cates {
		if v.Name == category {
			return nil
		}
	}

	cate.Name = category
	return cate.Add()
}