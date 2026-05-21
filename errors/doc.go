// Package errors provides typed application errors.
//
// Eight error kinds with HTTP status mapping:
//   - NotFound    (404)
//   - Validation  (400)
//   - Conflict    (409)
//   - Unauthorized (401)
//   - Forbidden   (403)
//   - RateLimited (429)
//   - Unavailable (503)
//   - Internal    (500)
//
// Example:
//
//   return errors.New(errors.NotFound, "order.not_found", "order not found")
//   return errors.New(errors.Validation, "invalid_email", "email is required").WithCause(err)
package errors
