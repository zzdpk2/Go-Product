package main

type HashSet struct {
	Map map[interface{}]bool
}

type Set interface {
	Put(key string)
	Keys() []string
	Contains(key string) bool
	Remove(key string)
	PutIfAbsent(key string) (old string, absent bool)
}

func (set *HashSet) Put(key string) {
	if _, ok := set.Map[key]; !ok {
		set.Map[key] = true
	}
}

// func (set *HashSet) Keys() []string {
// 	curMap := set.Map
// 	res := make([]string, len(curMap))
// 	if len(curMap) == 0 {
// 		return res
// 	}
//
// 	for v := range curMap {
// 		res = append(res, v.(string))
// 	}
// 	return res
// }

func (set *HashSet) Keys() []string {
	curMap := set.Map
	res := []string{}

	for v := range curMap {
		res = append(res, v.(string))
	}
	return res
}

func (set *HashSet) Contains(key string) bool {
	curMap := set.Map
	if _, ok := curMap[key]; ok {
		return true
	}
	return false
}

func (set *HashSet) Remove(key string) bool {
	if set.Contains(key) {
		delete(set.Map, key)
		return true
	}
	return false
}

func (set *HashSet) PutIfAbsent(key string) (old string, absent bool) {
	if set.Contains(key) {
		return key, false
	} else {
		set.Put(key)
		return "new string added!", true
	}
}
