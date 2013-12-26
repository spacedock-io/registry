package context

var context = make(map[interface{}]interface{})

func Set(key interface{}, value interface{}) {
  context[key] = value
}

func Get(key interface{}) interface{}{
  return context[key]
}
