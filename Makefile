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