package gmail

import (
	"fmt"
	gmailApi "google.golang.org/api/gmail/v1"
)

type API struct {
	client gmailApi.UsersService
}

func NewAPI(client gmailApi.UsersService) *API {
	return &API{
		client: client,
	}
}

func (api *API) GetUserEmailsWithLabels(userID string, labels []string) error {
	messagesCall := api.client.Messages.List(userID)
	messagesCall.LabelIds(labels...)

	response, err := messagesCall.Do()
	if err != nil {
		// TODO: Mejorar el error
		return err
	}

	fmt.Println(response)
	return nil
}
