package redis

import (
	"time"

	"github.com/NKTKLN/ShortURLGenerator/pkg/others"
	"github.com/NKTKLN/ShortURLGenerator/pkg/user"
)
// Create a temporary access key to verify mail fidelity and then send it
func RedisKeyGenerate(email, emailType, confirmationLinkId string, val interface{}) {
	client := RedisClientConnect()
	user.EmailSend(email, confirmationLinkId, emailType)
	others.ErrorChecking(client.Set(confirmationLinkId, val, 600*time.Second).Err())
}
