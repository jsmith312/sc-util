package scUtil

/**
 * Pretty much anything to do with uploading or storing of SC info
 **/
import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jsmith312/dblib"
	sc "github.com/jsmith312/soundcloud-api"
)

//RemoveFromGroup removes a given track from a group
func RemoveFromGroup(track sc.Track, client *sc.Client, group sc.Group /*, wg *sync.WaitGroup*/) {
	//defer wg.Done()
	resp, err := client.RemoveFromGroup(group.ID, track.ID)
	fmt.Printf("\n\033[1;33m[INFO]\tRemoved %s from %s with response: %d\n", track.Title, group.Name, resp)
	if err != nil || resp != 200 {
		log.Printf("\033[0;31m[ERROR]\tThere was an error in the removal of group: %s.\n", group.Name)
		// log this kind of error
		return
	}
}

//AddToGroup to the group from the rmv group lists
func AddToGroup(track sc.Track, client *sc.Client, group sc.Group, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := client.AddToGroup(group.ID, track.ID)
	if err != nil || resp != 201 {
		log.Printf("\033[0;31m[ERROR]\tThere was an error in the adding of group: %s.\n", group.Name)
		// log this kinf of error
		return
	}
	fmt.Printf("\n\033[0;32m[INFO]\tadded %s to %s with response: %d\n", track.Title, group.Name, resp)
}

//StoreGroups stores the id and name of the group in the DB
func StoreGroups(db *sql.DB, groups []sc.Group, wg *sync.WaitGroup) {
	defer wg.Done()
	t1 := time.Now()
	dblib.StoreItem(db, groups)
	fmt.Printf("stored %d records in %s\n", len(groups), time.Since(t1))
}
