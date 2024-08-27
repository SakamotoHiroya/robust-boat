.PHONY: redocly
redocly:
	redocly lint ./docs/openapi.yaml
	redocly bundle ./docs/openapi.yaml -o ./docs/bundle.yaml

.PHONY: ogen
ogen:
	ogen -target ui/api -clean docs/bundle.yaml