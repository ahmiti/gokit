// Package id generates ULIDs with optional prefixes.
//
// ULIDs are preferred over UUIDs because they are lexicographically sortable by creation time.
//
// Example:
//
//   requestID := id.New("req")  // req_01HX9T7Z2KQM...
//   orderID := id.New("ord")    // ord_01HX9T...
package id
