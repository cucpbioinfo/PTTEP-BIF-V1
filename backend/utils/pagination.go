package utils

func GetLimitAndOffset(pageNumber int, pageSize int) (int, int) {
	limit := pageSize
	offset := (pageNumber - 1) * pageSize

	return limit, offset
}

func GetValidPagination(pageNumber int, pageSize int) (int, int) {

	defaultPageNumber := 1
	defaultPageSize := 20
	maximumPageSize := 100

	isValidPageNumber := pageNumber > 0
	isValidPageSize := pageSize > 0 && pageSize <= maximumPageSize

	if isValidPageNumber {
		if isValidPageSize {
			return pageNumber, pageSize
		} else {
			return pageNumber, defaultPageSize
		}
	} else {
		if isValidPageSize {
			return defaultPageNumber, pageSize
		} else {
			return defaultPageNumber, defaultPageSize
		}
	}
}
