package phrase

import (
	"github.com/stinkyfingers/catchphraseAPI/db"
	"gopkg.in/mgo.v2/bson"
)

type Phrase string

type Category struct {
	ID      bson.ObjectId `bson:"_id" json:"id"`
	Name    string        `bson:"name" json:"name"`
	Phrases []Phrase      `bson:"phrases" json:"phrases"`
}

const collection = "category"

func All() ([]Category, error) {
	var cats []Category
	sess, err := db.CreateSession()
	if err != nil {
		return cats, err
	}
	err = sess.DB(db.DB).C(collection).Find(nil).All(&cats)
	return cats, err
}

func (c *Category) Insert() error {
	sess, err := db.CreateSession()
	if err != nil {
		return err
	}
	c.ID = bson.NewObjectId()
	return sess.DB(db.DB).C(collection).Insert(c)
}

func (c *Category) Remove() error {
	sess, err := db.CreateSession()
	if err != nil {
		return err
	}
	return sess.DB(db.DB).C(collection).RemoveId(c.ID)
}

func (c *Category) Find() error {
	sess, err := db.CreateSession()
	if err != nil {
		return err
	}
	return sess.DB(db.DB).C(collection).Find(bson.M{"name": c.Name}).One(&c)
}
