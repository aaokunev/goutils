package numbers

func GetInt64Value(value *int64, defaultValue int64) int64 {
	if value == nil {
		return defaultValue
	}

	return *value
}
