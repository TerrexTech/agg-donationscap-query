package donation

import (
	"encoding/json"

	util "github.com/TerrexTech/go-commonutils/commonutil"

	"github.com/TerrexTech/uuuid"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/pkg/errors"
)

// AggregateID is the global AggregateID for Inventory Aggregate.
const AggregateID int8 = 9

type Donation struct {
	ID           objectid.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	DonationID   uuuid.UUID        `bson:"donationID,omitempty" json:"donationID,omitempty"`
	ItemID       uuuid.UUID        `bson:"itemID,omitempty" json:"itemID,omitempty"`
	SKU          string            `bson:"sku,omitempty" json:"sku,omitempty"`
	Name         string            `bson:"name,omitempty" json:"name,omitempty"`
	SoldWeight   float64           `bson:"soldWeight,omitempty" json:"soldWeight,omitempty"`
	TotalWeight  float64           `bson:"totalWeight,omitempty" json:"totalWeight,omitempty"`
	UnsoldWeight float64           `bson:"unsoldWeight,omitempty" json:"unsoldWeight,omitempty"`
	Lot          string            `bson:"lot,omitempty" json:"lot,omitempty"`
	Status       string            `bson:"status,omitempty" json:"status,omitempty"`
	Timestamp    int64             `bson:"timestamp,omitempty" json:"timestamp,omitempty"`
	DonateWeight float64           `bson:"donateWeight,omitempty" json:"donateWeight,omitempty"`
}

func (i Donation) MarshalBSON() ([]byte, error) {
	in := map[string]interface{}{
		"donationID":   i.DonationID.String(),
		"itemID":       i.ItemID.String(),
		"sku":          i.SKU,
		"name":         i.Name,
		"soldWeight":   i.SoldWeight,
		"totalWeight":  i.TotalWeight,
		"timestamp":    i.Timestamp,
		"unsoldWeight": i.UnsoldWeight,
		"status":       i.Status,
		"lot":          i.Lot,
		"donateWeight": i.DonateWeight,
	}

	if i.ID != objectid.NilObjectID {
		in["_id"] = i.ID
	}
	return bson.Marshal(in)
}

// MarshalJSON returns bytes of JSON-type.
func (i Donation) MarshalJSON() ([]byte, error) {
	in := map[string]interface{}{
		"donationID":   i.DonationID.String(),
		"itemID":       i.ItemID.String(),
		"sku":          i.SKU,
		"name":         i.Name,
		"soldWeight":   i.SoldWeight,
		"totalWeight":  i.TotalWeight,
		"timestamp":    i.Timestamp,
		"unsoldWeight": i.UnsoldWeight,
		"status":       i.Status,
		"lot":          i.Lot,
		"donateWeight": i.DonateWeight,
	}

	if i.ID != objectid.NilObjectID {
		in["_id"] = i.ID.Hex()
	}
	return json.Marshal(in)
}

func (i *Donation) UnmarshalBSON(in []byte) error {
	m := make(map[string]interface{})
	err := bson.Unmarshal(in, m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = i.unmarshalFromMap(m)
	return err
}

func (i *Donation) UnmarshalJSON(in []byte) error {
	m := make(map[string]interface{})
	err := json.Unmarshal(in, &m)
	if err != nil {
		err = errors.Wrap(err, "Unmarshal Error")
		return err
	}

	err = i.unmarshalFromMap(m)
	return err
}

// unmarshalFromMap unmarshals Map into Inventory.
func (i *Donation) unmarshalFromMap(m map[string]interface{}) error {
	var err error
	var assertOK bool

	// Hoping to discover a better way to do this someday
	if m["_id"] != nil {
		i.ID, assertOK = m["_id"].(objectid.ObjectID)
		if !assertOK {
			i.ID, err = objectid.FromHex(m["_id"].(string))
			if err != nil {
				err = errors.Wrap(err, "Error while asserting ObjectID")
				return err
			}
		}
	}

	if m["donationID"] != nil {
		i.DonationID, err = uuuid.FromString(m["donationID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting donationID")
			return err
		}
	}

	if m["itemID"] != nil {
		i.ItemID, err = uuuid.FromString(m["itemID"].(string))
		if err != nil {
			err = errors.Wrap(err, "Error while asserting ItemID")
			return err
		}
	}

	if m["lot"] != nil {
		i.Lot, assertOK = m["lot"].(string)
		if !assertOK {
			return errors.New("Error while asserting Lot")
		}
	}
	if m["name"] != nil {
		i.Name, assertOK = m["name"].(string)
		if !assertOK {
			return errors.New("Error while asserting Name")
		}
	}

	if m["status"] != nil {
		i.Status, assertOK = m["status"].(string)
		if !assertOK {
			return errors.New("Error while asserting Name")
		}
	}

	if m["sku"] != nil {
		i.SKU, assertOK = m["sku"].(string)
		if !assertOK {
			return errors.New("Error while asserting Sku")
		}
	}
	if m["soldWeight"] != nil {
		i.SoldWeight, err = util.AssertFloat64(m["soldWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting SoldWeight")
			return err
		}
	}
	if m["timestamp"] != nil {
		i.Timestamp, err = util.AssertInt64(m["timestamp"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting Timestamp")
			return err
		}
	}
	if m["totalWeight"] != nil {
		i.TotalWeight, err = util.AssertFloat64(m["totalWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting TotalWeight")
			return err
		}
	}
	if m["unsoldWeight"] != nil {
		i.UnsoldWeight, err = util.AssertFloat64(m["unsoldWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting unsoldWeight")
			return err
		}
	}
	if m["donateWeight"] != nil {
		i.DonateWeight, err = util.AssertFloat64(m["donateWeight"])
		if err != nil {
			err = errors.Wrap(err, "Error while asserting donateWeight")
			return err
		}
	}

	return nil
}
