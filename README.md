# validity.red

## Services

```mermaid
graph TB
    A(frontend-service) --> |REST| B((gateway-service))
    B -->|gRPC| C(user-service) --> F[(Postgres)]
    B -->|gRPC| D(document-service) --> G[(Postgres)]
    B -->|gRPC| E(calendar-service)
    B -->|REST| H(mail-service)
    click A "./frontend-service" "Frontend SPA written in VueJS and TypeScript"
    click B "./gateway-service" "Gateway service written in Go"
    click C "./user-service" "Users service written in Go"
    click D "./document-service" "Documents service written in Go"
    click E "./calendar-service" "Calendar service written in Rust"
```
