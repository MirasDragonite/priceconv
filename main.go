package priceconv

// Initialization valute converter to work with it methods
func CreateValuteConverter() Valutes {
	result := Valutes{
		Valutes: map[string]Valute{},
	}
	return result
}
