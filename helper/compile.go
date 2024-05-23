package helper

import "strings"

func compileMigrationName(migrationNameFile string) string {
	parts := strings.Split(migrationNameFile, "_")

	var functionParts []string

	for i, part := range parts {
		if i < 5 {
			continue
		}
		functionParts = append(functionParts, strings.Title(part))
	}

	return "Create" + strings.Join(functionParts, "")
}
