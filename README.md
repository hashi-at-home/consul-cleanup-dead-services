# Consul Cleanup Dead Services

A Go application to periodically remove consul catalog of dead services.

This application connects to the Consul API to periodically query the catalog for services which are registered, but have bad health checks. It then removes these services from the catalog.
