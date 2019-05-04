package autorestapi

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/siskinc/mgorm"
)

type Model interface {
	GetID() bson.ObjectId
	New() (Model, error)
	Check() error
	Query() interface{}
	Collection() (*mgo.Collection, error)
}

func Save(model Model) (err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	err = mgorm.Save(collection, model, model.GetID())
	return
}

func Find(model Model) (err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	query := model.Query()
	err = mgorm.FindOne(collection, query, model)
	return
}

func FindOne(model Model, query interface{}) (result Model, err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	err = mgorm.FindOne(collection, query, result)
	return
}

func FindPage(model Model, query interface{}, pageSize, pageIndex int, sorted string) (result []Model, count int, err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	count, err = Count(model, query)
	if err != nil {
		return
	}
	err = collection.Find(query).Skip((pageIndex - 1) * pageSize).Limit(pageSize).Sort(sorted).All(&result)
	return
}

func Count(model Model, query interface{}) (count int, err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	count, err = mgorm.Count(collection, query)
	return
}

func Delete(model Model) (err error) {
	query := model.Query()
	err = DeleteOne(model, query)
	return
}

func DeleteOne(model Model, query interface{}) (err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	err = mgorm.Delete(collection, query)
	return
}

func DeleteAll(model Model, query interface{}) (err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	err = mgorm.DeleteAll(collection, query)
	return
}

func Update(model Model, update interface{}) (err error) {
	query := model.Query()
	UpdateOne(model, query, update)
	return
}

func UpdateOne(model Model, query interface{}, update interface{}) (err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	err = mgorm.UpdateOne(collection, query, update)
	return
}

func UpdateAll(model Model, query interface{}, update interface{}) (err error) {
	var collection *mgo.Collection
	collection, err = model.Collection()
	if err != nil {
		return
	}
	err = mgorm.UpdateAll(collection, query, update)
	return
}
