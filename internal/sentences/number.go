package sentences

type Number int64

func (n Number) GetValueIsItSetting() (int64, bool) {
	return int64(n), true
}

func (n Number) SubscribeOnValueWhenItCalculated() (result chan int64, isSubscribed bool) {
	return nil, false
}
