package key

import (
	"github.com/erathorus/quickstore"
)

const (
	KindProject = "prj"
)

func ValidateProjectKey(k quickstore.Key) bool {
	return k.Parent == "" && k.Kind == KindProject && k.Identifier != "" && isBase64URL(k.Identifier)
}

func GenerateProjectKey() quickstore.Key {
	return quickstore.Key{
		Kind:       KindProject,
		Identifier: quickstore.RandIdentifier(),
	}
}

func isBase64URL(s string) bool {
	for c := range s {
		if ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '-' || c == '_' || c == '=' {
			continue
		}
		return false
	}
	return true
}
