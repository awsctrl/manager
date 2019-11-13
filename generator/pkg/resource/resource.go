/*
Copyright © 2019 AWS Controller authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resource

// GetDocumentation returns the documentation link
func (in *BaseResource) GetDocumentation() string {
	return in.Documentation
}

// GetProperties returns the properties
func (in *BaseResource) GetProperties() map[string]Property {
	return in.Properties
}

// GetAttributes returns the attrs
func (in *BaseResource) GetAttributes() map[string]map[string]string {
	return in.Attributes
}

// GetDocumentation returns the documentation link
func (in *BaseProperty) GetDocumentation() string {
	return in.Documentation
}

// IsString will return if property can be a string
func (in *BaseProperty) IsString() bool {
	return in.Type == "String" || in.Type == "List"
}

// GetType returns the type from the object
func (in *BaseProperty) GetType() string {
	return in.Type
}

// GetDefault returns default values for params
func (in *BaseProperty) GetDefault() string {
	return ""
}

// IsParameter will make a property a parameter
func (in *BaseProperty) IsParameter() bool {
	switch in.GetType() {
	case "String":
		return true
	case "List":
		if in.GetItemType() != "String" {
			return false
		}
		return true
	case "Json":
		return true
	default:
		return false
	}
}

func (in *BaseProperty) GetGoType() string {
	switch in.Type {
	case "Json":
		return "string"
	case "List":
		itemtype := in.ItemType
		if in.ItemType == "Tag" {
			itemtype = "metav1alpha1.Tag"
		}

		return "[]" + itemtype
	case "String":
		return "string"
	}
	return in.Type
}

// GetRequired returns if the property is required
func (in *BaseProperty) GetRequired() bool {
	return in.Required
}

// GetUpdateType returns how it could be updated
func (in *BaseProperty) GetUpdateType() UpdateType {
	return in.UpdateType
}

// GetItemType returns the item type if it's a list or map
func (in *BaseProperty) GetItemType() string {
	return in.ItemType
}
