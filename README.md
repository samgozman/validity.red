# validity.red

[![Unit tests](https://github.com/samgozman/validity.red/actions/workflows/unit_test.yml/badge.svg?branch=main)](https://github.com/samgozman/validity.red/actions/workflows/unit_test.yml)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=samgozman_validity.red&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=samgozman_validity.red)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=samgozman_validity.red&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=samgozman_validity.red)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=samgozman_validity.red&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=samgozman_validity.red)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=samgozman_validity.red&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=samgozman_validity.red)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=samgozman_validity.red&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=samgozman_validity.red)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/760a8bc3dd734e10b7b820cda8b6267f)](https://www.codacy.com/gh/samgozman/validity.red/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=samgozman/validity.red&amp;utm_campaign=Badge_Grade)

## Services

```mermaid
graph LR
    A(frontend-service: TS, Vue) === |REST| B((gateway-service: Go))
    B --- J[(Redis DB)]
    B --- |gRPC| C(user-service: Go) --- F[(Postgres)]
    B --- |gRPC| D(document-service: Go) --- G[(Postgres)]
    B --- |gRPC| E(calendar-service: Rust)
    B -.- |REST| H([MailerSend: emails service])
    K([Sentry: app error tracking])
    L([New Relic: server monitoring])
    M([Backblaze cloud storage as S3])
    
    click A "https://github.com/samgozman/validity.red/tree/main/frontend-service" "Frontend SPA written in VueJS and TypeScript"
    click B "https://github.com/samgozman/validity.red/tree/main/gateway-service" "Gateway service written in Go"
    click C "https://github.com/samgozman/validity.red/tree/main/user-service" "Users service written in Go"
    click D "https://github.com/samgozman/validity.red/tree/main/document-service" "Documents service written in Go"
    click E "https://github.com/samgozman/validity.red/tree/main/calendar-service" "Calendar service written in Rust"
```
