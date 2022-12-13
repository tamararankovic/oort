package neo4j

import (
	"errors"
	"github.com/c12s/oort/domain/model"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
)

type TransactionManager struct {
	driver neo4j.Driver
	dbName string
}

func NewTransactionManager(uri, dbName string) (*TransactionManager, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.NoAuth())
	if err != nil {
		return nil, err
	}
	return &TransactionManager{
		driver: driver,
		dbName: dbName,
	}, nil
}

type TransactionFunction func(transaction neo4j.Transaction) (interface{}, error)

func (manager *TransactionManager) WriteTransaction(cypher string, params map[string]interface{}, callback func(error) *model.OutboxMessage) error {
	_, err := manager.writeTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		outboxMessage := callback(err)
		if outboxMessage == nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be created")
		}
		_, err = transaction.Run(getOutboxMessageCypher(*outboxMessage))
		if err != nil {
			_ = transaction.Rollback()
			return nil, errors.New("outbox message could not be stored - " + err.Error())
		}

		return nil, result.Err()
	})
	return err
}

func (manager *TransactionManager) ReadTransaction(cypher string, params map[string]interface{}) (interface{}, error) {
	return manager.readTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(cypher, params)
		if err != nil {
			return nil, err
		}
		if result.Err() != nil {
			return nil, result.Err()
		}

		records := make([]interface{}, 0)
		for result.Next() {
			records = append(records, result.Record().Values)
		}
		return records, nil
	})
}

func (manager *TransactionManager) writeTransaction(txFunc TransactionFunction) (interface{}, error) {
	session := manager.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeWrite,
		DatabaseName: manager.dbName})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Println(err)
		}
	}(session)

	result, err := session.WriteTransaction(neo4j.TransactionWork(txFunc))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *TransactionManager) readTransaction(txFunc TransactionFunction) (interface{}, error) {
	session := manager.driver.NewSession(neo4j.SessionConfig{
		AccessMode:   neo4j.AccessModeRead,
		DatabaseName: manager.dbName})
	defer func(session neo4j.Session) {
		err := session.Close()
		if err != nil {
			log.Println(err)
		}
	}(session)

	result, err := session.ReadTransaction(neo4j.TransactionWork(txFunc))
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (manager *TransactionManager) Stop() {
	err := manager.driver.Close()
	if err != nil {
		log.Println("error while closing neo4j conn: ", err)
	}
}
