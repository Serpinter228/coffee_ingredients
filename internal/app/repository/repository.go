package repository

type Ingredient struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Unit        string
	Status      string
	ImageKey    string
	VideoKey    string
	Likes       []int
}

var ingredients = []Ingredient{
	{ID: 1, Name: "Зёрна арабики", Description: "Мягкий шоколадный вкус с лёгкими нотами карамели.", Price: 850, Unit: "кг", Status: "published", ImageKey: "arabica_beans.jpg", VideoKey: "arabica_beans.mp4", Likes: []int{1, 2, 5, 8}},
	{ID: 2, Name: "Овсяное молоко", Description: "Нежное и сливочное.", Price: 220, Unit: "л", Status: "published", ImageKey: "oat_milk.jpg", VideoKey: "oat_milk.mp4", Likes: []int{3, 4}},
	{ID: 3, Name: "Карамельный сироп", Description: "Тёплая сладость.", Price: 490, Unit: "л", Status: "published", ImageKey: "caramel_syrup.jpg", VideoKey: "caramel_syrup.mp4", Likes: []int{1, 6, 7}},
	{ID: 4, Name: "Какао", Description: "Насыщенный вкус.", Price: 640, Unit: "кг", Status: "published", ImageKey: "cocoa.jpg", VideoKey: "cocoa.mp4", Likes: []int{2, 5}},
	{ID: 5, Name: "Корица", Description: "Пряный аромат.", Price: 380, Unit: "кг", Status: "published", ImageKey: "cinnamon.jpg", VideoKey: "cinnamon.mp4", Likes: []int{4, 8}},
	{ID: 6, Name: "Ванильный сироп", Description: "Мягкая ваниль.", Price: 520, Unit: "л", Status: "published", ImageKey: "vanilla_syrup.jpg", VideoKey: "vanilla_syrup.mp4", Likes: []int{1, 3}},
	{ID: 7, Name: "Фундучный сироп", Description: "Черновик для тестирования.", Price: 350, Unit: "л", Status: "draft", ImageKey: "hazelnut_syrup.jpg", VideoKey: "hazelnut_syrup.mp4", Likes: []int{}},
	{ID: 8, Name: "Старый бленд", Description: "Удалённый товар.", Price: 600, Unit: "кг", Status: "deleted", ImageKey: "old_blend.jpg", VideoKey: "old_blend.mp4", Likes: []int{1}},
}

type Repository struct{}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) GetByID(id int) *Ingredient {
	for i := range ingredients {
		if ingredients[i].ID == id && ingredients[i].Status != "deleted" {
			return &ingredients[i]
		}
	}
	return nil
}

func (r *Repository) GetNext(id int) *Ingredient {
	published := r.GetPublished()
	for i, ing := range published {
		if ing.ID == id {
			if i+1 < len(published) {
				return &published[i+1]
			}
			return &published[0]
		}
	}
	if len(published) > 0 {
		return &published[0]
	}
	return nil
}

func (r *Repository) GetDraft() *Ingredient {
	for i := range ingredients {
		if ingredients[i].Status == "draft" {
			return &ingredients[i]
		}
	}
	return nil
}

func (r *Repository) GetPublished() []Ingredient {
	var result []Ingredient
	for _, ing := range ingredients {
		if ing.Status == "published" {
			result = append(result, ing)
		}
	}
	return result
}

func (r *Repository) GetPublishedByPrice(maxPrice float64) []Ingredient {
	var result []Ingredient
	for _, ing := range ingredients {
		if ing.Status == "published" && ing.Price <= maxPrice {
			result = append(result, ing)
		}
	}
	return result
}
