package account

import "sync"

// Account has the following properties
type Account struct {
	mu       sync.Mutex
	sold     int64
	isClosed bool
}

// Close closes the account
func (a *Account) Close() (payout int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isClosed {
		return 0, false
	}
	a.isClosed = true
	return a.sold, true
}

// Balance returns account balance
func (a *Account) Balance() (balance int64, ok bool) {
	if a.isClosed {
		return 0, false
	}
	return a.sold, true
}

// Deposit deposit the account
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if a.isClosed {
		return 0, false
	}
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.sold+amount < 0 {
		return 0, false
	}
	a.sold += amount
	return a.sold, true
}

// Open opens the account
func Open(initialDeposit int64) (a *Account) {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{sold: initialDeposit}
}
