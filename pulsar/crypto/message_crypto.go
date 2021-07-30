// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package crypto

// MessageCrypto implement this interface to encrypt and decrypt messages
type MessageCrypto interface {
	// AddPublicKeyCipher
	AddPublicKeyCipher(keyNames []string, keyReader KeyReader) error

	// RemoveKeyCipher remove the key indentified by the keyname from the list
	RemoveKeyCipher(keyName string) bool

	// Encrypt the payload using the data key and update
	// message metadata with the keyname and encrypted data key
	Encrypt(encKeys []string, KeyReader KeyReader, msgMetadata MessageMetadataSupplier, payload []byte) ([]byte, error)

	// Decrypt the payload using the data key.
	// Keys used to ecnrypt the data key can be retrieved from msgMetadata
	Decrypt(msgMetadata MessageMetadataSupplier, payload []byte, KeyReader KeyReader) ([]byte, error)
}
