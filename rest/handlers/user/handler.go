package user


type Handler struct {
	svc UserService
}

func NewHandler(svc UserService) *Handler {
	return &Handler{
		svc: svc,
	}
}