// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.

package common

import (
	"context"
	"reflect"
	"testing"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	commonv1 "github.com/elastic/cloud-on-k8s/v3/pkg/apis/common/v1"
)

func TestLowestVersionFromPods(t *testing.T) {
	versionLabel := "version-label"
	type args struct {
		currentVersion string
		pods           []corev1.Pod
		versionLabel   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "all pods have the same version: return it",
			args: args{
				pods: []corev1.Pod{
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{versionLabel: "7.7.0"}}},
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{versionLabel: "7.7.0"}}},
				},
				currentVersion: "",
				versionLabel:   versionLabel,
			},
			want: "7.7.0",
		},
		{
			name: "return the lowest running version",
			args: args{
				pods: []corev1.Pod{
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{versionLabel: "7.7.0"}}},
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{versionLabel: "7.6.0"}}},
				},
				currentVersion: "",
				versionLabel:   versionLabel,
			},
			want: "7.6.0",
		},
		{
			name: "cannot parse version from pods: return the current version",
			args: args{
				pods: []corev1.Pod{
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{versionLabel: "invalid"}}},
				},
				currentVersion: "7.7.0",
				versionLabel:   versionLabel,
			},
			want: "7.7.0",
		},
		{
			name: "no pods: return the current version",
			args: args{
				pods:           []corev1.Pod{},
				currentVersion: "7.7.0",
				versionLabel:   versionLabel,
			},
			want: "7.7.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowestVersionFromPods(context.Background(), tt.args.currentVersion, tt.args.pods, tt.args.versionLabel); got != tt.want {
				t.Errorf("LowestVersionFromPods() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeploymentStatus(t *testing.T) {
	type args struct {
		current      commonv1.DeploymentStatus
		dep          appsv1.Deployment
		pods         []corev1.Pod
		versionLabel string
	}
	tests := []struct {
		name    string
		args    args
		want    commonv1.DeploymentStatus
		wantErr bool
	}{
		{
			name: "happy path: set all status fields",
			args: args{
				current: commonv1.DeploymentStatus{},
				dep: appsv1.Deployment{
					Status: appsv1.DeploymentStatus{
						Replicas:          4,
						AvailableReplicas: 3,
						Conditions: []appsv1.DeploymentCondition{
							{
								Type:   appsv1.DeploymentAvailable,
								Status: corev1.ConditionTrue,
							},
						},
					},
				},
				pods: []corev1.Pod{
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"version-label": "7.7.0"}}},
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"version-label": "7.7.0"}}},
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"version-label": "7.7.0"}}},
				},
				versionLabel: "version-label",
			},
			want: commonv1.DeploymentStatus{
				Count:          4,
				AvailableNodes: 3,
				Version:        "7.7.0",
				Health:         commonv1.GreenHealth,
			},
		},
		{
			name: "red health",
			args: args{
				current: commonv1.DeploymentStatus{},
				dep: appsv1.Deployment{
					Status: appsv1.DeploymentStatus{
						Replicas:          4,
						AvailableReplicas: 3,
						Conditions: []appsv1.DeploymentCondition{
							{
								Type:   appsv1.DeploymentAvailable,
								Status: corev1.ConditionFalse,
							},
						},
					},
				},
				pods: []corev1.Pod{
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"version-label": "7.7.0"}}},
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"version-label": "7.7.0"}}},
					{ObjectMeta: metav1.ObjectMeta{Labels: map[string]string{"version-label": "7.7.0"}}},
				},
				versionLabel: "version-label",
			},
			want: commonv1.DeploymentStatus{
				Count:          4,
				AvailableNodes: 3,
				Version:        "7.7.0",
				Health:         commonv1.RedHealth,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeploymentStatus(context.Background(), tt.args.current, tt.args.dep, tt.args.pods, tt.args.versionLabel)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeploymentStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeploymentStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
