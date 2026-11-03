package discordgo

import (
	"encoding/json"
	"testing"
)

func TestMessageSendNonceJSON(t *testing.T) {
	data, err := json.Marshal(&MessageSend{Content: "hello", Nonce: "abc", EnforceNonce: true})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var fields map[string]interface{}
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if fields["nonce"] != "abc" {
		t.Errorf("nonce = %v, want abc", fields["nonce"])
	}
	if fields["enforce_nonce"] != true {
		t.Errorf("enforce_nonce = %v, want true", fields["enforce_nonce"])
	}

	data, err = json.Marshal(&MessageSend{Content: "hello"})
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	fields = nil
	if err := json.Unmarshal(data, &fields); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if _, ok := fields["nonce"]; ok {
		t.Error("nonce is sent when unset")
	}
	if _, ok := fields["enforce_nonce"]; ok {
		t.Error("enforce_nonce is sent when unset")
	}
}
