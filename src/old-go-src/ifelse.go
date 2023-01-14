package solution

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		locale = "en" 
	}

	return locale + "." + domain
}
