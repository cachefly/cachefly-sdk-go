# Makefile for running Service usage examples

EXAMPLES_DIR := examples/services

# Default target: run all service examples
default: service-all

# Run the service-list example
service-list:
	@echo "Running example: service-list"
	@go run $(EXAMPLES_DIR)/list/main.go

# Run the get service by ID example, prompting if SERVICE_ID isn’t set
service-getbyid:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-getbyid with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/getbyid/main.go $$SERVICE_ID; \
	'

# Run the update service by ID example, prompting if SERVICE_ID isn’t set
service-updatebyid:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-updatebyid with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/updatebyid/main.go $$SERVICE_ID; \
	'

# Run the activate service example, prompting if SERVICE_ID isn’t set
service-activate:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-activate with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/activate/main.go $$SERVICE_ID; \
	'

# Run the activate service example, prompting if SERVICE_ID isn’t set
service-deactivate:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-deactivate with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/deactivate/main.go $$SERVICE_ID; \
	'

# Run the enable access logging example, prompting if SERVICE_ID isn’t set
service-enableaccesslogging:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-enableaccesslogging with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/enableaccesslogging/main.go $$SERVICE_ID; \
	'

# Run the enable origin logging example, prompting if SERVICE_ID isn’t set
service-enableoriginlogging:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-enableoriginlogging with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/enableoriginlogging/main.go $$SERVICE_ID; \
	'

# Run the disable origin logging example, prompting if SERVICE_ID isn’t set
service-disableoriginlogging:
	@bash -c '\
	  SERVICE_ID="${SERVICE_ID:-}"; \
	  if [ -z "$$SERVICE_ID" ]; then \
	    read -p "Enter service ID: " SERVICE_ID; \
	  fi; \
	  echo "Running example: service-disableoriginlogging with ID $$SERVICE_ID"; \
	  go run $(EXAMPLES_DIR)/disableoriginlogging/main.go $$SERVICE_ID; \
	'

# Run all service examples
service-all: service-list service-getbyid service-updatebyid service-activate service-deactivate
	@echo "Running all service examples"
