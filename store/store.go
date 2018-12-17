package store

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/edouardparis/spark/resources"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

var charges = struct {
	objects map[string]*resources.Charge
	sync.RWMutex
}{
	objects: make(map[string]*resources.Charge),
}

func GetCharge(id string) (*resources.Charge, error) {
	charge, ok := charges.objects[id]
	if !ok {
		return nil, fmt.Errorf("charge with id: %s does not exist", id)
	}
	return charge, nil
}

func InsertCharge(charge *resources.Charge) {
	charges.Lock()
	defer charges.Unlock()

	var id string
	exists := true
	for exists {
		id = fmt.Sprintf("ch_%s", RandString(29))
		_, exists = charges.objects[id]
	}
	charge.ID = id
	charges.objects[id] = charge
}

func UpdateCharge(charge *resources.Charge) error {
	charges.Lock()
	defer charges.Unlock()

	_, ok := charges.objects[charge.ID]
	if !ok {
		return fmt.Errorf("charge with id: %s does not exist", charge.ID)
	}

	charges.objects[charge.ID] = charge
	return nil
}
