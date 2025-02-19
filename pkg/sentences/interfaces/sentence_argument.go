package interfaces

type SentenceArgument interface {
	// GetValueIsItSetting возвращает значение аргумента мат.выражения и true,
	// если значение аргумента посчитано. Если значение аргумента еще не готово,
	// то возвращается 0 и false
	GetValueIsItSetting() (value int64, isValueSetting bool)
	// SubscribeOnValueWhenItCalculated позволяет нам подписаться на значение, когда оно будет вычислено.
	// Если значение уже вычислено, то вернется nil канал, и флаг isSubscribed = false
	SubscribeOnValueWhenItCalculated() (result chan int64, isSubscribed bool)
}

type SentenceArgumentVariable interface {
	SentenceArgument
	// SetValue устанавливает значение для аргумента мат.выражения
	SetValue(value int64) error
	Name() string
}
