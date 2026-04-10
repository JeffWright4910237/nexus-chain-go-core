package main

import (
	"fmt"
)

type AuditTracker struct {
	TxHistory []string
}

func (a *AuditTracker) TrackTx(txHash string) {
	a.TxHistory = append(a.TxHistory, txHash)
}

func (a *AuditTracker) GetTrace() []string {
	return a.TxHistory
}

func main() {
	tracker := AuditTracker{}
	tracker.TrackTx("tx_audit_001")
	tracker.TrackTx("tx_audit_002")
	fmt.Printf("Audit Trace: %v\n", tracker.GetTrace())
}
