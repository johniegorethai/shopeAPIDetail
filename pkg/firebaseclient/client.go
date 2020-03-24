package firebaseclient

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
	"shopeAPIDetail/internal/config"
	"shopeAPIDetail/pkg/errors"
)

var (
	sharedClient = &firestore.Client{}
	credentials  = map[string]string{
		"type": "service_account",
		"project_id": "mutasishopee",
		"private_key_id": "46db0f15aacd562f44f7eafedb438870bfe0d0eb",
		"private_key": "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQC0QtcYPEqdlQnw\nOnlJhfp2KRDQNr61T2KGe1PNDZT3vVjuYK63aEZK0yvtl8HWEbL3/kDyx6zgt2R+\niTb+Q67iHD5orpL9/O4cs50tN1tRcVoKJk0RtuP9jMx+Bq+jmjrskzenD8RwVsM9\nu3bjBX2dyGYvC3jgOkX266DgeVl/ZklCZUtclJ55PdqL/utuh1koD8XGawX0xIo3\n74cNAMdn1ZTX2XOxbSDo1X1/XbIQn52NM1hAN2B7Px2FmHqjPvvRCfvzMqybBVCt\n8CzOn6B8ThuPWFuKANIYfNVgd1wvD859e1R5u15c4j2Fx9tzDzF3JI8s7hD2PYBc\nb3fQ/GjLAgMBAAECggEABVogscmvJiDz3mvrXaP6jOyTi7hnDtIQaPwbaltXYMhV\n7h9Ka2FPXk8JHwZM991RqixPOFt2W2Zb4lKSTUKsMW7I6ZRurSzdKFaq5C9LE6Qh\nGhZgs8BxjNMpNSq0p0+1JStf0SLWNccAXVewuXUtxslZmKxSBpz5pnxWK/leZLpo\n0BjMc+5/42gNs9mLbXX8ze1ajnV18HfP95NKpBhbVNpoc9qVvvcGnquomIdwTHLc\ngb6oN2JJstfBSHqCDpvH2QuLiApV2WOwylkMYZ66t+Kqn/Ca57I6Arbkx5s5ruYQ\neh+fMQDawVZZOnNoLhZVO04hjnK++5fUX0NbRIYlMQKBgQDcTSu3tJygAkvcm643\nz1ULMDp1HWSTpVUjh8yJIrozSz40qu48V7CCiphqFfp4iLPRYTQMnSJ7pwskJMZg\nEUzbVvROVuAqOZh3V1zYK6vM8G+8rWnv53WpzDGyAZpEXtgnIxzXa6XCH+qgz5eK\nsz65IMo+OQhnbYIDz6hxvdRFcQKBgQDReKns2Pt106OhucWPHMNGMfaf6GoY21HJ\nOMDAYWJvjakH3LRGhIC/CSV4E6A0ewddqrFkrCiu6iiC9mnhD8kQHsWiQ1ZccFo1\niXfCiRBUv4tUvPnG90Zd1H7z8SlGG3X8tnjkKYJhCAsTKrbHF1L18KaNqKAH0f5v\nUMxkT8wD+wKBgAClOTuQi+TsHMzIp/oB4X3m5kTxkRndoiI4g6DjOShtUAFXftsu\nZpX7Ufb9mX5A6EX0wvJGg5NZKe8xLGObqV37IzwXhRCampe+6ca4Wgh/q9Qhre3F\n/9I/huW+m3UX0gpLUApRhmrLLSTLduYxID3qmq0T/bJP39GjChtLMxQBAoGBAL8S\n50XyQ6aeMEQE1k4OOZ4dU09YTPdxj/ASQdj6vDTvroFKdHNiyKH58ODQXjGhC/4I\nBxrC47VySRP1PG2c3ZZDy5mQ/QUDQ3ZUeBbOukPkGW9plhpFUz2h6VAR6slVoVGt\nDSdrKc/i619HdkSzZOyM4RCVAa1Ag2v88wSJuZrjAoGBAMhxaIByOu3j70HwYo0b\n0u0Re/Mnt6l3q4iOBafQR/XMZFqYqiyUokumgjZH3LYT8SxmG9jwsFk9RmQaeqOR\nZd7cgjZ3zpslMyUALHZuswpjQyUrtLLMceGqM3Jtx8BCDvGSLDD7WKEoFFUWzwv/\nVJzvmHlKRla+I1na2+VcoWW9\n-----END PRIVATE KEY-----\n",
		"client_email": "firebase-adminsdk-z2c62@mutasishopee.iam.gserviceaccount.com",
		"client_id": "110225468731651732695",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-z2c62%40mutasishopee.iam.gserviceaccount.com",
	}

)

// Client ...
type Client struct {
	Client *firestore.Client
}

// NewClient ...
func NewClient(cfg *config.Config) (*Client, error) {
	var c Client
	cb, err := json.Marshal(credentials)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to marshal credentials!")
	}
	option := option.WithCredentialsJSON(cb)
	c.Client, err = firestore.NewClient(context.Background(), cfg.Firebase.ProjectID, option)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firebase client!")
	}
	return &c, err
}
