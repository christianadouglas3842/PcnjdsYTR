package tcp

// NameAddress wraps a string and implements the
// net.Listener interface.
type NameAddress struct {
}

// Network returns the network type. Always returns "tcp".
func (n NameAddress) Network() string {
}

// String returns the address.
func (n NameAddress) String() string {
	return n.Address
}
