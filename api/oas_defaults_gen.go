// Code generated by ogen, DO NOT EDIT.

package api

// setDefaults set default value of fields.
func (s *Input) setDefaults() {
	{
		val := string("backstage.io/v1alpha1")
		s.ApiVersion.SetTo(val)
	}
	{
		val := string("User")
		s.Kind.SetTo(val)
	}
}

// setDefaults set default value of fields.
func (s *User) setDefaults() {
	{
		val := string("backstage.io/v1alpha1")
		s.ApiVersion.SetTo(val)
	}
	{
		val := string("User")
		s.Kind.SetTo(val)
	}
}