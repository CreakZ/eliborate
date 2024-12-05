package utils

import (
	"fmt"
	"reflect"
)

func CollectTagValuesFromStructByTag(tag string, src interface{}) []string {
	srcType := reflect.TypeOf(src)

	fmt.Println(srcType.String())

	if srcType.Kind() != reflect.Struct {
		return []string{}
	}

	var tags []string
	for i := range srcType.NumField() {
		field := srcType.Field(i)

		if structTag, ok := field.Tag.Lookup(tag); ok {
			tags = append(tags, structTag)
		}
	}
	return tags
}

func CollectMapKeys(m map[string]interface{}) []string {
	if len(m) == 0 {
		return []string{}
	}

	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	return keys
}
