package styles

// Fix import path

// Theme contains global style configuration
type Theme struct {
	Dark bool
}

var ActiveTheme = &Theme{
	Dark: true,
}
