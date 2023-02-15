package main

import (
	"fmt"
	"strings"
)

func main() {
	in := "| || $k; $k; $b; $b; $sdsad; $ $$ $$;$;;;;;$"
	encoded := EncodeOptimized(in)
	out := DecodeOptimized(encoded)
	res := in == out
	fmt.Println(in)
	fmt.Println(encoded)
	fmt.Println(out)
	fmt.Println(res)
}

func Encode(value string) string {
	if value == "" {
		return value
	}
	return strings.Replace(strings.Replace(value, "$", "$k;", -1), "|", "$b;", -1)
}

func EncodeOptimized(value string) string {
	if value == "" {
		return value
	}

	var b []byte
	var last int
	for i := 0; i < len(value); i++ {
		if value[i] == '$' {
			b = append(b, value[last:i]...)
			b = append(b, '$', 'k', ';')
			last = i + 1
		} else if value[i] == '|' {
			b = append(b, value[last:i]...)
			b = append(b, '$', 'b', ';')
			last = i + 1
		}
	}
	b = append(b, value[last:]...)

	return string(b)
}

func Decode(input string) string {
	if input == "" {
		return ""
	}

	var sb strings.Builder
	index := 0

	for index < len(input) {
		if input[index] == '$' {
			semicolonIndex := strings.Index(input[index+1:], ";") + index + 1

			if semicolonIndex != -1 {
				entity := input[index+1 : semicolonIndex]

				switch entity {
				case "k":
					sb.WriteRune('$')
					index = semicolonIndex + 1
					continue

				case "b":
					sb.WriteRune('|')
					index = semicolonIndex + 1
					continue
				}
			}
		}

		sb.WriteRune(rune(input[index]))
		index++
	}

	return sb.String()
}

func DecodeOptimized(input string) string {
	if input == "" {
		return ""
	}

	var b []byte
	last := 0

	for i := 0; i < len(input); i++ {
		if input[i] == '$' {
			b = append(b, input[last:i]...)
			if i+2 < len(input) && input[i+1] == 'k' && input[i+2] == ';' {
				b = append(b, '$')
				i += 2
			} else if i+2 < len(input) && input[i+1] == 'b' && input[i+2] == ';' {
				b = append(b, '|')
				i += 2
			} else {
				b = append(b, '$')
			}
			last = i + 1
		}
	}
	b = append(b, input[last:]...)

	return string(b)
}

func DecodeOptimizedSimplified(input string) string {
	if input == "" {
		return ""
	}

	var b []byte
	last := 0

	for i := 0; i < len(input); i++ {
		if input[i] == '$' {
			b = append(b, input[last:i]...)
			if i+2 < len(input) && input[i+2] == ';' {
				switch input[i+1] {
				case 'k':
					b = append(b, '$')
				case 'b':
					b = append(b, '|')
				default:
					b = append(b, '$')
				}
				i += 2
			} else {
				b = append(b, '$')
			}
			last = i + 1
		}
	}
	b = append(b, input[last:]...)

	return string(b)
}
