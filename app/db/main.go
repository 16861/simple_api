package db

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

//Painting stores name of workpiece and author name
type Painting struct {
	Name   string `bson:"Name"`
	Author string `bson:"Author"`
}

//ColectionOfPaintings is collecion of paintings and its name
type ColectionOfPaintings struct {
	Name      string     `bson:"Name"`
	Paintings []Painting `bson:"Paintings"`
}

//DB credentials
type DB struct {
	User           string
	Password       string
	Path           string
	DBName         string
	CollectionName string
}

func (d *DB) getConnection() *mgo.Session {
	// dialinfo, err := mgo.ParseURL(`mongodb://` + d.User + ":" + d.Password + "@" + d.Path + `/` + d.DBName)
	// if err != nil {
	// 	log.Fatalln("Can't parse url, err: ", err)
	// 	return nil
	// }
	//fmt.Println(`mongodb://` + d.User + ":" + d.Password + "@" + d.Path + `/` + d.DBName)
	if session, err := mgo.Dial(`mongodb://` + d.User + ":" + d.Password + "@" + d.Path + `/` + d.DBName); err != nil {
		log.Fatalln("Can't connect to DB, err: ", err)
	} else {
		return session
	}
	return nil
}

//GetPantings returns pointer to collection of paintings struct
func (d *DB) GetPantings() *ColectionOfPaintings {
	c := d.getConnection()
	defer c.Close()

	c.SetMode(mgo.Monotonic, true)

	paintings := ColectionOfPaintings{}
	coll := c.DB(d.DBName).C(d.CollectionName)

	err := coll.Find(nil).One(&paintings)
	if err != nil {
		log.Println("Can't fetch recodrds from db: err: ", err)
	}

	return &paintings

}

//AddPainting add new paintings to db
func (d *DB) AddPainting(p Painting) error {
	c := d.getConnection()
	defer c.Close()

	c.SetMode(mgo.Monotonic, true)

	coll := c.DB(d.DBName).C(d.CollectionName)
	paintings := ColectionOfPaintings{}

	err := coll.Find(nil).One(&paintings)
	if err != nil {
		return err
		//log.Println("error when fetching record for update, err: ", err)
	}

	newPaintings := ColectionOfPaintings{
		Name:      paintings.Name,
		Paintings: append(paintings.Paintings, p),
	}

	return coll.Update(paintings, newPaintings)
}
