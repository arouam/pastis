package pastis

type Object map[string]interface{}

func Error(err error)Object{
	return Object{"err": err.Error()}
}