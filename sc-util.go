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
	"github.com/jsmith312/go-soundcloud-api"
	"github.com/jsmith312/go-soundcloud-api/group"
	"github.com/jsmith312/go-soundcloud-api/track"
)

//RemoveFromGroup removes a given track from a group
func RemoveFromGroup(track track.Track, client *soundcloud.Client,
	group group.Group, addChannel chan group.Group, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := client.RemoveFromGroup(group.ID, track.ID)
	fmt.Printf("\n\033[1;33m[INFO]\tRemoved %s from %s with response: %d\n", track.Title, group.Name, resp)
	if err != nil || resp != 200 {
		log.Printf("\033[0;31m[ERROR]\tThere was an error in the removal of group: %s.\n", group.Name)
		// log this kind of error
		return
	}
	addChannel <- group
}

//AddToGroup to the group from the rmv group lists
func AddToGroup(track track.Track, client *soundcloud.Client,
	addChannel chan group.Group, wg *sync.WaitGroup) {
	defer wg.Done()
	group := <-addChannel
	resp, err := client.AddToGroup(group.ID, track.ID)
	if err != nil || resp != 201 {
		log.Printf("\033[0;31m[ERROR]\tThere was an error in the adding of group: %s.\n", group.Name)
		// log this kinf of error
		return
	}
	fmt.Printf("\n\033[0;32m[INFO]\tadded %s to %s with response: %d\n", track.Title, group.Name, resp)
}

//StoreGroups stores the id and name of the group in the DB
func StoreGroups(db *sql.DB, groups []group.Group, wg *sync.WaitGroup) {
	defer wg.Done()
	t1 := time.Now()
	dblib.StoreItem(db, groups)
	fmt.Printf("stored %d records in %s\n", len(groups), time.Since(t1))
}
