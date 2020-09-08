// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by "modeldecoder/generator". DO NOT EDIT.

package rumv3

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"
)

func (m *metadataRoot) IsSet() bool {
	return m.Metadata.IsSet()
}

func (m *metadataRoot) Reset() {
	m.Metadata.Reset()
}

func (m *metadataRoot) validate() error {
	if err := m.Metadata.validate(); err != nil {
		return err
	}
	if !m.Metadata.IsSet() {
		return fmt.Errorf("'m' required")
	}
	return nil
}

func (m *metadata) IsSet() bool {
	return len(m.Labels) > 0 || m.Service.IsSet() || m.User.IsSet()
}

func (m *metadata) Reset() {
	for k := range m.Labels {
		delete(m.Labels, k)
	}
	m.Service.Reset()
	m.User.Reset()
}

func (m *metadata) validate() error {
	if !m.IsSet() {
		return nil
	}
	for k, v := range m.Labels {
		if !labelsRegex.MatchString(k) {
			return fmt.Errorf("validation rule 'patternKeys(labelsRegex)' violated for 'm.l'")
		}
		switch t := v.(type) {
		case nil:
		case string:
			if utf8.RuneCountInString(t) > 1024 {
				return fmt.Errorf("validation rule 'maxVals(1024)' violated for 'm.l'")
			}
		case bool:
		case json.Number:
		default:
			return fmt.Errorf("validation rule 'typesVals(string;bool;number)' violated for 'm.l' for key %s", k)
		}
	}
	if err := m.Service.validate(); err != nil {
		return err
	}
	if !m.Service.IsSet() {
		return fmt.Errorf("'m.se' required")
	}
	if err := m.User.validate(); err != nil {
		return err
	}
	return nil
}

func (m *metadataService) IsSet() bool {
	return m.Agent.IsSet() || m.Environment.IsSet() || m.Framework.IsSet() || m.Language.IsSet() || m.Name.IsSet() || m.Runtime.IsSet() || m.Version.IsSet()
}

func (m *metadataService) Reset() {
	m.Agent.Reset()
	m.Environment.Reset()
	m.Framework.Reset()
	m.Language.Reset()
	m.Name.Reset()
	m.Runtime.Reset()
	m.Version.Reset()
}

func (m *metadataService) validate() error {
	if !m.IsSet() {
		return nil
	}
	if err := m.Agent.validate(); err != nil {
		return err
	}
	if !m.Agent.IsSet() {
		return fmt.Errorf("'m.se.a' required")
	}
	if utf8.RuneCountInString(m.Environment.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.en'")
	}
	if err := m.Framework.validate(); err != nil {
		return err
	}
	if err := m.Language.validate(); err != nil {
		return err
	}
	if utf8.RuneCountInString(m.Name.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.n'")
	}
	if !alphaNumericExtRegex.MatchString(m.Name.Val) {
		return fmt.Errorf("validation rule 'pattern(alphaNumericExtRegex)' violated for 'm.se.n'")
	}
	if !m.Name.IsSet() {
		return fmt.Errorf("'m.se.n' required")
	}
	if err := m.Runtime.validate(); err != nil {
		return err
	}
	if utf8.RuneCountInString(m.Version.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.ve'")
	}
	return nil
}

func (m *metadataServiceAgent) IsSet() bool {
	return m.Name.IsSet() || m.Version.IsSet()
}

func (m *metadataServiceAgent) Reset() {
	m.Name.Reset()
	m.Version.Reset()
}

func (m *metadataServiceAgent) validate() error {
	if !m.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(m.Name.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.a.n'")
	}
	if !m.Name.IsSet() {
		return fmt.Errorf("'m.se.a.n' required")
	}
	if utf8.RuneCountInString(m.Version.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.a.ve'")
	}
	if !m.Version.IsSet() {
		return fmt.Errorf("'m.se.a.ve' required")
	}
	return nil
}

func (m *MetadataServiceFramework) IsSet() bool {
	return m.Name.IsSet() || m.Version.IsSet()
}

func (m *MetadataServiceFramework) Reset() {
	m.Name.Reset()
	m.Version.Reset()
}

func (m *MetadataServiceFramework) validate() error {
	if !m.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(m.Name.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.fw.n'")
	}
	if utf8.RuneCountInString(m.Version.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.fw.ve'")
	}
	return nil
}

func (m *metadataServiceLanguage) IsSet() bool {
	return m.Name.IsSet() || m.Version.IsSet()
}

func (m *metadataServiceLanguage) Reset() {
	m.Name.Reset()
	m.Version.Reset()
}

func (m *metadataServiceLanguage) validate() error {
	if !m.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(m.Name.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.la.n'")
	}
	if !m.Name.IsSet() {
		return fmt.Errorf("'m.se.la.n' required")
	}
	if utf8.RuneCountInString(m.Version.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.la.ve'")
	}
	return nil
}

func (m *metadataServiceRuntime) IsSet() bool {
	return m.Name.IsSet() || m.Version.IsSet()
}

func (m *metadataServiceRuntime) Reset() {
	m.Name.Reset()
	m.Version.Reset()
}

func (m *metadataServiceRuntime) validate() error {
	if !m.IsSet() {
		return nil
	}
	if utf8.RuneCountInString(m.Name.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.ru.n'")
	}
	if !m.Name.IsSet() {
		return fmt.Errorf("'m.se.ru.n' required")
	}
	if utf8.RuneCountInString(m.Version.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.se.ru.ve'")
	}
	if !m.Version.IsSet() {
		return fmt.Errorf("'m.se.ru.ve' required")
	}
	return nil
}

func (m *metadataUser) IsSet() bool {
	return m.ID.IsSet() || m.Email.IsSet() || m.Name.IsSet()
}

func (m *metadataUser) Reset() {
	m.ID.Reset()
	m.Email.Reset()
	m.Name.Reset()
}

func (m *metadataUser) validate() error {
	if !m.IsSet() {
		return nil
	}
	switch t := m.ID.Val.(type) {
	case string:
		if utf8.RuneCountInString(t) > 1024 {
			return fmt.Errorf("validation rule 'max(1024)' violated for 'm.u.id'")
		}
	case json.Number:
		if _, err := t.Int64(); err != nil {
			return fmt.Errorf("validation rule 'types(string;int)' violated for 'm.u.id'")
		}
	case int:
	case nil:
	default:
		return fmt.Errorf("validation rule 'types(string;int)' violated for 'm.u.id'")
	}
	if utf8.RuneCountInString(m.Email.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.u.em'")
	}
	if utf8.RuneCountInString(m.Name.Val) > 1024 {
		return fmt.Errorf("validation rule 'max(1024)' violated for 'm.u.un'")
	}
	return nil
}
