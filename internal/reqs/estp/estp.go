package estp

type Estp struct {
	PageProps struct{
		Announces []struct{
			Number string
			Title string
		}
	}
}

func New() *Estp {
	return &Estp{}
}