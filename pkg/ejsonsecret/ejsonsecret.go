package ejsonsecret

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/Shopify/ejson"
	c "github.com/bcaldwell/ci-scripts/internal/CIScriptsHelpers"
)

type ejsonSecret struct {
	Name      string                 `json:"_name"`
	Namespace string                 `json:"_namespace"`
	Data      map[string]interface{} `json:"data"`
}

type kubeSecret struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Type       string            `json:"type"`
	Metadata   map[string]string `json:"metadata"`
	Data       interface{}       `json:"data"`
}

func DeploySecret(secretsFile string, ejsonKey string) error {
	c.LogInfo("Create kubernetes secret from %s", secretsFile)

	decryptedSource, err := ejson.DecryptFile(secretsFile, "/opt/ejson/keys", ejsonKey)
	if err != nil {
		fmt.Printf("Error: failed to decrypt ejson file %s\n", err)
		os.Exit(1)
	}

	var inputSecret ejsonSecret

	if err := json.Unmarshal(decryptedSource, &inputSecret); err != nil {
		return fmt.Errorf("Failed to unmarshal decrypted json file %w", err)
	}

	if inputSecret.Name == "" {
		return fmt.Errorf("Error parsing ejson secret: _name can not be blank")
	}

	if inputSecret.Namespace == "" {
		return fmt.Errorf("Error parsing ejson secret: _namespace can not be blank")
	}
	// convert secrets to base64
	for key, value := range inputSecret.Data {
		var bytes []byte
		if s, ok := value.(string); ok {
			bytes = []byte(s)
		} else {
			bytes, _ = json.Marshal(value)
		}

		inputSecret.Data[key] = base64.StdEncoding.EncodeToString(bytes)
	}

	secret := kubeSecret{
		APIVersion: "v1",
		Kind:       "Secret",
		Type:       "Opaque",
		Metadata: map[string]string{
			"name":      inputSecret.Name,
			"namespace": inputSecret.Namespace,
		},
		Data: inputSecret.Data,
	}

	manifest, err := json.Marshal(secret)
	if err != nil {
		return fmt.Errorf("failed to marshal kubernetes secret %w", err)
	}

	c.LogInfo("Creating secret %s in %s", inputSecret.Name, inputSecret.Namespace)

	kubeApply := exec.Command("kubectl", "apply", "-f", "-")
	kubeApply.Stdin = bytes.NewReader(manifest)
	err = c.RunCommand(kubeApply)

	if err != nil {
		return fmt.Errorf("failed to apply secret manifest %v", err)
	}

	return nil
}
