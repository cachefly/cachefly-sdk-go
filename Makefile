# Makefile for running Service usage examples

EXAMPLES_DIR_SERVICES := examples/services
EXAMPLES_DIR_SERVICE_DOMAINS := examples/service_domains

# Default target: run all service examples
default: service-all

# Run the service-list example
service-list:
	@echo "Running example: service-list"
	@go run $(EXAMPLES_DIR_SERVICES)/list/main.go

# Run the get service by ID example, prompting if SERVICE_ID isn’t set
service-getbyid:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-getbyid with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/getbyid/main.go $$SERVICE_ID; \
	'

# Run the update service by ID example, prompting if SERVICE_ID isn’t set
service-updatebyid:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-updatebyid with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/updatebyid/main.go $$SERVICE_ID; \
	'

# Run the activate service example, prompting if SERVICE_ID isn’t set
service-activate:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-activate with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/activate/main.go $$SERVICE_ID; \
	'

# Run the activate service example, prompting if SERVICE_ID isn’t set
service-deactivate:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-deactivate with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/deactivate/main.go $$SERVICE_ID; \
	'

# Run the enable access logging example, prompting if SERVICE_ID isn’t set
service-enableaccesslogging:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-enableaccesslogging with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/enableaccesslogging/main.go $$SERVICE_ID; \
	'

# Run the enable origin logging example, prompting if SERVICE_ID isn’t set
service-enableoriginlogging:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-enableoriginlogging with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/enableoriginlogging/main.go $$SERVICE_ID; \
	'

# Run the disable origin logging example, prompting if SERVICE_ID isn’t set
service-disableoriginlogging:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-disableoriginlogging with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICES)/disableoriginlogging/main.go $$SERVICE_ID; \
	'

# Run all service examples
service-all: service-list service-getbyid service-updatebyid service-activate service-deactivate
	@echo "Running all service examples"

# SERVICE_DOMAINS

# Run the service domains list example
service-domain-list:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-domain-list with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_DOMAINS)/list/main.go $$SERVICE_ID; \
	'

# Run the delete service domain by ID example
service-domain-delete:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  DOMAIN_ID="${DOMAIN_ID:-}"; \
	  if [ -z "$$DOMAIN_ID" ]; then \
	    read -p "Enter domain ID: " DOMAIN_ID; \
	  fi; \
	  echo "Running example: service-domain-delete with service ID $$SERVICE_ID and domain ID $$DOMAIN_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_DOMAINS)/delete/main.go $$SERVICE_ID $$DOMAIN_ID; \
	'

# Run the update service domain by ID example
service-domain-update:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  DOMAIN_ID="${DOMAIN_ID:-}"; \
	  if [ -z "$$DOMAIN_ID" ]; then \
	    read -p "Enter domain ID: " DOMAIN_ID; \
	  fi; \
	  echo "Running example: service-domain-update with service ID $$SERVICE_ID and domain ID $$DOMAIN_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_DOMAINS)/update/main.go $$SERVICE_ID $$DOMAIN_ID; \
	'

# SERVICE_OPTIONS

EXAMPLES_DIR_SERVICE_OPTIONS := examples/service_options

# Run the basic service options example
service-options-get-basic:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-get-basic with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/get_basic/main.go $$SERVICE_ID; \
	'

# Run the save basic service options example
service-options-save:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-save with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/save/main.go $$SERVICE_ID; \
	'

# Run the get legacy API key example
service-options-get-legacy-apikey:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-get-legacy-apikey with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/get_legacy_apikey/main.go $$SERVICE_ID; \
	'

# Run the regenerate legacy API key example
service-options-regenerate-legacy-apikey:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-regenerate-legacy-apikey with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/regenerate_legacy_apikey/main.go $$SERVICE_ID; \
	'

# Run the delete legacy API key example
service-options-delete-legacy-apikey:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-delete-legacy-apikey with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/delete_legacy_apikey/main.go $$SERVICE_ID; \
	'

# Run the get ProtectServe key example
get-protectserve-key:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-get-protectserve-key with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/get_protectserve_key/main.go $$SERVICE_ID; \
	'

# Run the regenerate ProtectServe key example
regenerate-protectserve-key:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-regenerate-protectserve-key with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/regenerate_protectserve_key/main.go $$SERVICE_ID; \
	'

# Run the update ProtectServe key options example
update-protectserve-key-options:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-update-protectserve-key-options with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/update_protectserve_key_options/main.go $$SERVICE_ID; \
	'

# Run the delete ProtectServe key example
delete-protectserve-key:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-delete-protectserve-key with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/delete_protectserve_key/main.go $$SERVICE_ID; \
	'

# Run the get FTP settings example
get-ftp-settings:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-get-ftp-settings with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/get_ftp_settings/main.go $$SERVICE_ID; \
	'

# Run the regenerate FTP password example
regenerate-ftp-password:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-options-regenerate-ftp-password with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_SERVICE_OPTIONS)/regenerate_ftp_password/main.go $$SERVICE_ID; \
	'


# SERVICE_OPTIONS_REFERER_RULES

EXAMPLES_DIR_REFERER_RULES := examples/referer_rules

# List referer rules
referer-rules-list:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: referer-rules-list with service ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_REFERER_RULES)/list/main.go $$SERVICE_ID; \
	'

# Get referer rule by ID
referer-rule-get:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  RULE_ID="${RULE_ID:-}"; \
	  if [ -z "$$RULE_ID" ]; then \
	    read -p "Enter rule ID: " RULE_ID; \
	  fi; \
	  echo "Running example: referer-rule-get with service ID $$SERVICE_ID and rule ID $$RULE_ID"; \
	  go run $(EXAMPLES_DIR_REFERER_RULES)/getbyid/main.go $$SERVICE_ID $$RULE_ID; \
	'

# Create a referer rule
referer-rule-create:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: referer-rule-create with service ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR_REFERER_RULES)/create/main.go $$SERVICE_ID; \
	'

# Update a referer rule by ID
referer-rule-update:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  RULE_ID="${RULE_ID:-}"; \
	  if [ -z "$$RULE_ID" ]; then \
	    read -p "Enter rule ID: " RULE_ID; \
	  fi; \
	  echo "Running example: referer-rule-update with service ID $$SERVICE_ID and rule ID $$RULE_ID"; \
	  go run $(EXAMPLES_DIR_REFERER_RULES)/update/main.go $$SERVICE_ID $$RULE_ID; \
	'

# Delete a referer rule by ID
referer-rule-delete:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  RULE_ID="${RULE_ID:-}"; \
	  if [ -z "$$RULE_ID" ]; then \
	    read -p "Enter rule ID: " RULE_ID; \
	  fi; \
	  echo "Running example: referer-rule-delete with service ID $$SERVICE_ID and rule ID $$RULE_ID"; \
	  go run $(EXAMPLES_DIR_REFERER_RULES)/delete/main.go $$SERVICE_ID $$RULE_ID; \
	'
