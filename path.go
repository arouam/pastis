package pastis


type params map[string]string

func (p *params) get(key string) string{
	return (*p)[key]
}

func (p *params) add(key, value string) {
	(*p)[key] = value
}

type queries map[string]string

func (q *queries) get(key string) string{
	return (*q)[key]
}

func (q *queries) add(key, value string) {
	(*q)[key] = value
}
