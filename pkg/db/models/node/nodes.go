package node

import (
	"encoding/json"
	"fmt"
	"errors"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"bitbucket.com/metamorph/pkg/config"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)



func getDB() *gorm.DB {
	dbPath := config.Get("database.path")
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(
		&Node{}, 
		&NameServer{}, 
		&Partition{},
		&Filesystem{},
		&KvmPolicy{},
		&SSHPubKey{},
		&BondInterface{},
		&BondParameters{},
		&VirtualDisk{},
		&PhysicalDisk{},
		)
	return db
}

// Only for Controller, Internal
//Functions. Will return all Details
//About node. Including Credentials
func GetNodes() ([]Node, error) {

	nodes := []Node{}
	db := getDB()
	defer db.Close()
	//db.Find(&nodes)
	db.Not("state", []string{"failed", "in-transition"}).Find(&nodes)
    if len(nodes)  > 0 {
		return nodes, nil
	} else {
		return nil, errors.New("Nodes not found")
	}
}

func GetNameServers(node_uuid string) ([]NameServer, error){
	node := Node{}
	nameservers := []NameServer{}
	db := getDB()
	defer db.Close()
	node_uuid1 ,_ := uuid.Parse(node_uuid)
	db.Where("node_uuid = ?", node_uuid1).First(&node)
	db.Model(&node).Related(&nameservers)
	if len(nameservers) > 0 {
		return nameservers, nil
	} else {
		return nil, errors.New(" No record Found")
	}
}


func GetPartitions(node_uuid string) ([]Partition, error){
	node := Node{}
	partitions := []Partition{}
	db := getDB()
	defer db.Close()
	node_uuid1 ,_ := uuid.Parse(node_uuid)
	db.Where("node_uuid = ?", node_uuid1).First(&node)
	db.Model(&node).Related(&partitions)
	if len(partitions) > 0 {
		return partitions, nil
	} else {
		return nil, errors.New(" No record Found")
	}
}


func GetSSHPubKeys(node_uuid string) ([]SSHPubKey, error){
	node := Node{}
	sshPubKeys := []SSHPubKey{}
	db := getDB()
	defer db.Close()
	node_uuid1 ,_ := uuid.Parse(node_uuid)
	db.Where("node_uuid = ?", node_uuid1).First(&node)
	db.Model(&node).Related(&sshPubKeys)
	if len(sshPubKeys) > 0 {
		return sshPubKeys, nil
	} else {
		return nil, errors.New(" No record Found")
	}
}

func GetBondInterfaces(node_uuid string) ([]BondInterface, error){
	node := Node{}
	bondInterfaces := []BondInterface{}
	db := getDB()
	defer db.Close()
	node_uuid1 ,_ := uuid.Parse(node_uuid)
	db.Where("node_uuid = ?", node_uuid1).First(&node)
	db.Model(&node).Related(&bondInterfaces)
	if len(bondInterfaces) > 0 {
		return bondInterfaces, nil
	} else {
		return nil, errors.New("No record Found")
	}
}

//VirtualDisk

func GetVirtualDisks(node_uuid string) ([]VirtualDisk, error){
	node := Node{}
	virtualdisks := []VirtualDisk{}
	db := getDB()
	defer db.Close()
	node_uuid1 ,_ := uuid.Parse(node_uuid)
	db.Where("node_uuid = ?", node_uuid1).First(&node)
	db.Model(&node).Related(&virtualdisks)
	if len(virtualdisks) > 0 {
		return virtualdisks, nil
	} else {
		return nil, errors.New(" No record Found")
	}
}

func GetPhysicalDisks(virtualDiskID uint) ([]PhysicalDisk, error) {
	 
	   vdisk := VirtualDisk{}
	   physcical_disks := []PhysicalDisk{}
	   db := getDB()
	   defer db.Close()
	   db.Where("id = ?", virtualDiskID).First(&vdisk)
	   db.Model(&vdisk).Related(&physcical_disks)
	   if len(physcical_disks) > 0 {
		return physcical_disks, nil
	} else {
		return nil, errors.New(" No record Found")
	}
}


func Describe(node_uuid string) ([]byte, error) {
	node := Node{}
	db := getDB()
	defer db.Close()

	node_uuid1 ,_ := uuid.Parse(node_uuid)
	db.Where("node_uuid = ?", node_uuid1).First(&node)
	if node.NodeUUID.String() == node_uuid {
		fmt.Println(node)
		res, _ := json.Marshal(node)
		return res, nil
	} else {
		return  nil, errors.New("Node not found")
	}
}

func Create(data []byte) ( string, error) {
	db := getDB()
	defer db.Close()

	var node Node
	UUID, err := uuid.NewRandom()
	//TODO : Get UUID using Redfish Library.
	err = json.Unmarshal(data, &node)
	node.NodeUUID = UUID
	err = db.Create(&node).Error
	if err != nil {
		return "", err
	} else {
		return node.NodeUUID.String(), nil
	}
}