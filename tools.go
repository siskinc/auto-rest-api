package autorestapi

func checkPageSizeAndPageIndex(pageSize, pageIndex int) bool {
	if pageIndex > 0 && pageSize > 0 {
		return true
	}
	return false
}
