package utils

import (
	"fmt"
)

// BuildImageURL generates the correct URL for a product image,
// from local server (default) or Cloudinary (custom)
func BuildImageURL(filename, baseHost, scheme string) string {
	if filename == "product_default.webp" {
		return fmt.Sprintf("%s%s/public/%s", scheme, baseHost, filename)
	}
	return fmt.Sprintf(
		"https://res.cloudinary.com/dye6aug6j/image/upload/v1750059966/coffee_shop/product_imgs/%s",
		filename,
	)
}

func BuildImageProfileURL(filename, baseHost, scheme string) string {
	if filename == "" || filename == "avatar_default.webp" {
		return fmt.Sprintf("%s%s/public/%s", scheme, baseHost, filename)
	}
	return fmt.Sprintf(
		"https://res.cloudinary.com/dye6aug6j/image/upload/v1750059966/coffee_shop/profile_imgs/%s",
		filename,
	)
}

func BuildImageProfileURLV2(filename, baseHost, scheme string) string {
	if filename == "" {
		filename = "avatar_default.webp"
	}

	if filename == "avatar_default.webp" {
		return fmt.Sprintf("%s%s/public/%s", scheme, baseHost, filename)
	}

	return fmt.Sprintf(
		"https://res.cloudinary.com/dye6aug6j/image/upload/v1750059966/coffee_shop/profile_imgs/%s",
		filename,
	)
}
