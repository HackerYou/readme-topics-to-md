package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Topic struct for data needed from topic
type Topic struct {
	Body     string `bson:"body"`
	Title    string `bson:"title"`
	Category string `bson:"category"`
}

func main() {
	//Create CLI flag for directory location
	dir := flag.String("dir", ".", "Used to define the directory to output the file.")
	flag.Parse()

	directory := *dir + "/readme_topics"
	//Connect to mongodb
	session, err := mgo.Dial("mongodb://localhost/notes")
	//Check for err
	if err != nil {
		log.Fatal(err)
	}
	//Connect to the notes.topics collections
	c := session.DB("notes").C("topics")
	//Slice of Topic for data
	var Topics []Topic
	//Query collection for all topics, pass pointer to Topic slice
	err = c.Find(bson.M{}).All(&Topics)
	//Check for err
	if err != nil {
		log.Fatal(err)
	}
	//Need to make folder for topics
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.Mkdir(directory, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	re := regexp.MustCompile(" |/")
	//Interate through topics and create a file for each
	for _, v := range Topics {
		subdir := re.ReplaceAllString(v.Category, "_")

		if subdir == "" {
			subdir = "uncategorised"
		}

		if _, err := os.Stat(directory + "/" + subdir); os.IsNotExist(err) {
			err = os.Mkdir(directory+"/"+subdir, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}

		name := directory + "/" + subdir + "/" + re.ReplaceAllString(v.Title, "_") + ".md"
		fmt.Println(name)
		file, err := os.Create(name)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		file.Write([]byte(v.Body))
	}

}
