package entity

import "time"

type PinjamBuku struct {
	IDPinjam       uint
	IDBook         int
	IDuser         int
	TanggalPinjam  time.Time
	TanggalKembali time.Time
}
