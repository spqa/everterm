package client

import (
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/spqa/everterm/pkg/config"
	"github.com/spqa/everterm/pkg/gen-go/evernote"
)

func GetUserStoreClient() (*evernote.UserStoreClient, error) {
	thriftTransport, err := thrift.NewTHttpClient(config.GetUserStoreURI())
	if err != nil {
		return nil, err
	}
	client := thrift.NewTStandardClient(
		thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(thriftTransport),
		thrift.NewTBinaryProtocolFactory(true, true).GetProtocol(thriftTransport),
	)
	userStore := evernote.NewUserStoreClient(client)
	return userStore, nil
}

func GetNoteStoreClient() (*evernote.NoteStoreClient, error) {
	thriftTransport, err := thrift.NewTHttpClient(config.GetNoteStoreURI())
	if err != nil {
		return nil, err
	}
	client := thrift.NewTStandardClient(
		thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(thriftTransport),
		thrift.NewTBinaryProtocolFactory(true, true).GetProtocol(thriftTransport),
	)
	noteStore := evernote.NewNoteStoreClient(client)
	return noteStore, nil
}
