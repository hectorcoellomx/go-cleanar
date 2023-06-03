package user

type PagoService interface {
	ProcesarPago(monto float64) error
}

type pagoService struct{}

func NewPagoService() PagoService {
	return &pagoService{}
}

func (s *pagoService) ProcesarPago(monto float64) error {
	// Lógica para procesar el pago...
	return nil
}
