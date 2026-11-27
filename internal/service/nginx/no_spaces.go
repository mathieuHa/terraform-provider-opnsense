package nginx

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

var noSpaces = stringvalidator.RegexMatches(regexp.MustCompile(`^[^" \t]*$`), "must not contain spaces, tabs or double quotes")
