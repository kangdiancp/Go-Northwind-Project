package payment

type Category struct {
	CategoryID   int16  `db:"category_id" json:"categoryId"`
	CategoryName string `db:"category_name" json:"categoryName"`
	Description  string `db:"description" json:"description"`
	Picture      []byte `db:"picture" json:"picture"`
}
