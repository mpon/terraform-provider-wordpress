package twilio

import (
	"testing"

	"encoding/json"
)

const applicationRawJSON = `
{
	"account_sid": "ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"api_version": "2010-04-01",
	"date_created": "Mon, 22 Aug 2011 20:59:45 +0000",
	"date_updated": "Tue, 18 Aug 2015 16:48:57 +0000",
	"friendly_name": "Application Friendly Name",
	"message_status_callback": "http://www.example.com/sms-status-callback",
	"sid": "APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
	"sms_fallback_method": "GET",
	"sms_fallback_url": "http://www.example.com/sms-fallback",
	"sms_method": "GET",
	"sms_status_callback": "http://www.example.com/sms-status-callback",
	"sms_url": "http://example.com",
	"status_callback": "http://example.com",
	"status_callback_method": "GET",
	"uri": "/2010-04-01/Accounts/ACXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX/Applications/APXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX.json",
	"voice_caller_id_lookup": false,
	"voice_fallback_method": "GET",
	"voice_fallback_url": "http://www.example.com/voice-callback",
	"voice_method": "GET",
	"voice_url": "http://example.com"
}`

func TestApplicationSerialization(t *testing.T) {
	var application Application
	err := json.Unmarshal([]byte(applicationRawJSON), &application)

	if err != nil {
		t.Fatalf("error: %s", err.Error())
	}
}
