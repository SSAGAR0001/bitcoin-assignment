package db

import (
	"github.com/jinzhu/gorm"
)

type Manager interface {
	Insert(fromAddr, toAddr, blockNumber, tHash string) error
	Fetch(blockNumber string) ([]TransactionStruct, error)
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func (mgr *manager) Insert(fromAddr, toAddr, blockNumber, tHash string) error {
	var ts TransactionStruct
	ts.FromAddr, ts.ToAddr, ts.BlockNumber, ts.TransactionHash = fromAddr, toAddr, blockNumber, tHash
	mgr.db.Save(ts)
	err := mgr.db.GetErrors()
	if len(err) > 0 {
		return err[0]
	}
	return nil
}

func (mgr *manager) Fetch(blockNumber string) ([]TransactionStruct, error) {
	var ts []TransactionStruct
	// mgr.db.Where(&TransactionStruct{BlockNumber: blockNumber}).Find(&ts)
	mgr.db.Where("block_number = ?", blockNumber).Find(&ts)
	err := mgr.db.GetErrors()
	if len(err) > 0 {
		return nil, err[0]
	}
	return ts, nil
}
