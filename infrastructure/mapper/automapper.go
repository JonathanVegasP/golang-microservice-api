package mapper

import "reflect"

func autoMapper(actual interface{}, dest interface{}, channel chan bool) {
	va := reflect.TypeOf(actual)

	values := reflect.ValueOf(actual)

	values = reflect.Indirect(values)

	vd := reflect.TypeOf(dest)

	valuesD := reflect.ValueOf(dest)

	valuesD = reflect.Indirect(valuesD)

	if va.Kind() == reflect.Ptr {
		va = va.Elem()
	}

	if vd.Kind() == reflect.Ptr {
		vd = vd.Elem()
	}

	for i := 0; i < va.NumField(); i++ {
		fa := va.Field(i)
		fak := fa.Type.Kind()
		fan := fa.Name
		value := values.Field(i)
		for j := 0; j < vd.NumField(); j++ {
			fd := vd.Field(j)
			fdk := fd.Type.Kind()
			fdn := fd.Name
			if fak == fdk && fan == fdn {
				field := valuesD.FieldByName(fdn)
				if field.CanSet() {
					field.Set(value)
				}
			}
		}
	}

	channel <- true
}

func AutoMapper(actual interface{}, dest interface{}) {
	channel := make(chan bool)

	go autoMapper(actual, dest, channel)

	<-channel

	close(channel)
}
