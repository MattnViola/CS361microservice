package storage

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func UploadJSON(bucketName, objectName string, jsonData map[string]interface{}) error {
    ctx := context.Background()

    // Create a Google Cloud Storage client
    client, err := storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
    if err != nil {
        return fmt.Errorf("failed to create client: %v", err)
    }
    defer client.Close()

    // Get a handle to the bucket and object
    bucket := client.Bucket(bucketName)
    obj := bucket.Object(objectName)

    // Convert the map to JSON
    jsonBytes, err := json.Marshal(jsonData)
    if err != nil {
        return fmt.Errorf("json.Marshal: %v", err)
    }

    // Create a writer to write the JSON data to the bucket
    wc := obj.NewWriter(ctx)
    defer wc.Close()

    // Write the JSON data
    if _, err := wc.Write(jsonBytes); err != nil {
        return fmt.Errorf("wc.Write: %v", err)
    }

    fmt.Printf("JSON data uploaded to bucket %s as %s\n", bucketName, objectName)
    return nil
}

func DownloadAndParseJSON(bucketName, objectName string) (map[string]interface{}, error) {
    ctx := context.Background()

    // Create a Google Cloud Storage client
    client, err := storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
    if err != nil {
        return nil, fmt.Errorf("failed to create client: %v", err)
    }
    defer client.Close()

    // Get a handle to the bucket and object
    bucket := client.Bucket(bucketName)
    obj := bucket.Object(objectName)

    // Create a reader to read the object from the bucket
    rc, err := obj.NewReader(ctx)
    if err != nil {
        return nil, fmt.Errorf("Object.NewReader: %v", err)
    }
    defer rc.Close()

    // Decode the JSON into a map
    var data map[string]interface{}
    if err := json.NewDecoder(rc).Decode(&data); err != nil {
        return nil, fmt.Errorf("json.NewDecoder: %v", err)
    }

    return data, nil
}