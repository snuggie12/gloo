// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"go.uber.org/zap"
)

type ApiSnapshot struct {
	Proxies   ProxiesByNamespace
	Artifacts ArtifactsByNamespace
	Endpoints EndpointsByNamespace
	Secrets   SecretsByNamespace
	Upstreams UpstreamsByNamespace
}

func (s ApiSnapshot) Clone() ApiSnapshot {
	return ApiSnapshot{
		Proxies:   s.Proxies.Clone(),
		Artifacts: s.Artifacts.Clone(),
		Endpoints: s.Endpoints.Clone(),
		Secrets:   s.Secrets.Clone(),
		Upstreams: s.Upstreams.Clone(),
	}
}

func (s ApiSnapshot) snapshotToHash() ApiSnapshot {
	snapshotForHashing := s.Clone()
	for _, proxy := range snapshotForHashing.Proxies.List() {
		resources.UpdateMetadata(proxy, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		proxy.SetStatus(core.Status{})
	}
	for _, artifact := range snapshotForHashing.Artifacts.List() {
		resources.UpdateMetadata(artifact, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}
	for _, endpoint := range snapshotForHashing.Endpoints.List() {
		resources.UpdateMetadata(endpoint, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}
	for _, secret := range snapshotForHashing.Secrets.List() {
		resources.UpdateMetadata(secret, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}
	for _, upstream := range snapshotForHashing.Upstreams.List() {
		resources.UpdateMetadata(upstream, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
		upstream.SetStatus(core.Status{})
	}

	return snapshotForHashing
}

func (s ApiSnapshot) Hash() uint64 {
	return s.hashStruct(s.snapshotToHash())
}

func (s ApiSnapshot) HashFields() []zap.Field {
	snapshotForHashing := s.snapshotToHash()
	var fields []zap.Field
	proxies := s.hashStruct(snapshotForHashing.Proxies.List())
	fields = append(fields, zap.Uint64("proxies", proxies))
	artifacts := s.hashStruct(snapshotForHashing.Artifacts.List())
	fields = append(fields, zap.Uint64("artifacts", artifacts))
	endpoints := s.hashStruct(snapshotForHashing.Endpoints.List())
	fields = append(fields, zap.Uint64("endpoints", endpoints))
	secrets := s.hashStruct(snapshotForHashing.Secrets.List())
	fields = append(fields, zap.Uint64("secrets", secrets))
	upstreams := s.hashStruct(snapshotForHashing.Upstreams.List())
	fields = append(fields, zap.Uint64("upstreams", upstreams))

	return append(fields, zap.Uint64("snapshotHash", s.hashStruct(snapshotForHashing)))
}

func (s ApiSnapshot) hashStruct(v interface{}) uint64 {
	h, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}
	return h
}