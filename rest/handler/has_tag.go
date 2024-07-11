package handler

import "reflect"

func hasTag(t reflect.Type, tagName string) bool {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Traverse the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Check if the field has the specified tag
		if _, ok := field.Tag.Lookup(tagName); ok {
			return true
		}

		// If the field is an embedded struct, traverse its fields
		if field.Anonymous {
			fieldType := field.Type
			if fieldType.Kind() == reflect.Ptr {
				fieldType = fieldType.Elem()
			}

			if fieldType.Kind() == reflect.Struct {
				if hasTag(fieldType, tagName) {
					return true
				}
			}
		}
	}

	return false
}
