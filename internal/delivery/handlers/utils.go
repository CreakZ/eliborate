package handlers

import (
	"eliborate/internal/constants"
	"fmt"
)

func GetIdFromKeys(keys map[string]any) (int, error) {
	idRaw, exists := keys[constants.KeyUserID]
	if !exists {
		return 0, fmt.Errorf("no user_id information provided")
	}

	id, ok := idRaw.(int)
	if !ok {
		return 0, fmt.Errorf("jwt contains non-int user_id value")
	}

	return id, nil
}

func GetRoleFromKeys(keys map[string]any) (string, error) {
	roleRaw, exists := keys[constants.KeyRole]
	if !exists {
		return "", fmt.Errorf("no role information provided")
	}

	role, ok := roleRaw.(string)
	if !ok {
		return "", fmt.Errorf("no role information provided")
	}

	return role, nil
}

func GetIdAndRoleFromKeys(keys map[string]any) (int, string, error) {
	id, err := GetIdFromKeys(keys)
	if err != nil {
		return 0, "", err
	}

	role, err := GetRoleFromKeys(keys)
	if err != nil {
		return 0, "", err
	}

	return id, role, nil
}
