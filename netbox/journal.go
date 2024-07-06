package netbox

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	netboxclient "github.com/fbreckle/go-netbox/netbox/client"
	netboxextras "github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/go-openapi/runtime"
)

type journalAPITransport struct {
	inner  runtime.ClientTransport
	client *netboxclient.NetBoxAPI
	entry  string
}

func newJournalTransport(inner runtime.ClientTransport, client *netboxclient.NetBoxAPI, entry string) runtime.ClientTransport {
	return &journalAPITransport{
		inner:  inner,
		client: client,
		entry:  entry,
	}
}

func (jt *journalAPITransport) Submit(op *runtime.ClientOperation) (interface{}, error) {
	res, err := jt.inner.Submit(op)
	if err != nil {
		return res, err
	}

	if op.ID == "extras_journal-entries_create" {
		// avoid loops when writing journal
		return res, nil
	}

	// skip for some methods
	if op.Method == http.MethodGet || op.Method == http.MethodDelete {
		return res, nil
	}

	// write journal
	objectType, ok := getObjectType(op.ID)
	if !ok {
		return res, nil
	}

	id, ok := getID(res)
	if !ok {
		return res, nil
	}

	m := models.WritableJournalEntry{
		AssignedObjectType: &objectType,
		AssignedObjectID:   &id,

		Kind:     models.WritableJournalEntryKindSuccess,
		Comments: &jt.entry,
		Tags:     []*models.NestedTag{},
	}
	p := netboxextras.NewExtrasJournalEntriesCreateParams().WithData(&m)
	if _, err := jt.client.Extras.ExtrasJournalEntriesCreate(p, nil); err != nil {
		return nil, fmt.Errorf("failed to create journal entry: %w", err)
	}

	return res, nil
}

func getObjectType(opID string) (string, bool) {
	parts := strings.SplitN(opID, "_", 3)
	if len(parts) < 3 {
		return "", false
	}
	group := parts[0]
	model := strings.TrimSuffix(parts[1], "s")
	return group + "." + model, true
}

func getID(res interface{}) (int64, bool) {
	getter := reflect.ValueOf(res).MethodByName("GetPayload")
	if getter.IsZero() {
		return 0, false
	}
	pl := getter.Call(nil)[0]
	if pl.IsNil() {
		return 0, false
	}
	pl = pl.Elem() // deref pointer
	id := pl.FieldByName("ID")
	if id.IsZero() {
		return 0, false
	}
	return id.Int(), true
}
