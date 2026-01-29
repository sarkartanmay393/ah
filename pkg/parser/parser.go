package parser

import (
	"bufio"
	"os"
	"strings"
)

type AliasDef struct {
	Name    string
	Command string
	Source  string
}

// ParseAliases strictly extracts "alias key=value" lines.
// It ignores everything else (comments, functions, logic).
func ParseAliases(filePath string) ([]AliasDef, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var aliases []AliasDef
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// 1. Filter: Must start with "alias "
		if !strings.HasPrefix(line, "alias ") {
			continue
		}

		// 2. Split by first "="
		// alias foo='bar baz'
		// left: "alias foo", right: "'bar baz'"
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		// 3. Extract Name
		// "alias foo" -> "foo"
		namePart := strings.TrimSpace(parts[0])
		name := strings.TrimPrefix(namePart, "alias ")
		name = strings.TrimSpace(name) // "foo"

		// 4. Extract Value
		value := strings.TrimSpace(parts[1])

		// 5. Unquote (Basic)
		// We want the raw command to re-quote it safely later.
		if len(value) >= 2 {
			first := value[0]
			last := value[len(value)-1]
			if (first == '"' && last == '"') || (first == '\'' && last == '\'') {
				value = value[1 : len(value)-1]
			}
		}

		if name != "" && value != "" {
			aliases = append(aliases, AliasDef{
				Name:    name,
				Command: value,
				Source:  filePath,
			})
		}
	}

	return aliases, scanner.Err()
}
