package handler

type Handler interface {

}

type HiHandler struct {
	Handler
}
func (this *HiHandler) SayHi() string  {
	return "hi~"
}
