package internal

import (
	"log"
	"sync"
)

type IConfigBuilder interface {
	AddProvider(provider IConfigProvider) IConfigBuilder
	AddJSON(filenames ...string) IConfigBuilder
	AddENV(filenames ...string) IConfigBuilder
	AddExternalEnvironment() IConfigBuilder

	Build() (IConfig, error)
	MustBuild() IConfig
}

type t_ConfigBuilder struct {
	providers []IConfigProvider

	mutex sync.RWMutex
}

func NewConfigBuilder() IConfigBuilder {
	return &t_ConfigBuilder{
		providers: make([]IConfigProvider, 0),
	}
}

func (b *t_ConfigBuilder) AddProvider(provider IConfigProvider) IConfigBuilder {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	b.providers = append(b.providers, provider)
	return b
}

func (b *t_ConfigBuilder) AddJSON(filenames ...string) IConfigBuilder {
	return b.AddProvider(NewConfigProviderJSON(filenames...))
}

func (b *t_ConfigBuilder) AddENV(filenames ...string) IConfigBuilder {
	return b.AddProvider(NewConfigProviderENV(filenames...))
}

func (b *t_ConfigBuilder) AddExternalEnvironment() IConfigBuilder {
	return b.AddENV()
}

func MergeProviderData(provider_data map[string]any, all_data map[string]any) map[string]any {
	if provider_data == nil {
		return all_data
	}

	if all_data == nil {
		return provider_data
	}

	for provider_key, provider_value := range provider_data {
		all_value, all_has_value := all_data[provider_key]

		if !all_has_value {
			all_data[provider_key] = provider_value
			continue
		}

		sub_provider_map, is_provider_sub_map := provider_value.(map[string]any)
		sub_all_map, is_all_sub_map := all_value.(map[string]any)

		if is_provider_sub_map && is_all_sub_map {
			all_data[provider_key] = MergeProviderData(sub_provider_map, sub_all_map)
			continue
		}

		all_data[provider_key] = provider_value
	}

	return all_data
}

func (b *t_ConfigBuilder) Build() (IConfig, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	var data map[string]any = nil

	for _, provider := range b.providers {
		provider_data, err := provider.Load()

		if err != nil {
			return nil, err
		}

		data = MergeProviderData(provider_data, data)
	}

	if data == nil {
		data = make(map[string]any)
	}

	return NewConfig(data), nil
}

func (b *t_ConfigBuilder) MustBuild() IConfig {
	cfg, err := b.Build()

	if err != nil {
		log.Panicf("[Error]: Config could not be built due to an error: %v", err)
	}

	return cfg
}
