package tls

var _ = &Config{WrapSession: (&Config{}).EncryptTicket}
var _ = &Config{UnwrapSession: (&Config{}).DecryptTicket}
