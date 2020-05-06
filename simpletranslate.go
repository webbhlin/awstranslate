package awstranslate

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
)

/*TranslateTo ...
this is a simple aws translation function for Exchanging language*/
func TranslateTo(term, source, target, region string) string {
	mySession := session.Must(session.NewSession())
	// Create a Translate client from just a session.
	client := translate.New(mySession, aws.NewConfig().WithRegion(region))
	var ti translate.TextInput
	ti.SetText(term)
	ti.SetSourceLanguageCode(source)
	ti.SetTargetLanguageCode(target)
	result, err := client.Text(&ti)
	if err != nil {
		log.Printf("%v", err)
	}
	return *result.TranslatedText
}
