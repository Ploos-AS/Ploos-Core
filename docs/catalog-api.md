# Catalog read API

Proposed GET /api/v1/catalog returns plans whose plan and provider are both enabled. Results are ordered by product, provider and plan code. Seat limit is nullable. No prices, secrets or customer data are returned.

The handler and PostgreSQL query are implemented but not mounted in the executable. Mount behind verified identity middleware after provider validation. No commercial plans are seeded or activated. Database integration tests for filtering and pagination remain follow-up work.
