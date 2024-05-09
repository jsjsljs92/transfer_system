package common

import "strconv"

func ConvertToFloat32(value string) (float32, error) {
	// Convert string to float64
	floatVal, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0, err
	}

	// Format float to have at most 5 decimal points
	formattedVal := float32(floatVal)
	return formattedVal, nil
}
