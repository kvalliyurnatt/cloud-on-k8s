// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package kibana

import (
	"context"
	"fmt"
	"maps"

	"go.elastic.co/apm/v2"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	kbv1 "github.com/elastic/cloud-on-k8s/v3/pkg/apis/kibana/v1"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/labels"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/metadata"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/reconciler"
	"github.com/elastic/cloud-on-k8s/v3/pkg/controller/common/tracing"
	"github.com/elastic/cloud-on-k8s/v3/pkg/utils/k8s"
)

// Constants to use for the config files in a Kibana pod.
const (
	TelemetryFilename = "telemetry.yml"
)

// ReconcileConfigSecret reconciles the expected Kibana config secret for the given Kibana resource.
// This managed secret is mounted into each pod of the Kibana deployment.
func ReconcileConfigSecret(
	ctx context.Context,
	client k8s.Client,
	kb kbv1.Kibana,
	kbSettings CanonicalConfig,
	meta metadata.Metadata,
) error {
	span, ctx := apm.StartSpan(ctx, "reconcile_config_secret", tracing.SpanTypeApp)
	defer span.End()

	settingsYamlBytes, err := kbSettings.Render()
	if err != nil {
		return err
	}

	telemetryYamlBytes, err := getTelemetryYamlBytes(client, kb)
	if err != nil {
		return err
	}

	data := map[string][]byte{
		SettingsFilename: settingsYamlBytes,
	}

	if telemetryYamlBytes != nil {
		data[TelemetryFilename] = telemetryYamlBytes
	}

	expected := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:   kb.Namespace,
			Name:        kbv1.ConfigSecret(kb.Name),
			Labels:      labels.AddCredentialsLabel(maps.Clone(meta.Labels)),
			Annotations: meta.Annotations,
		},
		Data: data,
	}

	_, err = reconciler.ReconcileSecret(ctx, client, expected, &kb)
	return err
}

// getUsage returns usage map object and its YAML bytes from this Kibana configuration Secret or nil
// if the Secret or usage key doesn't exist yet.
func getTelemetryYamlBytes(client k8s.Client, kb kbv1.Kibana) ([]byte, error) {
	var secret corev1.Secret
	if err := client.Get(context.Background(), types.NamespacedName{Namespace: kb.Namespace, Name: kbv1.ConfigSecret(kb.Name)}, &secret); err != nil {
		if apierrors.IsNotFound(err) {
			// this secret is just about to be created, we don't know usage yet
			return nil, nil
		}

		return nil, fmt.Errorf("unexpected error while getting usage secret: %w", err)
	}

	telemetryBytes, ok := secret.Data[TelemetryFilename]
	if !ok || telemetryBytes == nil {
		// secret is there, but telemetry not populated yet
		return nil, nil
	}

	return telemetryBytes, nil
}
