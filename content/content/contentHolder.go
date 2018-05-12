package content

type Provider struct {
}

const contentMessage = "Välkommen till Spexflix!"

func (c *Provider) GetContentForUser(user string) string {
	return contentMessage
}
