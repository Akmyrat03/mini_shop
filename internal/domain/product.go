package domain

type Product struct {
	ID            int
	CategoryID    int
	TitleTK       string
	TitleEN       string
	TitleRU       string
	DescriptionTK string
	DescriptionEN string
	DescriptionRU string
	Price         float64
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
}
