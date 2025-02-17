package gameplay

type Handler struct {
	Executor *ExecutionObject
}

func (h *Handler) HandleGameplay() error {

	return nil
}

func New(executor ExecutionObject) *Handler {
	return &Handler{
		Executor: &executor,
	}
}
