provider "twilio" {}


resource "twilio_chat_service" "test" {
    friendly_name = "test10"
}