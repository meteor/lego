package resolver

import (
	"bytes"
	"fmt"
	"sort"
)

// ObtainError is returned when there are specific errors available per domain.
type ObtainError map[string]error

func (e ObtainError) Error() string {
	buffer := bytes.NewBufferString("acme: Error -> One or more domains had a problem:\n")

	var domains []string
	for domain := range e {
		domains = append(domains, domain)
	}
	sort.Strings(domains)

	for _, domain := range domains {
		buffer.WriteString(fmt.Sprintf("[%s] %s\n", domain, e[domain]))
	}
	return buffer.String()
}
