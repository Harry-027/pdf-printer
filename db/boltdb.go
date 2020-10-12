package db

import (
	"github.com/Harry-027/pdf-printer/models"
	"github.com/Harry-027/pdf-printer/utils"
	"github.com/boltdb/bolt"
)

var (
	db  *bolt.DB
	err error
)

// Connect with bolt db ...
func Init(dbPath string) error {
	db, err = bolt.Open(dbPath, 0600, nil)
	utils.FatalErr(err)
	return err
}

// Creates the bucket ...
func CreateBucket(bucketName []byte) error {
	fn := func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucketName)
		return err
	}
	return db.Update(fn)
}

// Removes the bucket ...
func DeleteBucket(bucketName []byte) error {
	fn := func(tx *bolt.Tx) error {
		err := tx.DeleteBucket(bucketName)
		return err
	}
	return db.Update(fn)
}

// Create the db transaction ...
func CreateTxn(dataSlice [][]byte, bucketName string) {
	bucket := []byte(bucketName)
	err := DeleteBucket(bucket)
	err = CreateBucket(bucket)
	utils.LogErr(err)
	for _, rec := range dataSlice {
		err := updateRecord(rec, bucket)
		utils.LogErr(err)
	}
}

// Update the DB record for a given transaction ...
func updateRecord(record []byte, bucketName []byte) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(bucketName)
		id64, _ := b.NextSequence()
		key := utils.Itob(int(id64))
		return b.Put(key, record)
	})
	return err
}

// View the transaction ...
func ViewTxn(bucketType string) []*models.Record {
	var records []*models.Record
	_ = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketType))
		if b != nil {
			c := b.Cursor()
			for k, v := c.First(); k != nil; k, v = c.Next() {
				record := &models.Record{
					Key:   utils.Btoi(k),
					Value: string(v),
				}
				records = append(records, record)
			}
		}
		return nil
	})
	return records
}
